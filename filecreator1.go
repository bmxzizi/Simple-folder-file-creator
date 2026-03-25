package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	dossiersChan := make(chan string)

	wg.Add(2)

	go func() {
		defer wg.Done()
		defer close(dossiersChan)
		for i := 0; i < 4; i++ {
			nomDossier := fmt.Sprintf("dossier_%d", i)
			err := os.Mkdir(nomDossier, 0755)
			if err != nil {
				fmt.Printf("Erreur création dossier : %v\n", err)
				continue
			}
			fmt.Printf("Dossier créé : %s\n", nomDossier)
			dossiersChan <- nomDossier
		}
	}()

	go func() {
		defer wg.Done()
		for nomDossier := range dossiersChan {
			for j := 0; j < 4; j++ {
				nomFichier := filepath.Join(nomDossier, fmt.Sprintf("fichier_%d.txt", j))
				contenu := []byte(fmt.Sprintf("Contenu du fichier %d dans %s", j, nomDossier))
				
				err := os.WriteFile(nomFichier, contenu, 0644)
				if err != nil {
					fmt.Printf("Erreur fichier : %v\n", err)
					continue
				}
				fmt.Printf("\tFichier créé : %s\n", nomFichier)
			}
		}
	}()

	wg.Wait()
	fmt.Println("Processus terminé.")
}

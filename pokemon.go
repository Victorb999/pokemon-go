package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
    "os"
    "bufio"
    "strings"
)

type PokemonDetails struct {
	Name  string   `json:"name"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
    Weight int `json:"weight"`
    Species struct {
        Url  string `json:"url"`
    } `json:"species"`
}

type PokemonSpecie struct {
    Generation struct {
        Name string `json:"name"`
        Url  string `json:"url"`
    }
    Habitat struct {
        Name string `json:"name"`
        Url  string `json:"url"`
    }
    FlavorTextEntries []struct {
        FlavorText string `json:"flavor_text"`
    }  `json:"flavor_text_entries"`
}

func main() {
    for {
        fmt.Println("Digite nome do pokemon ou digite 's' para sair:")
        reader := bufio.NewReader(os.Stdin)
        input, err := reader.ReadString('\n')
        input = strings.TrimSpace(input)

        fmt.Println("/=====================================/")

        if input == "s" {
            break
        }

        fmt.Println("https://pokeapi.co/api/v2/pokemon/"+input)

        pokemonDetails, err := getPokemonDetails(input)
        if err != nil {
            fmt.Println("Erro ao obter detalhes do Pokémon:", err)
            return
        }

        typeArray := make([]string, 0)
        for _, t := range pokemonDetails.Types {
            typeArray = append(typeArray, t.Type.Name)
        }

       /*  jsonData, err := json.Marshal(p)
        if err != nil {
            fmt.Println("Erro ao serializar Pokémon para JSON:", err)
            return
        }

        fmt.Println(string(jsonData)) */


        pokemonSpecie, err := getPokemonSpecies(pokemonDetails.Species.Url)
        if err != nil {
            fmt.Println("Erro ao obter detalhes da espeça do pokemon:", err)
            return 
        }
        
        fmt.Println("Name:",pokemonDetails.Name)
        fmt.Println("Type:", typeArray)
        fmt.Println("Generation:", pokemonSpecie.Generation.Name)
        fmt.Println("Weight:", pokemonDetails.Weight)
        fmt.Println("Habitat:", pokemonSpecie.Habitat.Name)
        fmt.Println("Flavor Text:", pokemonSpecie.FlavorTextEntries[0].FlavorText)

        fmt.Println("/=====================================/")
    }
	
}
func getPokemonDetails(name string) (PokemonDetails, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", name)

	resp, err := http.Get(url)
	if err != nil {
		return PokemonDetails{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return PokemonDetails{}, err
	}

	var pokemonDetails PokemonDetails
	err = json.Unmarshal(body, &pokemonDetails)
	if err != nil {
		return PokemonDetails{}, err
	}

	return pokemonDetails, nil
}

func getPokemonSpecies(name string) (PokemonSpecie, error) {
    url := fmt.Sprintf(name)

    resp, err := http.Get(url)
    if err != nil {
        return PokemonSpecie{}, err
    }
    defer resp.Body.Close() 

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return PokemonSpecie{}, err
    }

    var pokemonSpecie PokemonSpecie
    err = json.Unmarshal(body, &pokemonSpecie)
    if err != nil {
        return PokemonSpecie{}, err
    }

    return pokemonSpecie, nil
}
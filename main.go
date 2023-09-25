package main

import (
	"encoding/base64"
	"fmt"
)

func criptografar(texto string, deslocamentos string) string {
	runes := []rune(texto)
	textoCriptografado := ""
	if len(deslocamentos) == 0 {
		return texto
	}
	for i, r := range runes {
		deslocamento := int(deslocamentos[i%len(deslocamentos)] - '0')
		runes[i] = r + rune(deslocamento)
		textoCriptografado += string(runes[i])
	}
	return base64.StdEncoding.EncodeToString([]byte(textoCriptografado))
}

func descriptografar(textoCriptografado string, deslocamentos string) string {
	textoCriptografadoBytes, err := base64.StdEncoding.DecodeString(textoCriptografado)
	if err != nil {
		fmt.Println("Erro na decodificação Base64:", err)
		return ""
	}

	runes := []rune(string(textoCriptografadoBytes))
	textoDescriptografado := ""
	if len(deslocamentos) == 0 {
		return textoDescriptografado
	}
	for i, r := range runes {
		deslocamento := int(deslocamentos[i%len(deslocamentos)] - '0')
		runes[i] = r - rune(deslocamento)
		textoDescriptografado += string(runes[i])
	}
	return textoDescriptografado
}

func main() {
	textoOriginal := "Olá, Mundo! 123 456 789"
	deslocamentos := "35754041543433289482284"

	fmt.Println("Texto Original:", textoOriginal)

	// Criptografar o texto original em Base64
	mensagemCriptografada := criptografar(textoOriginal, deslocamentos)
	fmt.Println("Mensagem Criptografada em Base64:", mensagemCriptografada)

	// Descriptografar a mensagem criptografada em Base64
	mensagemDescriptografada := descriptografar(mensagemCriptografada, deslocamentos)
	fmt.Println("Mensagem Descriptografada:", mensagemDescriptografada)
}

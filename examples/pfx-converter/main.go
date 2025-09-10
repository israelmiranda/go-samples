// package main

// import (
// 	"crypto/x509"
// 	"encoding/pem"
// 	"fmt"
// 	"os"
// 	"path/filepath"
// 	"strings"

// 	"software.sslmate.com/src/go-pkcs12"
// )

// func main() {
// 	// dir, _ := os.Getwd()
// 	exePath, _ := os.Executable()
// 	dir := filepath.Dir(exePath)

// 	// find certificate
// 	var certPath, keyPath string
// 	files, _ := os.ReadDir(dir)
// 	for _, f := range files {
// 		ext := strings.ToLower(filepath.Ext(f.Name()))
// 		if certPath == "" && (ext == ".crt" || ext == ".cer" || ext == ".cert") {
// 			certPath = f.Name()
// 		}
// 		if keyPath == "" && ext == ".key" {
// 			keyPath = f.Name()
// 		}
// 	}

// 	if certPath == "" || keyPath == "" {
// 		fmt.Println("❌ Coloque um certificado (.crt/.cer/.cert) e a chave (.key) na mesma pasta do exe")
// 		return
// 	}

// 	fmt.Printf("Certificado: %s\n", certPath)
// 	fmt.Printf("Chave: %s\n", keyPath)

// 	// read certificate
// 	certPEM, _ := os.ReadFile(certPath)
// 	block, _ := pem.Decode(certPEM)
// 	if block == nil {
// 		fmt.Println("❌ Certificado inválido")
// 		return
// 	}
// 	cert, err := x509.ParseCertificate(block.Bytes)
// 	if err != nil {
// 		fmt.Println("❌ Erro lendo certificado:", err)
// 		return
// 	}

// 	// read private key
// 	keyPEM, _ := os.ReadFile(keyPath)
// 	keyBlock, _ := pem.Decode(keyPEM)
// 	if keyBlock == nil {
// 		fmt.Println("❌ Chave inválida")
// 		return
// 	}
// 	priv, err := x509.ParsePKCS8PrivateKey(keyBlock.Bytes)
// 	if err != nil {
// 		priv, err = x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
// 		if err != nil {
// 			fmt.Println("❌ Erro lendo chave privada:", err)
// 			return
// 		}
// 	}

// 	// password
// 	fmt.Print("Senha para o PFX: ")
// 	var password string
// 	fmt.Scanln(&password)

// 	// generate PFX (legacy mode for broad compatibility)
// 	pfxData, err := pkcs12.Legacy.Encode(priv, cert, nil, password)
// 	if err != nil {
// 		fmt.Println("❌ Erro ao gerar PFX:", err)
// 		return
// 	}

// 	out := strings.TrimSuffix(certPath, filepath.Ext(certPath)) + ".pfx"
// 	if err := os.WriteFile(out, pfxData, 0644); err != nil {
// 		fmt.Println("❌ Erro ao salvar PFX:", err)
// 		return
// 	}

// 	fmt.Println("✅ PFX gerado com sucesso:", out)
// }

package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	pkcs12 "software.sslmate.com/src/go-pkcs12"
)

func main() {
	// find path
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("❌ Não foi possível localizar o diretório do executável:", err)
		return
	}
	dir := filepath.Dir(exePath)

	// find certificate
	certPath := ""
	keyPath := ""
	files, _ := os.ReadDir(dir)
	for _, f := range files {
		ext := strings.ToLower(filepath.Ext(f.Name()))
		if certPath == "" && (ext == ".crt" || ext == ".cer" || ext == ".cert" || ext == ".pem") {
			certPath = filepath.Join(dir, f.Name())
		}
		if keyPath == "" && (ext == ".key" || ext == ".pem" || ext == ".pk8") {
			keyPath = filepath.Join(dir, f.Name())
		}
	}

	if certPath == "" || keyPath == "" {
		fmt.Println("❌ Coloque um certificado (.crt/.cer/.cert/.pem) e a chave (.key/.pem/.pk8) na mesma pasta do executável")
		return
	}

	fmt.Printf("Certificado encontrado: %s\n", filepath.Base(certPath))
	fmt.Printf("Chave privada encontrada: %s\n", filepath.Base(keyPath))

	// read (PEM, DER)
	certData, _ := os.ReadFile(certPath)
	var cert *x509.Certificate
	block, _ := pem.Decode(certData)
	if block != nil && block.Type == "CERTIFICATE" {
		cert, err = x509.ParseCertificate(block.Bytes)
	} else {
		cert, err = x509.ParseCertificate(certData)
	}
	if err != nil {
		fmt.Println("❌ Erro lendo certificado:", err)
		return
	}

	// read private key (PEM, DER)
	keyData, _ := os.ReadFile(keyPath)
	var priv any
	keyBlock, _ := pem.Decode(keyData)
	if keyBlock != nil {
		priv, err = x509.ParsePKCS8PrivateKey(keyBlock.Bytes)
		if err != nil {
			priv, err = x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
			if err != nil {
				fmt.Println("❌ Erro lendo chave privada:", err)
				return
			}
		}
	} else {
		priv, err = x509.ParsePKCS8PrivateKey(keyData)
		if err != nil {
			priv, err = x509.ParsePKCS1PrivateKey(keyData)
			if err != nil {
				fmt.Println("❌ Erro lendo chave privada:", err)
				return
			}
		}
	}

	// password
	fmt.Print("Digite a senha para proteger o PFX: ")
	var password string
	fmt.Scanln(&password)

	// generate PFX
	pfxData, err := pkcs12.Legacy.Encode(priv, cert, nil, password)
	if err != nil {
		fmt.Println("❌ Erro ao gerar PFX:", err)
		return
	}

	// save PFX
	outPath := filepath.Join(dir, strings.TrimSuffix(filepath.Base(certPath), filepath.Ext(certPath))+".pfx")
	err = os.WriteFile(outPath, pfxData, 0644)
	if err != nil {
		fmt.Println("❌ Erro ao salvar PFX:", err)
		return
	}

	fmt.Println("✅ PFX gerado com sucesso:", filepath.Base(outPath))
}

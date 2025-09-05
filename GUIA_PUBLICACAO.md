# Como Preparar uma Biblioteca Go para o Go Reference (pkg.go.dev)

Este guia explica passo a passo como preparar uma biblioteca Go para aparecer corretamente no Go Reference (pkg.go.dev) com documentação profissional e exemplos.

## 📋 Pré-requisitos

- Repositório Git público (GitHub, GitLab, etc.)
- Go 1.21 ou superior
- Conhecimento básico de Go modules

## 🚀 Passos para Publicação

### 1. Estrutura Básica do Projeto

Primeiro, certifique-se de que seu projeto tenha a estrutura correta:

```
minha-lib/
├── go.mod
├── go.sum
├── README.md
├── LICENSE
├── doc.go              # Documentação do pacote
├── arquivo_principal.go
├── arquivo_test.go
├── .gitignore
├── .golangci.yml
├── Makefile
├── .github/
│   └── workflows/
│       ├── ci.yml
│       └── release.yml
├── cmd/
│   └── minha-cli/
│       └── main.go
└── example/
    └── main.go
```

### 2. Configuração do go.mod

Certifique-se de que seu `go.mod` esteja configurado corretamente:

```go
module github.com/seu-usuario/sua-biblioteca

go 1.21

// Descrição breve da biblioteca para melhor indexação
```

### 3. Documentação do Pacote (doc.go)

Crie um arquivo `doc.go` com documentação completa do pacote:

```go
// Package sua-biblioteca fornece funcionalidades para [descrição].
//
// Esta biblioteca oferece [características principais], incluindo
// [funcionalidade 1], [funcionalidade 2], e [funcionalidade 3].
//
// # Características
//
//   - Funcionalidade 1
//   - Funcionalidade 2
//   - API simples e intuitiva
//   - Zero dependências externas
//   - Alta performance
//
// # Uso Básico
//
// Exemplo básico de uso:
//
//	resultado := sua-biblioteca.FuncaoPrincipal()
//	fmt.Printf("Resultado: %v\n", resultado)
//
// # Casos de Uso Avançados
//
// Para operações mais complexas:
//
//	// Configuração avançada
//	config := sua-biblioteca.NovaConfig()
//	resultado := config.ExecutarOperacao()
//
// # Performance
//
// A biblioteca é otimizada para [características de performance].
package sua-biblioteca
```

### 4. Documentação das Funções

Documente todas as funções exportadas com comentários detalhados:

```go
// FuncaoPrincipal executa a operação principal da biblioteca.
//
// Esta função [descrição detalhada do que faz]. A função é segura para
// uso concorrente e [outras características importantes].
//
// Parâmetros:
//   - parametro1: Descrição do primeiro parâmetro
//   - parametro2: Descrição do segundo parâmetro
//
// Retorna:
//   - tipo: Descrição do que é retornado
//   - error: Erro caso a operação falhe
//
// Exemplo:
//
//	resultado, err := FuncaoPrincipal("valor1", 42)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Resultado: %v\n", resultado)
//
// Casos de uso comuns:
//
//	// Uso simples
//	resultado, _ := FuncaoPrincipal("teste", 1)
//
//	// Com verificação de erro
//	if resultado, err := FuncaoPrincipal("teste", 1); err != nil {
//		// tratar erro
//	}
func FuncaoPrincipal(parametro1 string, parametro2 int) (string, error) {
	// implementação
}
```

### 5. Testes com Exemplos

Crie testes que incluam exemplos executáveis:

```go
package sua-biblioteca

import (
	"fmt"
	"testing"
)

// Testes normais
func TestFuncaoPrincipal(t *testing.T) {
	resultado, err := FuncaoPrincipal("teste", 1)
	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}
	if resultado == "" {
		t.Error("Resultado não deve ser vazio")
	}
}

// Exemplos que aparecem na documentação
func ExampleFuncaoPrincipal() {
	resultado, err := FuncaoPrincipal("exemplo", 42)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}
	fmt.Printf("Resultado: %s\n", resultado)
	// Output: Resultado: exemplo processado
}

// Exemplo para caso de uso específico
func ExampleFuncaoPrincipal_operacaoCompleta() {
	// Configuração
	valor := "dados importantes"
	
	// Execução
	resultado, err := FuncaoPrincipal(valor, 100)
	if err != nil {
		// tratar erro
		return
	}
	
	// Uso do resultado
	fmt.Printf("Processado: %s\n", resultado)
	// Output: Processado: dados importantes processado
}

// Benchmarks
func BenchmarkFuncaoPrincipal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FuncaoPrincipal("benchmark", i)
	}
}
```

### 6. README.md Completo

Crie um README.md profissional:

```markdown
# Sua Biblioteca

[![Go Version](https://img.shields.io/github/go-mod/go-version/seu-usuario/sua-biblioteca)](https://golang.org/)
[![License](https://img.shields.io/github/license/seu-usuario/sua-biblioteca)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/seu-usuario/sua-biblioteca)](https://goreportcard.com/report/github.com/seu-usuario/sua-biblioteca)
[![GoDoc](https://godoc.org/github.com/seu-usuario/sua-biblioteca?status.svg)](https://godoc.org/github.com/seu-usuario/sua-biblioteca)

Descrição da sua biblioteca.

## Instalação

```bash
go get github.com/seu-usuario/sua-biblioteca
```

## Uso

[Exemplos de uso]

## API Reference

[Documentação da API]

## Contribuindo

[Instruções para contribuição]
```

### 7. GitHub Actions para CI/CD

Configure workflows para garantir qualidade:

`.github/workflows/ci.yml`:
```yaml
name: CI

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main, develop ]

jobs:
  test:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        go-version: ['1.21', '1.22']

    steps:
    - uses: actions/checkout@v4
    
    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: ${{ matrix.go-version }}

    - name: Test
      run: go test -v -race -coverprofile=coverage.out ./...

    - name: Upload coverage
      uses: codecov/codecov-action@v3
      with:
        file: ./coverage.out
```

### 8. Makefile para Automação

Crie um Makefile para facilitar o desenvolvimento:

```makefile
.PHONY: test lint fmt build clean

test:
	go test -v -race -coverprofile=coverage.out ./...

lint:
	golangci-lint run

fmt:
	gofmt -s -w .

build:
	go build -v ./...

clean:
	go clean
	rm -f coverage.out
```

### 9. Publicação e Tags

1. **Commit inicial**:
```bash
git add .
git commit -m "feat: initial library implementation"
git push origin main
```

2. **Criar tag de versão**:
```bash
git tag v1.0.0
git push origin v1.0.0
```

3. **Para versões subsequentes**:
```bash
git tag v1.0.1
git push origin v1.0.1
```

### 10. Forçar Indexação no pkg.go.dev

Após publicar, force a indexação:

```bash
# Verificar se está no Go Proxy
curl "https://proxy.golang.org/github.com/seu-usuario/sua-biblioteca/@latest"

# Solicitar indexação no pkg.go.dev
curl "https://pkg.go.dev/github.com/seu-usuario/sua-biblioteca@v1.0.0"
```

## ✅ Checklist Final

Antes de publicar, verifique:

- [ ] `go.mod` configurado corretamente
- [ ] Documentação completa em `doc.go`
- [ ] Todas as funções exportadas documentadas
- [ ] Exemplos testáveis criados
- [ ] Testes passando com `go test ./...`
- [ ] Código formatado com `go fmt`
- [ ] Linting sem erros
- [ ] README.md completo
- [ ] LICENSE incluída
- [ ] Tags de versão criadas
- [ ] Repository público no GitHub

## 📚 Exemplo Prático

Para ver um exemplo completo, consulte a biblioteca Chronus:
- **Repositório**: https://github.com/reinanbr/chronus
- **pkg.go.dev**: https://pkg.go.dev/github.com/reinanbr/chronus

## 🕐 Tempo de Indexação

Após seguir todos os passos:

1. **Go Proxy**: Disponível imediatamente após tag
2. **pkg.go.dev**: 5-15 minutos para indexação completa
3. **Primeira vez**: Pode levar até 1 hora

## 🔧 Troubleshooting

### Biblioteca não aparece no pkg.go.dev

1. Verifique se o repositório é público
2. Confirme que as tags foram enviadas: `git push origin --tags`
3. Teste se está no Go Proxy: `go list -m github.com/seu-usuario/sua-biblioteca@latest`
4. Force a indexação fazendo uma requisição para a URL

### Documentação não aparece corretamente

1. Verifique se os comentários seguem o formato Go Doc
2. Certifique-se de que `doc.go` está no pacote raiz
3. Teste localmente com `go doc`

### Exemplos não executam

1. Verifique se as funções `Example*` estão corretas
2. Confirme que o `// Output:` está presente quando necessário
3. Teste com `go test -run Example`

## 📖 Recursos Adicionais

- [Effective Go - Commentary](https://golang.org/doc/effective_go#commentary)
- [Go Doc Comments](https://go.dev/doc/comment)
- [pkg.go.dev About](https://pkg.go.dev/about)
- [Go Modules Reference](https://golang.org/ref/mod)

## 🎯 Dicas de Boas Práticas

1. **Versionamento Semântico**: Use tags no formato `vX.Y.Z`
2. **Documentação Clara**: Seja específico sobre o que cada função faz
3. **Exemplos Práticos**: Inclua casos de uso reais
4. **Testes Abrangentes**: Mantenha boa cobertura de testes
5. **API Estável**: Evite breaking changes em versões patch
6. **Performance**: Inclua benchmarks para funções críticas

---

Seguindo este guia, sua biblioteca Go será indexada corretamente no pkg.go.dev com documentação profissional e exemplos funcionais! 🚀

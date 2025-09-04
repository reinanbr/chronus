# Contributing to Chronus

Thank you for your interest in contributing to Chronus! We welcome contributions from everyone.

## Development Setup

1. **Fork the repository** on GitHub
2. **Clone your fork** locally:
   ```bash
   git clone https://github.com/yourusername/chronus.git
   cd chronus
   ```
3. **Install development tools**:
   ```bash
   make install-tools
   ```

## Development Workflow

1. **Create a feature branch**:
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make your changes** following our coding standards

3. **Run tests** to ensure everything works:
   ```bash
   make test
   ```

4. **Format and lint your code**:
   ```bash
   make fmt
   make lint
   ```

5. **Commit your changes**:
   ```bash
   git add .
   git commit -m "Add your feature description"
   ```

6. **Push to your fork**:
   ```bash
   git push origin feature/your-feature-name
   ```

7. **Create a Pull Request** on GitHub

## Code Standards

- Follow Go conventions and best practices
- Write clear, self-documenting code
- Add comments for exported functions and types
- Ensure all tests pass
- Maintain backward compatibility when possible

## Testing

- Write tests for new functionality
- Ensure existing tests continue to pass
- Aim for good test coverage
- Use table-driven tests when appropriate

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run benchmarks
go test -bench=.
```

## Documentation

- Update the README.md if you add new features
- Add examples for new functionality
- Document any breaking changes

## Pull Request Guidelines

1. **Use a clear title** that describes the change
2. **Provide a detailed description** of what changed and why
3. **Reference any related issues** using GitHub keywords (e.g., "Fixes #123")
4. **Ensure CI passes** before requesting review
5. **Keep PRs focused** - one feature or fix per PR

## Code Review Process

1. All changes require at least one review
2. Maintainers will review PRs when possible
3. Address feedback promptly
4. Once approved, maintainers will merge

## Reporting Issues

When reporting issues, please include:

- Go version
- Operating system
- Steps to reproduce
- Expected vs actual behavior
- Relevant code snippets or logs

## Questions?

Feel free to open an issue for questions or join discussions in existing issues.

Thank you for contributing!

# Contributing to Operata Events Go Module

Thank you for your interest in contributing to the Operata Events Go module! This document provides guidelines for contributing to this project.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Development Setup](#development-setup)
- [Conventional Commits](#conventional-commits)
- [Making Changes](#making-changes)
- [Testing](#testing)
- [Pull Request Process](#pull-request-process)
- [Release Process](#release-process)

## Code of Conduct

This project adheres to a code of conduct. By participating, you are expected to uphold this code. Please report unacceptable behavior to the project maintainers.

## Development Setup

### Prerequisites

- Go 1.20 or later
- Git
- Node.js (for commit linting and releases)
- Make

### Setup

1. Fork the repository
2. Clone your fork:
   ```bash
   git clone https://github.com/YOUR_USERNAME/operata-events.git
   cd operata-events
   ```

3. Install development tools:
   ```bash
   make install-tools
   ```

4. Setup git hooks:
   ```bash
   make setup-git-hooks
   ```

## Conventional Commits

This project uses [Conventional Commits](https://www.conventionalcommits.org/) specification for commit messages. This enables:

- Automatic semantic versioning
- Automatic changelog generation
- Better collaboration through clear commit history

### Format

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

### Types

- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code refactoring
- `perf`: Performance improvements
- `test`: Test changes
- `build`: Build system changes
- `ci`: CI configuration changes
- `chore`: Maintenance tasks

### Scopes

- `events`: Event struct definitions
- `examples`: Example code
- `utils`: Utility functions
- `ci`: CI/CD configuration
- `deps`: Dependencies

### Examples

```bash
# Feature commits
feat(events): add CallTranscript event support
feat(utils): add event validation helpers

# Bug fixes
fix(events): correct JSON tag for timestamp field
fix(utils): handle nil pointer in IsOperataEvent

# Documentation
docs: update README with new examples
docs(events): add godoc comments for CallSummary

# Breaking changes
feat!: remove deprecated Legacy field
feat(events)!: change ContactID structure

BREAKING CHANGE: ContactID now uses string instead of struct
```

### Tools

Get help with commit format:
```bash
make commit-help
```

Validate your commit:
```bash
make validate-commit
```

## Making Changes

### Before You Start

1. Check existing issues and pull requests
2. Create an issue for significant changes
3. Fork the repository if you haven't already

### Development Workflow

1. Create a feature branch:
   ```bash
   git checkout -b feat/your-feature-name
   ```

2. Make your changes following the project conventions

3. Write or update tests as needed

4. Run the test suite:
   ```bash
   make ci
   ```

5. Commit your changes using conventional commits:
   ```bash
   git add .
   git commit -m "feat(events): add new event type support"
   ```

6. Push to your fork:
   ```bash
   git push origin feat/your-feature-name
   ```

## Testing

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
make coverage

# Run all CI checks
make ci
```

### Test Guidelines

- Write tests for new functionality
- Ensure existing tests pass
- Add test cases for edge cases
- Use meaningful test names
- Include both positive and negative test cases

### Example Test Structure

```go
func TestNewEventType(t *testing.T) {
    tests := []struct {
        name     string
        input    []byte
        expected EventType
        wantErr  bool
    }{
        {
            name:     "valid event",
            input:    validEventJSON,
            expected: expectedEvent,
            wantErr:  false,
        },
        // Add more test cases
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation
        })
    }
}
```

## Pull Request Process

### Before Submitting

1. Ensure all tests pass: `make ci`
2. Update documentation if needed
3. Add yourself to contributors if it's your first contribution

### PR Guidelines

1. **Title**: Use conventional commit format
   - `feat(events): add new event support`
   - `fix(utils): correct validation logic`

2. **Description**: Include:
   - What changes were made
   - Why the changes were necessary
   - Any breaking changes
   - Related issues

3. **Checklist**:
   - [ ] Tests pass locally
   - [ ] Code follows project conventions
   - [ ] Documentation updated if needed
   - [ ] Commit messages follow conventional format
   - [ ] No breaking changes (or clearly documented)

### Review Process

1. Automated checks must pass (CI, linting, tests)
2. At least one maintainer review required
3. All conversations must be resolved
4. No merge conflicts

## Release Process

Releases are automated using semantic-release:

1. **Patch Release**: Bug fixes (`fix:` commits)
2. **Minor Release**: New features (`feat:` commits)
3. **Major Release**: Breaking changes (`feat!:` or `BREAKING CHANGE:`)

### Release Workflow

1. Changes are merged to `main` branch
2. GitHub Actions runs semantic-release
3. Version is automatically determined from commit history
4. Release notes are generated from commits
5. Git tag and GitHub release are created
6. CHANGELOG.md is updated

### Version Bumping

- `fix:` → patch (0.0.1)
- `feat:` → minor (0.1.0)
- `feat!:` or `BREAKING CHANGE:` → major (1.0.0)

## Questions?

If you have questions about contributing:

1. Check existing issues and discussions
2. Create a new issue with the `question` label
3. Reach out to maintainers

Thank you for contributing! 🎉

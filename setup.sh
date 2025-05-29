#!/bin/bash

# Setup script for g4n Math Library
# This script helps new users get started quickly

set -e

echo "g4n Math Library Setup"
echo "=========================="
echo

# Function to check if command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Check prerequisites
echo "Checking prerequisites..."

# Check Go
if ! command_exists go; then
    echo "❌ Go is not installed. Please install Go 1.24+ from https://golang.org/"
    exit 1
fi

go_version=$(go version | grep -oE 'go[0-9]+\.[0-9]+' | sed 's/go//')
echo "✅ Go $go_version found"

# Check GCC
if ! command_exists gcc; then
    echo "❌ GCC is not installed. Please install build-essential package."
    exit 1
fi
echo "✅ GCC found"

# Check Make
if ! command_exists make; then
    echo "❌ Make is not installed. Please install make."
    exit 1
fi
echo "✅ Make found"

# Check MinGW for Windows cross-compilation
if command_exists x86_64-w64-mingw32-gcc; then
    echo "✅ MinGW-w64 found (Windows cross-compilation available)"
    MINGW_AVAILABLE=true
else
    echo "⚠️  MinGW-w64 not found (Windows cross-compilation not available)"
    echo "   To install: sudo apt-get install gcc-mingw-w64 (Ubuntu/Debian)"
    MINGW_AVAILABLE=false
fi

echo

# Build options
echo "Build Options:"
echo "1. Build for current platform only"
echo "2. Build for Linux only"
if [ "$MINGW_AVAILABLE" = true ]; then
    echo "3. Build for Windows only"
    echo "4. Build for all platforms (Linux + Windows)"
else
    echo "3. Build for Windows only (not available - MinGW not installed)"
    echo "4. Build for all platforms (not available - MinGW not installed)"
fi
echo "5. Run tests only"
echo "6. Full setup (build + test + package)"
echo

read -p "Choose an option (1-6): " choice

case $choice in
    1)
        echo "Building for current platform..."
        make local
        echo "✅ Build completed: build/libcommon4n_local.so"
        ;;
    2)
        echo "Building for Linux..."
        make linux
        echo "✅ Build completed: build/libcommon4n.so"
        ;;
    3)
        if [ "$MINGW_AVAILABLE" = true ]; then
            echo "Building for Windows..."
            make windows
            echo "✅ Build completed: build/g4n.dll"
        else
            echo "❌ Cannot build for Windows - MinGW not available"
            exit 1
        fi
        ;;
    4)
        if [ "$MINGW_AVAILABLE" = true ]; then
            echo "Building for all platforms..."
            make all
            echo "✅ Build completed:"
            echo "   - Linux: build/libcommon4n.so"
            echo "   - Windows: build/g4n.dll"
        else
            echo "❌ Cannot build for all platforms - MinGW not available"
            exit 1
        fi
        ;;
    5)
        echo "Running tests..."
        make test
        echo "✅ All tests passed"
        ;;
    6)
        echo "Full setup: building, testing, and packaging..."
        
        # Run tests first
        echo "Step 1: Running tests..."
        make test
        echo "✅ Tests passed"
        
        # Build for available platforms
        echo "Step 2: Building libraries..."
        if [ "$MINGW_AVAILABLE" = true ]; then
            make all
            echo "✅ Built for Linux and Windows"
        else
            make linux
            echo "✅ Built for Linux (Windows skipped - no MinGW)"
        fi
        
        # Create package
        echo "Step 3: Creating package..."
        make package
        echo "✅ Package created in build/package/"
        
        # Run integration tests
        echo "Step 4: Running integration tests..."
        ./test_library.sh
        echo "✅ Integration tests passed"
        ;;
    *)
        echo "Invalid option"
        exit 1
        ;;
esac

echo
echo "Setup completed successfully! 🎉"
echo

# Show next steps
echo "Next Steps:"
echo "----------"
echo "1. Check the build/ directory for generated libraries"
echo "2. Copy the appropriate library to your project:"
echo "   - Windows: g4n.dll"
echo "   - Linux: libcommon4n.so"
echo "3. Include the header file: include/g4n.h"
echo "4. See examples/ directory for usage examples"
echo "5. Read README.md for detailed documentation"
echo

# Show available files
echo "Generated Files:"
echo "---------------"
if [ -f "build/libcommon4n.so" ]; then
    echo "✅ build/libcommon4n.so (Linux library)"
fi
if [ -f "build/g4n.dll" ]; then
    echo "✅ build/g4n.dll (Windows library)"
fi
if [ -f "include/g4n.h" ]; then
    echo "✅ include/g4n.h (C header file)"
fi
if [ -d "build/package" ]; then
    echo "✅ build/package/ (Complete distribution package)"
fi

echo
echo "Happy coding! 🚀"

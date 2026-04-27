#!/usr/bin/env python3
"""
Migration script for project database migrations.
Usage: python migrate.py <migration_name>
"""

import sys
import subprocess
from pathlib import Path


def run_migration(migration_name: str) -> int:
    """
    Create a new database migration file.
    
    Args:
        migration_name: Name of the migration (e.g., "create_users_table")
    
    Returns:
        Exit code (0 for success, non-zero for failure)
    """
    if not migration_name:
        print("Error: migration file name is required as first parameter", file=sys.stderr)
        return 1
    
    
    if not migration_name.replace('_', '').isalnum():
        print(f"Warning: Migration name '{migration_name}' contains special characters", file=sys.stderr)
    
    try:
        
        result = subprocess.run(
            [
                "migrate",
                "create",
                "-ext", "sql",
                "-dir", "migrations",
                "-seq",
                migration_name
            ],
            capture_output=True,
            text=True,
            check=False  
        )
        
        
        if result.returncode == 0:
            print(f"✓ Migration created successfully: {migration_name}")
            if result.stdout:
                print(result.stdout)
            return 0
        else:
            print(f"✗ Failed to create migration: {migration_name}", file=sys.stderr)
            if result.stderr:
                print(result.stderr, file=sys.stderr)
            return result.returncode
            
    except FileNotFoundError:
        print("Error: 'migrate' command not found. Please install golang-migrate.", file=sys.stderr)
        return 1
    except Exception as e:
        print(f"Unexpected error: {e}", file=sys.stderr)
        return 1


def main() -> int:
    """Main entry point."""
    if len(sys.argv) < 2:
        print("Error: migration file name is required as first parameter", file=sys.stderr)
        print(f"Usage: {sys.argv[0]} <migration_name>", file=sys.stderr)
        return 1
    
    migration_name = sys.argv[1]
    return run_migration(migration_name)


if __name__ == "__main__":
    sys.exit(main())
# Pygo

#### Description

pygo is a Go application designed to automate the management of Python environments, both locally and globally. Its primary functionality includes creating and managing Python environments, and automatically sourcing the environment when the program runs. Additionally, it sets up directories for source code (`src`) and tests (`tests`) during the setup process.

#### Technologies
Project uses golang cli package as a base and also relies on python libraries:
- [urfave/cli v2](https://cli.urfave.org/)
- [Python 3.10+](https://www.python.org/)
- [pip 24](https://pypi.org/project/pip/)
- [pytest 8.0.2](https://docs.pytest.org/en/8.0.x/)
- [mypy 1.9](https://mypy-lang.org/)

#### Features

- **Automatic Python Environment Management**: pygo simplifies the process of managing Python environments, allowing users to create and switch between local and global environments.
- **Environment Sourcing**: The application automatically sources the designated Python environment when the program is executed, ensuring that the correct environment settings are applied.
- **Directory Setup**: Upon setup, pygo creates the necessary directories (`src` and `tests`) for organizing source code and test files.
- **Pip Handler**: After setup, adding and removing python packages becomes seemless expirience.
- **Pytest/mypy included**: Following good programing practice, tests are running under `pygo test` and not only manually written tests are checked, but optional static type as well (with `mypy` library).  


#### Installation

1. **Download or Clone the Repository**:
    `git clone https://github.com/ReiterAdam/pygo.git`
2. **Build the Application**:
    Navigate to the project directory and build the application using the `go build` command:
    `cd pygo && go build`
3. **Move the Executable**:
    Move the generated executable (`pygo`) to a directory included in your system's PATH for easy access:
    `mv pygo /usr/local/bin`

#### Usage

- **Setup Python Environment**:
  Create your project directory and nevigate there. Then, to set up a Python environment, use the following command with required flag (`--type local` or `--type global`):
    `pygo setup --type local`

    This command will create `.venv` directory that contains python virtual environment and basic structure for your project:
```
.  
├── src  
│   ├── __init__.py  
│   └── main.py  
└── tests  
   └── __init__.py
```
    If flag `--type global` is used, then globabl virtual environment is created in path `~/.pygo/.venv`

- **Run Your Program**:
    
    After setting up the environment, simply run your program using the `pygo run` command in projects root directory:
    `pygo run`
    This command sources an existing environment (flag `--type` has value `local` by default) and runs `src/main.py`. If you want to provide command line arguments for your program, it is also supported:
    `pygo run arg1 arg2`
    
- **Test Your Program**:
    
    If you want to test your program, you need to write your tests in tests directory.
    For example, you could create file `test_basic.py` starting with line `from src.main import <your function>` to import your function from `src/main.py`.
    
    Then run test from project root directory with:
    `pygo test`
    This command sources existing environment (flag `--type` has value `local` by default) and runs `pytest tests/`. 
    While the command is running, it also kicks a type checking using `mypy` library. While static typing is purely optional, this library ensures, that no type has been mistaken without even running project code. 

- **Manage external libraries**:
  
  Adding external library to your python envirionment happens with command: 
  `pygo add <library name>`
  You can also remove some package from your virtual environment with command:
  `pygo remove <library name>`
  You also type more than one package name (just like in `pip`). 
  Note, similary to other commands, `add` and `remove` will try to use flag `--type local` by default and usage of global virtual environment must be exceptionaly written with `--type global`



#### Contributing

Contributions to pygo are welcome! If you encounter any issues or have suggestions for improvements, please feel free to open an issue or submit a pull request on the [GitHub repository](https://github.com/ReiterAdam/pygo).

#### License

This project is licensed under the MIT License - see the LICENSE file for details.

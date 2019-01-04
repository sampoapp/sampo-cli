**********************************
Sampo Command-Line Interface (CLI)
**********************************

.. image:: https://img.shields.io/badge/license-Public%20Domain-blue.svg
   :alt: Project license
   :target: https://unlicense.org

.. image:: https://img.shields.io/badge/godoc-reference-blue.svg
   :alt: GoDoc reference
   :target: https://godoc.org/github.com/sampoapp/sampo-cli

.. image:: https://goreportcard.com/badge/github.com/sampoapp/sampo-cli
   :alt: Go Report Card score
   :target: https://goreportcard.com/report/github.com/sampoapp/sampo-cli

.. image:: https://img.shields.io/travis/sampoapp/sampo-cli/master.svg
   :alt: Travis CI build status
   :target: https://travis-ci.org/sampoapp/sampo-cli

|

Prerequisites
=============

To build the Sampo CLI, you will need the following software:

- `GNU Make <https://www.gnu.org/software/make/>`__ 3.81+

- `Go <https://golang.org/>`__ 1.11+

Installation
============

::

   $ go get -u github.com/sampoapp/sampo-cli/sampo

   $ sampo init  # creates the $HOME/.sampo/ directory tree

Reference
=========

::

   Sampo is a personal information manager (PIM) app.
   This is the command-line interface (CLI) for Sampo.

   Usage:
     sampo [command]

   Available Commands:
     config      Show configuration variables
     export      Export data
     help        Help about any command
     history     Show history
     import      Import data
     list        List data
     search      Search data
     today       Show today's agenda
     tomorrow    Show tomorrow's agenda

   Flags:
     -C, --config string   Set config file (default: $HOME/.sampo/config.yaml)
     -d, --debug           Enable debugging
     -h, --help            help for sampo
     -v, --verbose         Be verbose

   Use "sampo [command] --help" for more information about a command.

See Also
========

- `Sampo App for Android <https://github.com/sampoapp/sampo>`__

- `Sampo Demo Scenarios <https://github.com/sampoapp/sampo-demos>`__

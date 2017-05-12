Drone PyPi Plugin
=================

Basic pypi plugin docker container that works with `Drone <https://github.com/drone/drone>`_.


Usage
-----

.. code-block:: yaml

    pipeline:
      pypi:
        image: reinbach/drone-pypi
        repository: https://pypi.python.org/pypi/
        username: <username>
        password: <password>
        distributions: sdist,bdist_wheel

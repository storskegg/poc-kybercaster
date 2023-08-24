# poc-kybercaster

A very naive PoC implementation of strongly encrypted UDP multicast message sharing with the following features:

- Naive attempt at local memory security for private and public keys
- Post-quantum KEM using CRYSTALS Kyber
- Arm-performant AEAD with chacha20poly1305

So far, message writing is to a channel, and decrypted on the receiving end. UDP multicast is my next step. Note that
this implementation will pay no mind to "optimal UDP MTU" size, since the shared key ciphertext is already beyond that
value. So...let the hardware fragment the packets all they want.

## TODO

- Polish up the tests
- Reimplement `main.go` as a pair of "peercasters" that speak to each other using `kybercaster`
- Define message types (e.g. pubkey, general data...)
- Benchmarking
- Pub/Priv Keys are created at every program start, but maybe consider periodic rotation for long-running services

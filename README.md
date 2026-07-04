# ForgeCore

Api testing tool.

### Run

by binary name (the binary is called `forge`)

``` 
cargo run --bin forge -- --file docs/example.yaml
```

or by crate

```
cargo run -p forge-cli -- --file docs/example.yaml
```

Everything after `--` goes to the app. So `-f docs/example.yaml` works too, and `--help` / `--version`:

```
cargo run --bin forge -- --help
```

Release / installed binary:

```
cargo build --release          # → target/release/forge
./target/release/forge --file docs/example.yaml
```

or install it onto your PATH so you can just type `forge`:

```
cargo install --path crates/forge-cli
forge --file docs/example.yaml
```
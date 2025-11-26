## Release New Version
Once you push/merge your changes into the main branch, run the following commands:   
1. `git tag <version>` The version needs to start with a `v`. For example `v0.2.8`
2. `git push origin <version>`

After releasing a new version, update the `clearblade-google` provider in `main.tf` in the https://github.com/ClearBlade/terraform-google-clearblade-iot-enterprise repository

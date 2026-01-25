# Changelog

## [1.0.6](https://github.com/qeeqez/work-email-validator/compare/v1.0.5...v1.0.6) (2026-01-25)


### Bug Fixes

* update domain lists ([c706e03](https://github.com/qeeqez/work-email-validator/commit/c706e035e112e080695eaf4662b9309f9d9ab5b6))

## [1.0.5](https://github.com/qeeqez/work-email-validator/compare/v1.0.4...v1.0.5) (2026-01-18)


### Bug Fixes

* update domain lists ([21d27e1](https://github.com/qeeqez/work-email-validator/commit/21d27e16a9292251c1a725fc103195460563bf3d))

## [1.0.4](https://github.com/qeeqez/work-email-validator/compare/v1.0.3...v1.0.4) (2026-01-11)


### Bug Fixes

* update domain lists ([1a1fe4f](https://github.com/qeeqez/work-email-validator/commit/1a1fe4f80cfffb2d6cff9a96d43bb0c648b70036))

## [1.0.3](https://github.com/qeeqez/work-email-validator/compare/v1.0.2...v1.0.3) (2026-01-04)


### Bug Fixes

* update domain lists ([fea45c1](https://github.com/qeeqez/work-email-validator/commit/fea45c140ae826a303cd893fe2b316cf9da47b2f))
* update domain lists ([8161852](https://github.com/qeeqez/work-email-validator/commit/81618520ba498f2b6af182ecda0d29c2c6d531dd))

## [1.0.2](https://github.com/qeeqez/work-email-validator/compare/v1.0.1...v1.0.2) (2025-12-23)


### Bug Fixes

* its not an automatic release ([680d1b5](https://github.com/qeeqez/work-email-validator/commit/680d1b5ae16c6a91c4656c1258a010d1ae950dad))
* make patch releases on list updates ([71ab3ff](https://github.com/qeeqez/work-email-validator/commit/71ab3ffb5353fb08e9db00c45966b3b3862aaea2))

## [1.0.1](https://github.com/qeeqez/work-email-validator/compare/v1.0.0...v1.0.1) (2025-12-17)


### Bug Fixes

* proper commit person ([d7a729c](https://github.com/qeeqez/work-email-validator/commit/d7a729c904c9eaf3154452c7823830aae86ca9b5))
* run release on list updates ([5074cb6](https://github.com/qeeqez/work-email-validator/commit/5074cb648fa7d648cd945c0dfd8d9b782f8dc100))

## 1.0.0 (2025-12-17)


### Features

* add .testcoverage.yml for coverage report config. ([374e565](https://github.com/qeeqez/work-email-validator/commit/374e56523c546b7fc128a503f808edfc8e4805f9))
* add dependabot configuration for GitHub Actions and Go modules ([bb7c814](https://github.com/qeeqez/work-email-validator/commit/bb7c814c82009015a83fcc5094a2337848ca8113))
* add exclusion for 'example' path in coverage thresholds ([410a05a](https://github.com/qeeqez/work-email-validator/commit/410a05a5ce55a7a2745969fa14ed3735abd42c80))
* add initial .goreleaser.yaml configuration for builds and archives ([7ac6f17](https://github.com/qeeqez/work-email-validator/commit/7ac6f1764b467badce1fb2a6053a7453580af2bf))
* add release workflow for automated version and tag ([7081ea6](https://github.com/qeeqez/work-email-validator/commit/7081ea648372f1e61ef3dbcf7e8981ec954e02c9))
* business domains should be validated first ([9f80c51](https://github.com/qeeqez/work-email-validator/commit/9f80c513dddc9e89eb67439c330c2685caf5598c))
* improve domain loading and handling ([bfcc972](https://github.com/qeeqez/work-email-validator/commit/bfcc97264b5e5cfa0c0b198bd32ae3a84a1ce0a3))
* improve update list script ([3e95e21](https://github.com/qeeqez/work-email-validator/commit/3e95e211e1c0900b9b6dffb17e208cc9cd659886))
* support internationalized domains ([6bac838](https://github.com/qeeqez/work-email-validator/commit/6bac838c6125837fb941546a5dcd298427271a0a))


### Bug Fixes

* all linting errors ([2e15003](https://github.com/qeeqez/work-email-validator/commit/2e150039a042e3f32bc314590bde25c4b0f97e6e))
* **ci:** change cache parameter to boolean value ([2733aad](https://github.com/qeeqez/work-email-validator/commit/2733aad301de01c650ccc89b49becdb1d62c6dcd))
* **ci:** downgrade to Go 1.23 for golangci-lint compatibility ([e5a47e9](https://github.com/qeeqez/work-email-validator/commit/e5a47e968d8eba3289894160cf0519296a77a1b7))
* **ci:** update CI config  for Go version compatibility and improve coverage reporting ([727560d](https://github.com/qeeqez/work-email-validator/commit/727560deb4272ad7c7bbea6d327a4dfe32b79e5a))
* **ci:** use only Go 1.25.4 to avoid version availability issues ([718dc3b](https://github.com/qeeqez/work-email-validator/commit/718dc3b75675450d1264462d6625e6edc59df286))
* goreleaser.yaml for build config. and archive formats ([b28f17d](https://github.com/qeeqez/work-email-validator/commit/b28f17d62b413af5c9e5b9164f235d6f26372cf2))
* migrate to golangci-lint v2 and apply proper rules ([880bd5b](https://github.com/qeeqez/work-email-validator/commit/880bd5b184ef53695e5187a2663cee6da4f428d6))
* proper permissions ([3447df8](https://github.com/qeeqez/work-email-validator/commit/3447df8b7cc5c0b8db8f8e9332fc4eb64cb68058))
* set main package path for goreleaser ([2302268](https://github.com/qeeqez/work-email-validator/commit/2302268c56bbdc106d3142edde0f539db48c2dc9))
* **validator:** handle subdomains correctly and add IsWorkEmail helper ([bf0c0f0](https://github.com/qeeqez/work-email-validator/commit/bf0c0f08719083e4e207153acf9c9f2630c96a1f))

package nestContent

func PackageLockJson() []byte {
	packageLockJson := []byte(`
	{
		"name": "nest-typescript-starter",
		"version": "1.0.0",
		"lockfileVersion": 1,
		"requires": true,
		"dependencies": {
		  "@angular-devkit/core": {
			"version": "9.0.6",
			"resolved": "https://registry.npmjs.org/@angular-devkit/core/-/core-9.0.6.tgz",
			"integrity": "sha512-hCZJbnqLEm1F5Bx+ILcdd3LPgQTn4WFWpfUqMEGGj7UirRInWcz+6UpYotKGTJw85/mV01LrIbtWIkAUXbkkhg==",
			"dev": true,
			"requires": {
			  "ajv": "6.10.2",
			  "fast-json-stable-stringify": "2.0.0",
			  "magic-string": "0.25.4",
			  "rxjs": "6.5.3",
			  "source-map": "0.7.3"
			},
			"dependencies": {
			  "rxjs": {
				"version": "6.5.3",
				"resolved": "https://registry.npmjs.org/rxjs/-/rxjs-6.5.3.tgz",
				"integrity": "sha512-wuYsAYYFdWTAnAaPoKGNhfpWwKZbJW+HgAJ+mImp+Epl7BG8oNWBCTyRM8gba9k4lk8BgWdoYm21Mo/RYhhbgA==",
				"dev": true,
				"requires": {
				  "tslib": "^1.9.0"
				}
			  }
			}
		  },
		  "@angular-devkit/schematics": {
			"version": "9.0.6",
			"resolved": "https://registry.npmjs.org/@angular-devkit/schematics/-/schematics-9.0.6.tgz",
			"integrity": "sha512-X7qZDJVrFcPUn+jNUeOH7Bx1D7YTpTFr0d3DBIsQzseReSGu7ugWziQPS4gc5Xm5K0nb8vx6DYtyW0FaIvX0ZA==",
			"dev": true,
			"requires": {
			  "@angular-devkit/core": "9.0.6",
			  "ora": "4.0.2",
			  "rxjs": "6.5.3"
			},
			"dependencies": {
			  "ora": {
				"version": "4.0.2",
				"resolved": "https://registry.npmjs.org/ora/-/ora-4.0.2.tgz",
				"integrity": "sha512-YUOZbamht5mfLxPmk4M35CD/5DuOkAacxlEUbStVXpBAt4fyhBf+vZHI/HRkI++QUp3sNoeA2Gw4C+hi4eGSig==",
				"dev": true,
				"requires": {
				  "chalk": "^2.4.2",
				  "cli-cursor": "^3.1.0",
				  "cli-spinners": "^2.2.0",
				  "is-interactive": "^1.0.0",
				  "log-symbols": "^3.0.0",
				  "strip-ansi": "^5.2.0",
				  "wcwidth": "^1.0.1"
				}
			  },
			  "rxjs": {
				"version": "6.5.3",
				"resolved": "https://registry.npmjs.org/rxjs/-/rxjs-6.5.3.tgz",
				"integrity": "sha512-wuYsAYYFdWTAnAaPoKGNhfpWwKZbJW+HgAJ+mImp+Epl7BG8oNWBCTyRM8gba9k4lk8BgWdoYm21Mo/RYhhbgA==",
				"dev": true,
				"requires": {
				  "tslib": "^1.9.0"
				}
			  }
			}
		  },
		  "@angular-devkit/schematics-cli": {
			"version": "0.901.7",
			"resolved": "https://registry.npmjs.org/@angular-devkit/schematics-cli/-/schematics-cli-0.901.7.tgz",
			"integrity": "sha512-9LTyCvgiU74nc5FcwRfgHGMZSfspSsqTFgivuS627Dju2D/knmXzdrNXeqU0D4CRsWGbJBKbcAjnlgmtZDTydg==",
			"dev": true,
			"requires": {
			  "@angular-devkit/core": "9.1.7",
			  "@angular-devkit/schematics": "9.1.7",
			  "@schematics/schematics": "0.901.7",
			  "inquirer": "7.1.0",
			  "minimist": "1.2.5",
			  "rxjs": "6.5.4",
			  "symbol-observable": "1.2.0"
			},
			"dependencies": {
			  "@angular-devkit/core": {
				"version": "9.1.7",
				"resolved": "https://registry.npmjs.org/@angular-devkit/core/-/core-9.1.7.tgz",
				"integrity": "sha512-guvolu9Cl+qYMTtedLZD9wCqustJjdqzJ2psD2C1Sr1LrX9T0mprmDldR/YnhsitThveJEb6sM/0EvqWxoSvKw==",
				"dev": true,
				"requires": {
				  "ajv": "6.12.0",
				  "fast-json-stable-stringify": "2.1.0",
				  "magic-string": "0.25.7",
				  "rxjs": "6.5.4",
				  "source-map": "0.7.3"
				}
			  },
			  "@angular-devkit/schematics": {
				"version": "9.1.7",
				"resolved": "https://registry.npmjs.org/@angular-devkit/schematics/-/schematics-9.1.7.tgz",
				"integrity": "sha512-oeHPJePBcPp/bd94jHQeFUnft93PGF5iJiKV9szxqS8WWC5OMZ5eK7icRY0PwvLyfenspAZxdZcNaqJqPMul5A==",
				"dev": true,
				"requires": {
				  "@angular-devkit/core": "9.1.7",
				  "ora": "4.0.3",
				  "rxjs": "6.5.4"
				}
			  },
			  "ajv": {
				"version": "6.12.0",
				"resolved": "https://registry.npmjs.org/ajv/-/ajv-6.12.0.tgz",
				"integrity": "sha512-D6gFiFA0RRLyUbvijN74DWAjXSFxWKaWP7mldxkVhyhAV3+SWA9HEJPHQ2c9soIeTFJqcSdFDGFgdqs1iUU2Hw==",
				"dev": true,
				"requires": {
				  "fast-deep-equal": "^3.1.1",
				  "fast-json-stable-stringify": "^2.0.0",
				  "json-schema-traverse": "^0.4.1",
				  "uri-js": "^4.2.2"
				}
			  },
			  "ansi-regex": {
				"version": "5.0.0",
				"resolved": "https://registry.npmjs.org/ansi-regex/-/ansi-regex-5.0.0.tgz",
				"integrity": "sha512-bY6fj56OUQ0hU1KjFNDQuJFezqKdrAyFdIevADiqrWHwSlbmBNMHp5ak2f40Pm8JTFyM2mqxkG6ngkHO11f/lg==",
				"dev": true
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "3.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-3.0.0.tgz",
				"integrity": "sha512-4D3B6Wf41KOYRFdszmDqMCGq5VV/uMAB273JILmO+3jAlh8X4qDtdtgCR3fxtbLEMzSx22QdhnDcJvu2u1fVwg==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "fast-deep-equal": {
				"version": "3.1.1",
				"resolved": "https://registry.npmjs.org/fast-deep-equal/-/fast-deep-equal-3.1.1.tgz",
				"integrity": "sha512-8UEa58QDLauDNfpbrX55Q9jrGHThw2ZMdOky5Gl1CDtVeJDPVrG4Jxx1N8jw2gkWaff5UUuX1KJd+9zGe2B+ZA==",
				"dev": true
			  },
			  "fast-json-stable-stringify": {
				"version": "2.1.0",
				"resolved": "https://registry.npmjs.org/fast-json-stable-stringify/-/fast-json-stable-stringify-2.1.0.tgz",
				"integrity": "sha512-lhd/wF+Lk98HZoTCtlVraHtfh5XYijIjalXck7saUtuanSDyLMxnHhSXEDJqHxD7msR8D0uCmqlkwjCV8xvwHw==",
				"dev": true
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "magic-string": {
				"version": "0.25.7",
				"resolved": "https://registry.npmjs.org/magic-string/-/magic-string-0.25.7.tgz",
				"integrity": "sha512-4CrMT5DOHTDk4HYDlzmwu4FVCcIYI8gauveasrdCu2IKIFOJ3f0v/8MDGJCDL9oD2ppz/Av1b0Nj345H9M+XIA==",
				"dev": true,
				"requires": {
				  "sourcemap-codec": "^1.4.4"
				}
			  },
			  "ora": {
				"version": "4.0.3",
				"resolved": "https://registry.npmjs.org/ora/-/ora-4.0.3.tgz",
				"integrity": "sha512-fnDebVFyz309A73cqCipVL1fBZewq4vwgSHfxh43vVy31mbyoQ8sCH3Oeaog/owYOs/lLlGVPCISQonTneg6Pg==",
				"dev": true,
				"requires": {
				  "chalk": "^3.0.0",
				  "cli-cursor": "^3.1.0",
				  "cli-spinners": "^2.2.0",
				  "is-interactive": "^1.0.0",
				  "log-symbols": "^3.0.0",
				  "mute-stream": "0.0.8",
				  "strip-ansi": "^6.0.0",
				  "wcwidth": "^1.0.1"
				}
			  },
			  "rxjs": {
				"version": "6.5.4",
				"resolved": "https://registry.npmjs.org/rxjs/-/rxjs-6.5.4.tgz",
				"integrity": "sha512-naMQXcgEo3csAEGvw/NydRA0fuS2nDZJiw1YUWFKU7aPPAPGZEsD4Iimit96qwCieH6y614MCLYwdkrWx7z/7Q==",
				"dev": true,
				"requires": {
				  "tslib": "^1.9.0"
				}
			  },
			  "strip-ansi": {
				"version": "6.0.0",
				"resolved": "https://registry.npmjs.org/strip-ansi/-/strip-ansi-6.0.0.tgz",
				"integrity": "sha512-AuvKTrTfQNYNIctbR1K/YGTR1756GycPsg7b9bdV9Duqur4gv6aKqHXah67Z8ImS7WEz5QVcOtlfW2rZEugt6w==",
				"dev": true,
				"requires": {
				  "ansi-regex": "^5.0.0"
				}
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "@babel/code-frame": {
			"version": "7.8.3",
			"resolved": "https://registry.npmjs.org/@babel/code-frame/-/code-frame-7.8.3.tgz",
			"integrity": "sha512-a9gxpmdXtZEInkCSHUJDLHZVBgb1QS0jhss4cPP93EW7s+uC5bikET2twEF3KV+7rDblJcmNvTR7VJejqd2C2g==",
			"dev": true,
			"requires": {
			  "@babel/highlight": "^7.8.3"
			}
		  },
		  "@babel/core": {
			"version": "7.9.6",
			"resolved": "https://registry.npmjs.org/@babel/core/-/core-7.9.6.tgz",
			"integrity": "sha512-nD3deLvbsApbHAHttzIssYqgb883yU/d9roe4RZymBCDaZryMJDbptVpEpeQuRh4BJ+SYI8le9YGxKvFEvl1Wg==",
			"dev": true,
			"requires": {
			  "@babel/code-frame": "^7.8.3",
			  "@babel/generator": "^7.9.6",
			  "@babel/helper-module-transforms": "^7.9.0",
			  "@babel/helpers": "^7.9.6",
			  "@babel/parser": "^7.9.6",
			  "@babel/template": "^7.8.6",
			  "@babel/traverse": "^7.9.6",
			  "@babel/types": "^7.9.6",
			  "convert-source-map": "^1.7.0",
			  "debug": "^4.1.0",
			  "gensync": "^1.0.0-beta.1",
			  "json5": "^2.1.2",
			  "lodash": "^4.17.13",
			  "resolve": "^1.3.2",
			  "semver": "^5.4.1",
			  "source-map": "^0.5.0"
			},
			"dependencies": {
			  "debug": {
				"version": "4.1.1",
				"resolved": "https://registry.npmjs.org/debug/-/debug-4.1.1.tgz",
				"integrity": "sha512-pYAIzeRo8J6KPEaJ0VWOh5Pzkbw/RetuzehGM7QRRX5he4fPHx2rdKMB256ehJCkX+XRQm16eZLqLNS8RSZXZw==",
				"dev": true,
				"requires": {
				  "ms": "^2.1.1"
				}
			  },
			  "json5": {
				"version": "2.1.3",
				"resolved": "https://registry.npmjs.org/json5/-/json5-2.1.3.tgz",
				"integrity": "sha512-KXPvOm8K9IJKFM0bmdn8QXh7udDh1g/giieX0NLCaMnb4hEiVFqnop2ImTXCc5e0/oHz3LTqmHGtExn5hfMkOA==",
				"dev": true,
				"requires": {
				  "minimist": "^1.2.5"
				}
			  },
			  "ms": {
				"version": "2.1.2",
				"resolved": "https://registry.npmjs.org/ms/-/ms-2.1.2.tgz",
				"integrity": "sha512-sGkPx+VjMtmA6MX27oA4FBFELFCZZ4S4XqeGOXCv68tT+jb3vk/RyaKWP0PTKyWtmLSM0b+adUTEvbs1PEaH2w==",
				"dev": true
			  },
			  "source-map": {
				"version": "0.5.7",
				"resolved": "https://registry.npmjs.org/source-map/-/source-map-0.5.7.tgz",
				"integrity": "sha1-igOdLRAh0i0eoUyA2OpGi6LvP8w=",
				"dev": true
			  }
			}
		  },
		  "@babel/generator": {
			"version": "7.9.6",
			"resolved": "https://registry.npmjs.org/@babel/generator/-/generator-7.9.6.tgz",
			"integrity": "sha512-+htwWKJbH2bL72HRluF8zumBxzuX0ZZUFl3JLNyoUjM/Ho8wnVpPXM6aUz8cfKDqQ/h7zHqKt4xzJteUosckqQ==",
			"dev": true,
			"requires": {
			  "@babel/types": "^7.9.6",
			  "jsesc": "^2.5.1",
			  "lodash": "^4.17.13",
			  "source-map": "^0.5.0"
			},
			"dependencies": {
			  "source-map": {
				"version": "0.5.7",
				"resolved": "https://registry.npmjs.org/source-map/-/source-map-0.5.7.tgz",
				"integrity": "sha1-igOdLRAh0i0eoUyA2OpGi6LvP8w=",
				"dev": true
			  }
			}
		  },
		  "@babel/helper-function-name": {
			"version": "7.9.5",
			"resolved": "https://registry.npmjs.org/@babel/helper-function-name/-/helper-function-name-7.9.5.tgz",
			"integrity": "sha512-JVcQZeXM59Cd1qanDUxv9fgJpt3NeKUaqBqUEvfmQ+BCOKq2xUgaWZW2hr0dkbyJgezYuplEoh5knmrnS68efw==",
			"dev": true,
			"requires": {
			  "@babel/helper-get-function-arity": "^7.8.3",
			  "@babel/template": "^7.8.3",
			  "@babel/types": "^7.9.5"
			}
		  },
		  "@babel/helper-get-function-arity": {
			"version": "7.8.3",
			"resolved": "https://registry.npmjs.org/@babel/helper-get-function-arity/-/helper-get-function-arity-7.8.3.tgz",
			"integrity": "sha512-FVDR+Gd9iLjUMY1fzE2SR0IuaJToR4RkCDARVfsBBPSP53GEqSFjD8gNyxg246VUyc/ALRxFaAK8rVG7UT7xRA==",
			"dev": true,
			"requires": {
			  "@babel/types": "^7.8.3"
			}
		  },
		  "@babel/helper-member-expression-to-functions": {
			"version": "7.8.3",
			"resolved": "https://registry.npmjs.org/@babel/helper-member-expression-to-functions/-/helper-member-expression-to-functions-7.8.3.tgz",
			"integrity": "sha512-fO4Egq88utkQFjbPrSHGmGLFqmrshs11d46WI+WZDESt7Wu7wN2G2Iu+NMMZJFDOVRHAMIkB5SNh30NtwCA7RA==",
			"dev": true,
			"requires": {
			  "@babel/types": "^7.8.3"
			}
		  },
		  "@babel/helper-module-imports": {
			"version": "7.8.3",
			"resolved": "https://registry.npmjs.org/@babel/helper-module-imports/-/helper-module-imports-7.8.3.tgz",
			"integrity": "sha512-R0Bx3jippsbAEtzkpZ/6FIiuzOURPcMjHp+Z6xPe6DtApDJx+w7UYyOLanZqO8+wKR9G10s/FmHXvxaMd9s6Kg==",
			"dev": true,
			"requires": {
			  "@babel/types": "^7.8.3"
			}
		  },
		  "@babel/helper-module-transforms": {
			"version": "7.9.0",
			"resolved": "https://registry.npmjs.org/@babel/helper-module-transforms/-/helper-module-transforms-7.9.0.tgz",
			"integrity": "sha512-0FvKyu0gpPfIQ8EkxlrAydOWROdHpBmiCiRwLkUiBGhCUPRRbVD2/tm3sFr/c/GWFrQ/ffutGUAnx7V0FzT2wA==",
			"dev": true,
			"requires": {
			  "@babel/helper-module-imports": "^7.8.3",
			  "@babel/helper-replace-supers": "^7.8.6",
			  "@babel/helper-simple-access": "^7.8.3",
			  "@babel/helper-split-export-declaration": "^7.8.3",
			  "@babel/template": "^7.8.6",
			  "@babel/types": "^7.9.0",
			  "lodash": "^4.17.13"
			}
		  },
		  "@babel/helper-optimise-call-expression": {
			"version": "7.8.3",
			"resolved": "https://registry.npmjs.org/@babel/helper-optimise-call-expression/-/helper-optimise-call-expression-7.8.3.tgz",
			"integrity": "sha512-Kag20n86cbO2AvHca6EJsvqAd82gc6VMGule4HwebwMlwkpXuVqrNRj6CkCV2sKxgi9MyAUnZVnZ6lJ1/vKhHQ==",
			"dev": true,
			"requires": {
			  "@babel/types": "^7.8.3"
			}
		  },
		  "@babel/helper-plugin-utils": {
			"version": "7.8.3",
			"resolved": "https://registry.npmjs.org/@babel/helper-plugin-utils/-/helper-plugin-utils-7.8.3.tgz",
			"integrity": "sha512-j+fq49Xds2smCUNYmEHF9kGNkhbet6yVIBp4e6oeQpH1RUs/Ir06xUKzDjDkGcaaokPiTNs2JBWHjaE4csUkZQ==",
			"dev": true
		  },
		  "@babel/helper-replace-supers": {
			"version": "7.9.6",
			"resolved": "https://registry.npmjs.org/@babel/helper-replace-supers/-/helper-replace-supers-7.9.6.tgz",
			"integrity": "sha512-qX+chbxkbArLyCImk3bWV+jB5gTNU/rsze+JlcF6Nf8tVTigPJSI1o1oBow/9Resa1yehUO9lIipsmu9oG4RzA==",
			"dev": true,
			"requires": {
			  "@babel/helper-member-expression-to-functions": "^7.8.3",
			  "@babel/helper-optimise-call-expression": "^7.8.3",
			  "@babel/traverse": "^7.9.6",
			  "@babel/types": "^7.9.6"
			}
		  },
		  "@babel/helper-simple-access": {
			"version": "7.8.3",
			"resolved": "https://registry.npmjs.org/@babel/helper-simple-access/-/helper-simple-access-7.8.3.tgz",
			"integrity": "sha512-VNGUDjx5cCWg4vvCTR8qQ7YJYZ+HBjxOgXEl7ounz+4Sn7+LMD3CFrCTEU6/qXKbA2nKg21CwhhBzO0RpRbdCw==",
			"dev": true,
			"requires": {
			  "@babel/template": "^7.8.3",
			  "@babel/types": "^7.8.3"
			}
		  },
		  "@babel/helper-split-export-declaration": {
			"version": "7.8.3",
			"resolved": "https://registry.npmjs.org/@babel/helper-split-export-declaration/-/helper-split-export-declaration-7.8.3.tgz",
			"integrity": "sha512-3x3yOeyBhW851hroze7ElzdkeRXQYQbFIb7gLK1WQYsw2GWDay5gAJNw1sWJ0VFP6z5J1whqeXH/WCdCjZv6dA==",
			"dev": true,
			"requires": {
			  "@babel/types": "^7.8.3"
			}
		  },
		  "@babel/helper-validator-identifier": {
			"version": "7.9.5",
			"resolved": "https://registry.npmjs.org/@babel/helper-validator-identifier/-/helper-validator-identifier-7.9.5.tgz",
			"integrity": "sha512-/8arLKUFq882w4tWGj9JYzRpAlZgiWUJ+dtteNTDqrRBz9Iguck9Rn3ykuBDoUwh2TO4tSAJlrxDUOXWklJe4g==",
			"dev": true
		  },
		  "@babel/helpers": {
			"version": "7.9.6",
			"resolved": "https://registry.npmjs.org/@babel/helpers/-/helpers-7.9.6.tgz",
			"integrity": "sha512-tI4bUbldloLcHWoRUMAj4g1bF313M/o6fBKhIsb3QnGVPwRm9JsNf/gqMkQ7zjqReABiffPV6RWj7hEglID5Iw==",
			"dev": true,
			"requires": {
			  "@babel/template": "^7.8.3",
			  "@babel/traverse": "^7.9.6",
			  "@babel/types": "^7.9.6"
			}
		  },
		  "@babel/highlight": {
			"version": "7.9.0",
			"resolved": "https://registry.npmjs.org/@babel/highlight/-/highlight-7.9.0.tgz",
			"integrity": "sha512-lJZPilxX7Op3Nv/2cvFdnlepPXDxi29wxteT57Q965oc5R9v86ztx0jfxVrTcBk8C2kcPkkDa2Z4T3ZsPPVWsQ==",
			"dev": true,
			"requires": {
			  "@babel/helper-validator-identifier": "^7.9.0",
			  "chalk": "^2.0.0",
			  "js-tokens": "^4.0.0"
			},
			"dependencies": {
			  "js-tokens": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/js-tokens/-/js-tokens-4.0.0.tgz",
				"integrity": "sha512-RdJUflcE3cUzKiMqQgsCu06FPu9UdIJO0beYbPhHN4k6apgJtifcoCtT9bcxOpYBtpD2kCM6Sbzg4CausW/PKQ==",
				"dev": true
			  }
			}
		  },
		  "@babel/parser": {
			"version": "7.9.6",
			"resolved": "https://registry.npmjs.org/@babel/parser/-/parser-7.9.6.tgz",
			"integrity": "sha512-AoeIEJn8vt+d/6+PXDRPaksYhnlbMIiejioBZvvMQsOjW/JYK6k/0dKnvvP3EhK5GfMBWDPtrxRtegWdAcdq9Q==",
			"dev": true
		  },
		  "@babel/plugin-syntax-async-generators": {
			"version": "7.8.4",
			"resolved": "https://registry.npmjs.org/@babel/plugin-syntax-async-generators/-/plugin-syntax-async-generators-7.8.4.tgz",
			"integrity": "sha512-tycmZxkGfZaxhMRbXlPXuVFpdWlXpir2W4AMhSJgRKzk/eDlIXOhb2LHWoLpDF7TEHylV5zNhykX6KAgHJmTNw==",
			"dev": true,
			"requires": {
			  "@babel/helper-plugin-utils": "^7.8.0"
			}
		  },
		  "@babel/plugin-syntax-bigint": {
			"version": "7.8.3",
			"resolved": "https://registry.npmjs.org/@babel/plugin-syntax-bigint/-/plugin-syntax-bigint-7.8.3.tgz",
			"integrity": "sha512-wnTnFlG+YxQm3vDxpGE57Pj0srRU4sHE/mDkt1qv2YJJSeUAec2ma4WLUnUPeKjyrfntVwe/N6dCXpU+zL3Npg==",
			"dev": true,
			"requires": {
			  "@babel/helper-plugin-utils": "^7.8.0"
			}
		  },
		  "@babel/plugin-syntax-class-properties": {
			"version": "7.8.3",
			"resolved": "https://registry.npmjs.org/@babel/plugin-syntax-class-properties/-/plugin-syntax-class-properties-7.8.3.tgz",
			"integrity": "sha512-UcAyQWg2bAN647Q+O811tG9MrJ38Z10jjhQdKNAL8fsyPzE3cCN/uT+f55cFVY4aGO4jqJAvmqsuY3GQDwAoXg==",
			"dev": true,
			"requires": {
			  "@babel/helper-plugin-utils": "^7.8.3"
			}
		  },
		  "@babel/plugin-syntax-json-strings": {
			"version": "7.8.3",
			"resolved": "https://registry.npmjs.org/@babel/plugin-syntax-json-strings/-/plugin-syntax-json-strings-7.8.3.tgz",
			"integrity": "sha512-lY6kdGpWHvjoe2vk4WrAapEuBR69EMxZl+RoGRhrFGNYVK8mOPAW8VfbT/ZgrFbXlDNiiaxQnAtgVCZ6jv30EA==",
			"dev": true,
			"requires": {
			  "@babel/helper-plugin-utils": "^7.8.0"
			}
		  },
		  "@babel/plugin-syntax-logical-assignment-operators": {
			"version": "7.8.3",
			"resolved": "https://registry.npmjs.org/@babel/plugin-syntax-logical-assignment-operators/-/plugin-syntax-logical-assignment-operators-7.8.3.tgz",
			"integrity": "sha512-Zpg2Sgc++37kuFl6ppq2Q7Awc6E6AIW671x5PY8E/f7MCIyPPGK/EoeZXvvY3P42exZ3Q4/t3YOzP/HiN79jDg==",
			"dev": true,
			"requires": {
			  "@babel/helper-plugin-utils": "^7.8.3"
			}
		  },
		  "@babel/plugin-syntax-nullish-coalescing-operator": {
			"version": "7.8.3",
			"resolved": "https://registry.npmjs.org/@babel/plugin-syntax-nullish-coalescing-operator/-/plugin-syntax-nullish-coalescing-operator-7.8.3.tgz",
			"integrity": "sha512-aSff4zPII1u2QD7y+F8oDsz19ew4IGEJg9SVW+bqwpwtfFleiQDMdzA/R+UlWDzfnHFCxxleFT0PMIrR36XLNQ==",
			"dev": true,
			"requires": {
			  "@babel/helper-plugin-utils": "^7.8.0"
			}
		  },
		  "@babel/plugin-syntax-numeric-separator": {
			"version": "7.8.3",
			"resolved": "https://registry.npmjs.org/@babel/plugin-syntax-numeric-separator/-/plugin-syntax-numeric-separator-7.8.3.tgz",
			"integrity": "sha512-H7dCMAdN83PcCmqmkHB5dtp+Xa9a6LKSvA2hiFBC/5alSHxM5VgWZXFqDi0YFe8XNGT6iCa+z4V4zSt/PdZ7Dw==",
			"dev": true,
			"requires": {
			  "@babel/helper-plugin-utils": "^7.8.3"
			}
		  },
		  "@babel/plugin-syntax-object-rest-spread": {
			"version": "7.8.3",
			"resolved": "https://registry.npmjs.org/@babel/plugin-syntax-object-rest-spread/-/plugin-syntax-object-rest-spread-7.8.3.tgz",
			"integrity": "sha512-XoqMijGZb9y3y2XskN+P1wUGiVwWZ5JmoDRwx5+3GmEplNyVM2s2Dg8ILFQm8rWM48orGy5YpI5Bl8U1y7ydlA==",
			"dev": true,
			"requires": {
			  "@babel/helper-plugin-utils": "^7.8.0"
			}
		  },
		  "@babel/plugin-syntax-optional-catch-binding": {
			"version": "7.8.3",
			"resolved": "https://registry.npmjs.org/@babel/plugin-syntax-optional-catch-binding/-/plugin-syntax-optional-catch-binding-7.8.3.tgz",
			"integrity": "sha512-6VPD0Pc1lpTqw0aKoeRTMiB+kWhAoT24PA+ksWSBrFtl5SIRVpZlwN3NNPQjehA2E/91FV3RjLWoVTglWcSV3Q==",
			"dev": true,
			"requires": {
			  "@babel/helper-plugin-utils": "^7.8.0"
			}
		  },
		  "@babel/plugin-syntax-optional-chaining": {
			"version": "7.8.3",
			"resolved": "https://registry.npmjs.org/@babel/plugin-syntax-optional-chaining/-/plugin-syntax-optional-chaining-7.8.3.tgz",
			"integrity": "sha512-KoK9ErH1MBlCPxV0VANkXW2/dw4vlbGDrFgz8bmUsBGYkFRcbRwMh6cIJubdPrkxRwuGdtCk0v/wPTKbQgBjkg==",
			"dev": true,
			"requires": {
			  "@babel/helper-plugin-utils": "^7.8.0"
			}
		  },
		  "@babel/template": {
			"version": "7.8.6",
			"resolved": "https://registry.npmjs.org/@babel/template/-/template-7.8.6.tgz",
			"integrity": "sha512-zbMsPMy/v0PWFZEhQJ66bqjhH+z0JgMoBWuikXybgG3Gkd/3t5oQ1Rw2WQhnSrsOmsKXnZOx15tkC4qON/+JPg==",
			"dev": true,
			"requires": {
			  "@babel/code-frame": "^7.8.3",
			  "@babel/parser": "^7.8.6",
			  "@babel/types": "^7.8.6"
			}
		  },
		  "@babel/traverse": {
			"version": "7.9.6",
			"resolved": "https://registry.npmjs.org/@babel/traverse/-/traverse-7.9.6.tgz",
			"integrity": "sha512-b3rAHSjbxy6VEAvlxM8OV/0X4XrG72zoxme6q1MOoe2vd0bEc+TwayhuC1+Dfgqh1QEG+pj7atQqvUprHIccsg==",
			"dev": true,
			"requires": {
			  "@babel/code-frame": "^7.8.3",
			  "@babel/generator": "^7.9.6",
			  "@babel/helper-function-name": "^7.9.5",
			  "@babel/helper-split-export-declaration": "^7.8.3",
			  "@babel/parser": "^7.9.6",
			  "@babel/types": "^7.9.6",
			  "debug": "^4.1.0",
			  "globals": "^11.1.0",
			  "lodash": "^4.17.13"
			},
			"dependencies": {
			  "debug": {
				"version": "4.1.1",
				"resolved": "https://registry.npmjs.org/debug/-/debug-4.1.1.tgz",
				"integrity": "sha512-pYAIzeRo8J6KPEaJ0VWOh5Pzkbw/RetuzehGM7QRRX5he4fPHx2rdKMB256ehJCkX+XRQm16eZLqLNS8RSZXZw==",
				"dev": true,
				"requires": {
				  "ms": "^2.1.1"
				}
			  },
			  "globals": {
				"version": "11.12.0",
				"resolved": "https://registry.npmjs.org/globals/-/globals-11.12.0.tgz",
				"integrity": "sha512-WOBp/EEGUiIsJSp7wcv/y6MO+lV9UoncWqxuFfm8eBwzWNgyfBd6Gz+IeKQ9jCmyhoH99g15M3T+QaVHFjizVA==",
				"dev": true
			  },
			  "ms": {
				"version": "2.1.2",
				"resolved": "https://registry.npmjs.org/ms/-/ms-2.1.2.tgz",
				"integrity": "sha512-sGkPx+VjMtmA6MX27oA4FBFELFCZZ4S4XqeGOXCv68tT+jb3vk/RyaKWP0PTKyWtmLSM0b+adUTEvbs1PEaH2w==",
				"dev": true
			  }
			}
		  },
		  "@babel/types": {
			"version": "7.9.6",
			"resolved": "https://registry.npmjs.org/@babel/types/-/types-7.9.6.tgz",
			"integrity": "sha512-qxXzvBO//jO9ZnoasKF1uJzHd2+M6Q2ZPIVfnFps8JJvXy0ZBbwbNOmE6SGIY5XOY6d1Bo5lb9d9RJ8nv3WSeA==",
			"dev": true,
			"requires": {
			  "@babel/helper-validator-identifier": "^7.9.5",
			  "lodash": "^4.17.13",
			  "to-fast-properties": "^2.0.0"
			}
		  },
		  "@bcoe/v8-coverage": {
			"version": "0.2.3",
			"resolved": "https://registry.npmjs.org/@bcoe/v8-coverage/-/v8-coverage-0.2.3.tgz",
			"integrity": "sha512-0hYQ8SB4Db5zvZB4axdMHGwEaQjkZzFjQiN9LVYvIFB2nSUHW9tYpxWriPrWDASIxiaXax83REcLxuSdnGPZtw==",
			"dev": true
		  },
		  "@cnakazawa/watch": {
			"version": "1.0.4",
			"resolved": "https://registry.npmjs.org/@cnakazawa/watch/-/watch-1.0.4.tgz",
			"integrity": "sha512-v9kIhKwjeZThiWrLmj0y17CWoyddASLj9O2yvbZkbvw/N3rWOYy9zkV66ursAoVr0mV15bL8g0c4QZUE6cdDoQ==",
			"dev": true,
			"requires": {
			  "exec-sh": "^0.3.2",
			  "minimist": "^1.2.0"
			}
		  },
		  "@istanbuljs/load-nyc-config": {
			"version": "1.0.0",
			"resolved": "https://registry.npmjs.org/@istanbuljs/load-nyc-config/-/load-nyc-config-1.0.0.tgz",
			"integrity": "sha512-ZR0rq/f/E4f4XcgnDvtMWXCUJpi8eO0rssVhmztsZqLIEFA9UUP9zmpE0VxlM+kv/E1ul2I876Fwil2ayptDVg==",
			"dev": true,
			"requires": {
			  "camelcase": "^5.3.1",
			  "find-up": "^4.1.0",
			  "js-yaml": "^3.13.1",
			  "resolve-from": "^5.0.0"
			},
			"dependencies": {
			  "find-up": {
				"version": "4.1.0",
				"resolved": "https://registry.npmjs.org/find-up/-/find-up-4.1.0.tgz",
				"integrity": "sha512-PpOwAdQ/YlXQ2vj8a3h8IipDuYRi3wceVQQGYWxNINccq40Anw7BlsEXCMbt1Zt+OLA6Fq9suIpIWD0OsnISlw==",
				"dev": true,
				"requires": {
				  "locate-path": "^5.0.0",
				  "path-exists": "^4.0.0"
				}
			  },
			  "locate-path": {
				"version": "5.0.0",
				"resolved": "https://registry.npmjs.org/locate-path/-/locate-path-5.0.0.tgz",
				"integrity": "sha512-t7hw9pI+WvuwNJXwk5zVHpyhIqzg2qTlklJOf0mVxGSbe3Fp2VieZcduNYjaLDoy6p9uGpQEGWG87WpMKlNq8g==",
				"dev": true,
				"requires": {
				  "p-locate": "^4.1.0"
				}
			  },
			  "p-locate": {
				"version": "4.1.0",
				"resolved": "https://registry.npmjs.org/p-locate/-/p-locate-4.1.0.tgz",
				"integrity": "sha512-R79ZZ/0wAxKGu3oYMlz8jy/kbhsNrS7SKZ7PxEHBgJ5+F2mtFW2fK2cOtBh1cHYkQsbzFV7I+EoRKe6Yt0oK7A==",
				"dev": true,
				"requires": {
				  "p-limit": "^2.2.0"
				}
			  },
			  "path-exists": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/path-exists/-/path-exists-4.0.0.tgz",
				"integrity": "sha512-ak9Qy5Q7jYb2Wwcey5Fpvg2KoAc/ZIhLSLOSBmRmygPsGwkVVt0fZa0qrtMz+m6tJTAHfZQ8FnmB4MG4LWy7/w==",
				"dev": true
			  },
			  "resolve-from": {
				"version": "5.0.0",
				"resolved": "https://registry.npmjs.org/resolve-from/-/resolve-from-5.0.0.tgz",
				"integrity": "sha512-qYg9KP24dD5qka9J47d0aVky0N+b4fTU89LN9iDnjB5waksiC49rvMB0PrUJQGoTmH50XPiqOvAjDfaijGxYZw==",
				"dev": true
			  }
			}
		  },
		  "@istanbuljs/schema": {
			"version": "0.1.2",
			"resolved": "https://registry.npmjs.org/@istanbuljs/schema/-/schema-0.1.2.tgz",
			"integrity": "sha512-tsAQNx32a8CoFhjhijUIhI4kccIAgmGhy8LZMZgGfmXcpMbPRUqn5LWmgRttILi6yeGmBJd2xsPkFMs0PzgPCw==",
			"dev": true
		  },
		  "@jest/console": {
			"version": "26.0.1",
			"resolved": "https://registry.npmjs.org/@jest/console/-/console-26.0.1.tgz",
			"integrity": "sha512-9t1KUe/93coV1rBSxMmBAOIK3/HVpwxArCA1CxskKyRiv6o8J70V8C/V3OJminVCTa2M0hQI9AWRd5wxu2dAHw==",
			"dev": true,
			"requires": {
			  "@jest/types": "^26.0.1",
			  "chalk": "^4.0.0",
			  "jest-message-util": "^26.0.1",
			  "jest-util": "^26.0.1",
			  "slash": "^3.0.0"
			},
			"dependencies": {
			  "@jest/types": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/@jest/types/-/types-26.0.1.tgz",
				"integrity": "sha512-IbtjvqI9+eS1qFnOIEL7ggWmT+iK/U+Vde9cGWtYb/b6XgKb3X44ZAe/z9YZzoAAZ/E92m0DqrilF934IGNnQA==",
				"dev": true,
				"requires": {
				  "@types/istanbul-lib-coverage": "^2.0.0",
				  "@types/istanbul-reports": "^1.1.1",
				  "@types/yargs": "^15.0.0",
				  "chalk": "^4.0.0"
				}
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-4.0.0.tgz",
				"integrity": "sha512-N9oWFcegS0sFr9oh1oz2d7Npos6vNoWW9HvtCg5N1KRFpUhaAhvTv5Y58g880fZaEYSNm3qDz8SU1UrGvp+n7A==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "@jest/core": {
			"version": "26.0.1",
			"resolved": "https://registry.npmjs.org/@jest/core/-/core-26.0.1.tgz",
			"integrity": "sha512-Xq3eqYnxsG9SjDC+WLeIgf7/8KU6rddBxH+SCt18gEpOhAGYC/Mq+YbtlNcIdwjnnT+wDseXSbU0e5X84Y4jTQ==",
			"dev": true,
			"requires": {
			  "@jest/console": "^26.0.1",
			  "@jest/reporters": "^26.0.1",
			  "@jest/test-result": "^26.0.1",
			  "@jest/transform": "^26.0.1",
			  "@jest/types": "^26.0.1",
			  "ansi-escapes": "^4.2.1",
			  "chalk": "^4.0.0",
			  "exit": "^0.1.2",
			  "graceful-fs": "^4.2.4",
			  "jest-changed-files": "^26.0.1",
			  "jest-config": "^26.0.1",
			  "jest-haste-map": "^26.0.1",
			  "jest-message-util": "^26.0.1",
			  "jest-regex-util": "^26.0.0",
			  "jest-resolve": "^26.0.1",
			  "jest-resolve-dependencies": "^26.0.1",
			  "jest-runner": "^26.0.1",
			  "jest-runtime": "^26.0.1",
			  "jest-snapshot": "^26.0.1",
			  "jest-util": "^26.0.1",
			  "jest-validate": "^26.0.1",
			  "jest-watcher": "^26.0.1",
			  "micromatch": "^4.0.2",
			  "p-each-series": "^2.1.0",
			  "rimraf": "^3.0.0",
			  "slash": "^3.0.0",
			  "strip-ansi": "^6.0.0"
			},
			"dependencies": {
			  "@jest/types": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/@jest/types/-/types-26.0.1.tgz",
				"integrity": "sha512-IbtjvqI9+eS1qFnOIEL7ggWmT+iK/U+Vde9cGWtYb/b6XgKb3X44ZAe/z9YZzoAAZ/E92m0DqrilF934IGNnQA==",
				"dev": true,
				"requires": {
				  "@types/istanbul-lib-coverage": "^2.0.0",
				  "@types/istanbul-reports": "^1.1.1",
				  "@types/yargs": "^15.0.0",
				  "chalk": "^4.0.0"
				}
			  },
			  "ansi-regex": {
				"version": "5.0.0",
				"resolved": "https://registry.npmjs.org/ansi-regex/-/ansi-regex-5.0.0.tgz",
				"integrity": "sha512-bY6fj56OUQ0hU1KjFNDQuJFezqKdrAyFdIevADiqrWHwSlbmBNMHp5ak2f40Pm8JTFyM2mqxkG6ngkHO11f/lg==",
				"dev": true
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-4.0.0.tgz",
				"integrity": "sha512-N9oWFcegS0sFr9oh1oz2d7Npos6vNoWW9HvtCg5N1KRFpUhaAhvTv5Y58g880fZaEYSNm3qDz8SU1UrGvp+n7A==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "graceful-fs": {
				"version": "4.2.4",
				"resolved": "https://registry.npmjs.org/graceful-fs/-/graceful-fs-4.2.4.tgz",
				"integrity": "sha512-WjKPNJF79dtJAVniUlGGWHYGz2jWxT6VhN/4m1NdkbZ2nOsEF+cI1Edgql5zCRhs/VsQYRvrXctxktVXZUkixw==",
				"dev": true
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "micromatch": {
				"version": "4.0.2",
				"resolved": "https://registry.npmjs.org/micromatch/-/micromatch-4.0.2.tgz",
				"integrity": "sha512-y7FpHSbMUMoyPbYUSzO6PaZ6FyRnQOpHuKwbo1G+Knck95XVU4QAiKdGEnj5wwoS7PlOgthX/09u5iFJ+aYf5Q==",
				"dev": true,
				"requires": {
				  "braces": "^3.0.1",
				  "picomatch": "^2.0.5"
				}
			  },
			  "strip-ansi": {
				"version": "6.0.0",
				"resolved": "https://registry.npmjs.org/strip-ansi/-/strip-ansi-6.0.0.tgz",
				"integrity": "sha512-AuvKTrTfQNYNIctbR1K/YGTR1756GycPsg7b9bdV9Duqur4gv6aKqHXah67Z8ImS7WEz5QVcOtlfW2rZEugt6w==",
				"dev": true,
				"requires": {
				  "ansi-regex": "^5.0.0"
				}
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "@jest/environment": {
			"version": "26.0.1",
			"resolved": "https://registry.npmjs.org/@jest/environment/-/environment-26.0.1.tgz",
			"integrity": "sha512-xBDxPe8/nx251u0VJ2dFAFz2H23Y98qdIaNwnMK6dFQr05jc+Ne/2np73lOAx+5mSBO/yuQldRrQOf6hP1h92g==",
			"dev": true,
			"requires": {
			  "@jest/fake-timers": "^26.0.1",
			  "@jest/types": "^26.0.1",
			  "jest-mock": "^26.0.1"
			},
			"dependencies": {
			  "@jest/types": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/@jest/types/-/types-26.0.1.tgz",
				"integrity": "sha512-IbtjvqI9+eS1qFnOIEL7ggWmT+iK/U+Vde9cGWtYb/b6XgKb3X44ZAe/z9YZzoAAZ/E92m0DqrilF934IGNnQA==",
				"dev": true,
				"requires": {
				  "@types/istanbul-lib-coverage": "^2.0.0",
				  "@types/istanbul-reports": "^1.1.1",
				  "@types/yargs": "^15.0.0",
				  "chalk": "^4.0.0"
				}
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-4.0.0.tgz",
				"integrity": "sha512-N9oWFcegS0sFr9oh1oz2d7Npos6vNoWW9HvtCg5N1KRFpUhaAhvTv5Y58g880fZaEYSNm3qDz8SU1UrGvp+n7A==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "@jest/fake-timers": {
			"version": "26.0.1",
			"resolved": "https://registry.npmjs.org/@jest/fake-timers/-/fake-timers-26.0.1.tgz",
			"integrity": "sha512-Oj/kCBnTKhm7CR+OJSjZty6N1bRDr9pgiYQr4wY221azLz5PHi08x/U+9+QpceAYOWheauLP8MhtSVFrqXQfhg==",
			"dev": true,
			"requires": {
			  "@jest/types": "^26.0.1",
			  "@sinonjs/fake-timers": "^6.0.1",
			  "jest-message-util": "^26.0.1",
			  "jest-mock": "^26.0.1",
			  "jest-util": "^26.0.1"
			},
			"dependencies": {
			  "@jest/types": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/@jest/types/-/types-26.0.1.tgz",
				"integrity": "sha512-IbtjvqI9+eS1qFnOIEL7ggWmT+iK/U+Vde9cGWtYb/b6XgKb3X44ZAe/z9YZzoAAZ/E92m0DqrilF934IGNnQA==",
				"dev": true,
				"requires": {
				  "@types/istanbul-lib-coverage": "^2.0.0",
				  "@types/istanbul-reports": "^1.1.1",
				  "@types/yargs": "^15.0.0",
				  "chalk": "^4.0.0"
				}
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-4.0.0.tgz",
				"integrity": "sha512-N9oWFcegS0sFr9oh1oz2d7Npos6vNoWW9HvtCg5N1KRFpUhaAhvTv5Y58g880fZaEYSNm3qDz8SU1UrGvp+n7A==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "@jest/globals": {
			"version": "26.0.1",
			"resolved": "https://registry.npmjs.org/@jest/globals/-/globals-26.0.1.tgz",
			"integrity": "sha512-iuucxOYB7BRCvT+TYBzUqUNuxFX1hqaR6G6IcGgEqkJ5x4htNKo1r7jk1ji9Zj8ZMiMw0oB5NaA7k5Tx6MVssA==",
			"dev": true,
			"requires": {
			  "@jest/environment": "^26.0.1",
			  "@jest/types": "^26.0.1",
			  "expect": "^26.0.1"
			},
			"dependencies": {
			  "@jest/types": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/@jest/types/-/types-26.0.1.tgz",
				"integrity": "sha512-IbtjvqI9+eS1qFnOIEL7ggWmT+iK/U+Vde9cGWtYb/b6XgKb3X44ZAe/z9YZzoAAZ/E92m0DqrilF934IGNnQA==",
				"dev": true,
				"requires": {
				  "@types/istanbul-lib-coverage": "^2.0.0",
				  "@types/istanbul-reports": "^1.1.1",
				  "@types/yargs": "^15.0.0",
				  "chalk": "^4.0.0"
				}
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-4.0.0.tgz",
				"integrity": "sha512-N9oWFcegS0sFr9oh1oz2d7Npos6vNoWW9HvtCg5N1KRFpUhaAhvTv5Y58g880fZaEYSNm3qDz8SU1UrGvp+n7A==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "@jest/reporters": {
			"version": "26.0.1",
			"resolved": "https://registry.npmjs.org/@jest/reporters/-/reporters-26.0.1.tgz",
			"integrity": "sha512-NWWy9KwRtE1iyG/m7huiFVF9YsYv/e+mbflKRV84WDoJfBqUrNRyDbL/vFxQcYLl8IRqI4P3MgPn386x76Gf2g==",
			"dev": true,
			"requires": {
			  "@bcoe/v8-coverage": "^0.2.3",
			  "@jest/console": "^26.0.1",
			  "@jest/test-result": "^26.0.1",
			  "@jest/transform": "^26.0.1",
			  "@jest/types": "^26.0.1",
			  "chalk": "^4.0.0",
			  "collect-v8-coverage": "^1.0.0",
			  "exit": "^0.1.2",
			  "glob": "^7.1.2",
			  "graceful-fs": "^4.2.4",
			  "istanbul-lib-coverage": "^3.0.0",
			  "istanbul-lib-instrument": "^4.0.0",
			  "istanbul-lib-report": "^3.0.0",
			  "istanbul-lib-source-maps": "^4.0.0",
			  "istanbul-reports": "^3.0.2",
			  "jest-haste-map": "^26.0.1",
			  "jest-resolve": "^26.0.1",
			  "jest-util": "^26.0.1",
			  "jest-worker": "^26.0.0",
			  "node-notifier": "^7.0.0",
			  "slash": "^3.0.0",
			  "source-map": "^0.6.0",
			  "string-length": "^4.0.1",
			  "terminal-link": "^2.0.0",
			  "v8-to-istanbul": "^4.1.3"
			},
			"dependencies": {
			  "@jest/types": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/@jest/types/-/types-26.0.1.tgz",
				"integrity": "sha512-IbtjvqI9+eS1qFnOIEL7ggWmT+iK/U+Vde9cGWtYb/b6XgKb3X44ZAe/z9YZzoAAZ/E92m0DqrilF934IGNnQA==",
				"dev": true,
				"requires": {
				  "@types/istanbul-lib-coverage": "^2.0.0",
				  "@types/istanbul-reports": "^1.1.1",
				  "@types/yargs": "^15.0.0",
				  "chalk": "^4.0.0"
				}
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-4.0.0.tgz",
				"integrity": "sha512-N9oWFcegS0sFr9oh1oz2d7Npos6vNoWW9HvtCg5N1KRFpUhaAhvTv5Y58g880fZaEYSNm3qDz8SU1UrGvp+n7A==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "graceful-fs": {
				"version": "4.2.4",
				"resolved": "https://registry.npmjs.org/graceful-fs/-/graceful-fs-4.2.4.tgz",
				"integrity": "sha512-WjKPNJF79dtJAVniUlGGWHYGz2jWxT6VhN/4m1NdkbZ2nOsEF+cI1Edgql5zCRhs/VsQYRvrXctxktVXZUkixw==",
				"dev": true
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "source-map": {
				"version": "0.6.1",
				"resolved": "https://registry.npmjs.org/source-map/-/source-map-0.6.1.tgz",
				"integrity": "sha512-UjgapumWlbMhkBgzT7Ykc5YXUT46F0iKu8SGXq0bcwP5dz/h0Plj6enJqjz1Zbq2l5WaqYnrVbwWOWMyF3F47g==",
				"dev": true
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "@jest/source-map": {
			"version": "26.0.0",
			"resolved": "https://registry.npmjs.org/@jest/source-map/-/source-map-26.0.0.tgz",
			"integrity": "sha512-S2Z+Aj/7KOSU2TfW0dyzBze7xr95bkm5YXNUqqCek+HE0VbNNSNzrRwfIi5lf7wvzDTSS0/ib8XQ1krFNyYgbQ==",
			"dev": true,
			"requires": {
			  "callsites": "^3.0.0",
			  "graceful-fs": "^4.2.4",
			  "source-map": "^0.6.0"
			},
			"dependencies": {
			  "graceful-fs": {
				"version": "4.2.4",
				"resolved": "https://registry.npmjs.org/graceful-fs/-/graceful-fs-4.2.4.tgz",
				"integrity": "sha512-WjKPNJF79dtJAVniUlGGWHYGz2jWxT6VhN/4m1NdkbZ2nOsEF+cI1Edgql5zCRhs/VsQYRvrXctxktVXZUkixw==",
				"dev": true
			  },
			  "source-map": {
				"version": "0.6.1",
				"resolved": "https://registry.npmjs.org/source-map/-/source-map-0.6.1.tgz",
				"integrity": "sha512-UjgapumWlbMhkBgzT7Ykc5YXUT46F0iKu8SGXq0bcwP5dz/h0Plj6enJqjz1Zbq2l5WaqYnrVbwWOWMyF3F47g==",
				"dev": true
			  }
			}
		  },
		  "@jest/test-result": {
			"version": "26.0.1",
			"resolved": "https://registry.npmjs.org/@jest/test-result/-/test-result-26.0.1.tgz",
			"integrity": "sha512-oKwHvOI73ICSYRPe8WwyYPTtiuOAkLSbY8/MfWF3qDEd/sa8EDyZzin3BaXTqufir/O/Gzea4E8Zl14XU4Mlyg==",
			"dev": true,
			"requires": {
			  "@jest/console": "^26.0.1",
			  "@jest/types": "^26.0.1",
			  "@types/istanbul-lib-coverage": "^2.0.0",
			  "collect-v8-coverage": "^1.0.0"
			},
			"dependencies": {
			  "@jest/types": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/@jest/types/-/types-26.0.1.tgz",
				"integrity": "sha512-IbtjvqI9+eS1qFnOIEL7ggWmT+iK/U+Vde9cGWtYb/b6XgKb3X44ZAe/z9YZzoAAZ/E92m0DqrilF934IGNnQA==",
				"dev": true,
				"requires": {
				  "@types/istanbul-lib-coverage": "^2.0.0",
				  "@types/istanbul-reports": "^1.1.1",
				  "@types/yargs": "^15.0.0",
				  "chalk": "^4.0.0"
				}
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-4.0.0.tgz",
				"integrity": "sha512-N9oWFcegS0sFr9oh1oz2d7Npos6vNoWW9HvtCg5N1KRFpUhaAhvTv5Y58g880fZaEYSNm3qDz8SU1UrGvp+n7A==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "@jest/test-sequencer": {
			"version": "26.0.1",
			"resolved": "https://registry.npmjs.org/@jest/test-sequencer/-/test-sequencer-26.0.1.tgz",
			"integrity": "sha512-ssga8XlwfP8YjbDcmVhwNlrmblddMfgUeAkWIXts1V22equp2GMIHxm7cyeD5Q/B0ZgKPK/tngt45sH99yLLGg==",
			"dev": true,
			"requires": {
			  "@jest/test-result": "^26.0.1",
			  "graceful-fs": "^4.2.4",
			  "jest-haste-map": "^26.0.1",
			  "jest-runner": "^26.0.1",
			  "jest-runtime": "^26.0.1"
			},
			"dependencies": {
			  "graceful-fs": {
				"version": "4.2.4",
				"resolved": "https://registry.npmjs.org/graceful-fs/-/graceful-fs-4.2.4.tgz",
				"integrity": "sha512-WjKPNJF79dtJAVniUlGGWHYGz2jWxT6VhN/4m1NdkbZ2nOsEF+cI1Edgql5zCRhs/VsQYRvrXctxktVXZUkixw==",
				"dev": true
			  }
			}
		  },
		  "@jest/transform": {
			"version": "26.0.1",
			"resolved": "https://registry.npmjs.org/@jest/transform/-/transform-26.0.1.tgz",
			"integrity": "sha512-pPRkVkAQ91drKGbzCfDOoHN838+FSbYaEAvBXvKuWeeRRUD8FjwXkqfUNUZL6Ke48aA/1cqq/Ni7kVMCoqagWA==",
			"dev": true,
			"requires": {
			  "@babel/core": "^7.1.0",
			  "@jest/types": "^26.0.1",
			  "babel-plugin-istanbul": "^6.0.0",
			  "chalk": "^4.0.0",
			  "convert-source-map": "^1.4.0",
			  "fast-json-stable-stringify": "^2.0.0",
			  "graceful-fs": "^4.2.4",
			  "jest-haste-map": "^26.0.1",
			  "jest-regex-util": "^26.0.0",
			  "jest-util": "^26.0.1",
			  "micromatch": "^4.0.2",
			  "pirates": "^4.0.1",
			  "slash": "^3.0.0",
			  "source-map": "^0.6.1",
			  "write-file-atomic": "^3.0.0"
			},
			"dependencies": {
			  "@jest/types": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/@jest/types/-/types-26.0.1.tgz",
				"integrity": "sha512-IbtjvqI9+eS1qFnOIEL7ggWmT+iK/U+Vde9cGWtYb/b6XgKb3X44ZAe/z9YZzoAAZ/E92m0DqrilF934IGNnQA==",
				"dev": true,
				"requires": {
				  "@types/istanbul-lib-coverage": "^2.0.0",
				  "@types/istanbul-reports": "^1.1.1",
				  "@types/yargs": "^15.0.0",
				  "chalk": "^4.0.0"
				}
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-4.0.0.tgz",
				"integrity": "sha512-N9oWFcegS0sFr9oh1oz2d7Npos6vNoWW9HvtCg5N1KRFpUhaAhvTv5Y58g880fZaEYSNm3qDz8SU1UrGvp+n7A==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "graceful-fs": {
				"version": "4.2.4",
				"resolved": "https://registry.npmjs.org/graceful-fs/-/graceful-fs-4.2.4.tgz",
				"integrity": "sha512-WjKPNJF79dtJAVniUlGGWHYGz2jWxT6VhN/4m1NdkbZ2nOsEF+cI1Edgql5zCRhs/VsQYRvrXctxktVXZUkixw==",
				"dev": true
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "micromatch": {
				"version": "4.0.2",
				"resolved": "https://registry.npmjs.org/micromatch/-/micromatch-4.0.2.tgz",
				"integrity": "sha512-y7FpHSbMUMoyPbYUSzO6PaZ6FyRnQOpHuKwbo1G+Knck95XVU4QAiKdGEnj5wwoS7PlOgthX/09u5iFJ+aYf5Q==",
				"dev": true,
				"requires": {
				  "braces": "^3.0.1",
				  "picomatch": "^2.0.5"
				}
			  },
			  "source-map": {
				"version": "0.6.1",
				"resolved": "https://registry.npmjs.org/source-map/-/source-map-0.6.1.tgz",
				"integrity": "sha512-UjgapumWlbMhkBgzT7Ykc5YXUT46F0iKu8SGXq0bcwP5dz/h0Plj6enJqjz1Zbq2l5WaqYnrVbwWOWMyF3F47g==",
				"dev": true
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "@jest/types": {
			"version": "25.5.0",
			"resolved": "https://registry.npmjs.org/@jest/types/-/types-25.5.0.tgz",
			"integrity": "sha512-OXD0RgQ86Tu3MazKo8bnrkDRaDXXMGUqd+kTtLtK1Zb7CRzQcaSRPPPV37SvYTdevXEBVxe0HXylEjs8ibkmCw==",
			"dev": true,
			"requires": {
			  "@types/istanbul-lib-coverage": "^2.0.0",
			  "@types/istanbul-reports": "^1.1.1",
			  "@types/yargs": "^15.0.0",
			  "chalk": "^3.0.0"
			},
			"dependencies": {
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "3.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-3.0.0.tgz",
				"integrity": "sha512-4D3B6Wf41KOYRFdszmDqMCGq5VV/uMAB273JILmO+3jAlh8X4qDtdtgCR3fxtbLEMzSx22QdhnDcJvu2u1fVwg==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "@nestjs/cli": {
			"version": "7.2.0",
			"resolved": "https://registry.npmjs.org/@nestjs/cli/-/cli-7.2.0.tgz",
			"integrity": "sha512-LMoi7CzOzteNsho6pz4g4KaqjBfIuugO27Om2hiZBNkxIXL2OiGgDX4DBaouwc+HhGs6v86rDg588p4Y0PJ/PQ==",
			"dev": true,
			"requires": {
			  "@angular-devkit/core": "9.1.7",
			  "@angular-devkit/schematics": "9.1.7",
			  "@angular-devkit/schematics-cli": "0.901.7",
			  "@nestjs/schematics": "^7.0.0",
			  "@types/webpack": "4.41.13",
			  "chalk": "3.0.0",
			  "chokidar": "3.4.0",
			  "cli-table3": "0.5.1",
			  "commander": "4.1.1",
			  "fork-ts-checker-webpack-plugin": "4.1.4",
			  "inquirer": "7.1.0",
			  "node-emoji": "1.10.0",
			  "ora": "4.0.4",
			  "os-name": "3.1.0",
			  "rimraf": "3.0.2",
			  "shelljs": "0.8.4",
			  "tree-kill": "1.2.2",
			  "tsconfig-paths": "3.9.0",
			  "tsconfig-paths-webpack-plugin": "3.2.0",
			  "typescript": "^3.6.4",
			  "webpack": "4.43.0",
			  "webpack-node-externals": "1.7.2"
			},
			"dependencies": {
			  "@angular-devkit/core": {
				"version": "9.1.7",
				"resolved": "https://registry.npmjs.org/@angular-devkit/core/-/core-9.1.7.tgz",
				"integrity": "sha512-guvolu9Cl+qYMTtedLZD9wCqustJjdqzJ2psD2C1Sr1LrX9T0mprmDldR/YnhsitThveJEb6sM/0EvqWxoSvKw==",
				"dev": true,
				"requires": {
				  "ajv": "6.12.0",
				  "fast-json-stable-stringify": "2.1.0",
				  "magic-string": "0.25.7",
				  "rxjs": "6.5.4",
				  "source-map": "0.7.3"
				}
			  },
			  "@angular-devkit/schematics": {
				"version": "9.1.7",
				"resolved": "https://registry.npmjs.org/@angular-devkit/schematics/-/schematics-9.1.7.tgz",
				"integrity": "sha512-oeHPJePBcPp/bd94jHQeFUnft93PGF5iJiKV9szxqS8WWC5OMZ5eK7icRY0PwvLyfenspAZxdZcNaqJqPMul5A==",
				"dev": true,
				"requires": {
				  "@angular-devkit/core": "9.1.7",
				  "ora": "4.0.3",
				  "rxjs": "6.5.4"
				},
				"dependencies": {
				  "ora": {
					"version": "4.0.3",
					"resolved": "https://registry.npmjs.org/ora/-/ora-4.0.3.tgz",
					"integrity": "sha512-fnDebVFyz309A73cqCipVL1fBZewq4vwgSHfxh43vVy31mbyoQ8sCH3Oeaog/owYOs/lLlGVPCISQonTneg6Pg==",
					"dev": true,
					"requires": {
					  "chalk": "^3.0.0",
					  "cli-cursor": "^3.1.0",
					  "cli-spinners": "^2.2.0",
					  "is-interactive": "^1.0.0",
					  "log-symbols": "^3.0.0",
					  "mute-stream": "0.0.8",
					  "strip-ansi": "^6.0.0",
					  "wcwidth": "^1.0.1"
					}
				  }
				}
			  },
			  "ajv": {
				"version": "6.12.0",
				"resolved": "https://registry.npmjs.org/ajv/-/ajv-6.12.0.tgz",
				"integrity": "sha512-D6gFiFA0RRLyUbvijN74DWAjXSFxWKaWP7mldxkVhyhAV3+SWA9HEJPHQ2c9soIeTFJqcSdFDGFgdqs1iUU2Hw==",
				"dev": true,
				"requires": {
				  "fast-deep-equal": "^3.1.1",
				  "fast-json-stable-stringify": "^2.0.0",
				  "json-schema-traverse": "^0.4.1",
				  "uri-js": "^4.2.2"
				}
			  },
			  "ansi-regex": {
				"version": "5.0.0",
				"resolved": "https://registry.npmjs.org/ansi-regex/-/ansi-regex-5.0.0.tgz",
				"integrity": "sha512-bY6fj56OUQ0hU1KjFNDQuJFezqKdrAyFdIevADiqrWHwSlbmBNMHp5ak2f40Pm8JTFyM2mqxkG6ngkHO11f/lg==",
				"dev": true
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "3.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-3.0.0.tgz",
				"integrity": "sha512-4D3B6Wf41KOYRFdszmDqMCGq5VV/uMAB273JILmO+3jAlh8X4qDtdtgCR3fxtbLEMzSx22QdhnDcJvu2u1fVwg==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "fast-deep-equal": {
				"version": "3.1.1",
				"resolved": "https://registry.npmjs.org/fast-deep-equal/-/fast-deep-equal-3.1.1.tgz",
				"integrity": "sha512-8UEa58QDLauDNfpbrX55Q9jrGHThw2ZMdOky5Gl1CDtVeJDPVrG4Jxx1N8jw2gkWaff5UUuX1KJd+9zGe2B+ZA==",
				"dev": true
			  },
			  "fast-json-stable-stringify": {
				"version": "2.1.0",
				"resolved": "https://registry.npmjs.org/fast-json-stable-stringify/-/fast-json-stable-stringify-2.1.0.tgz",
				"integrity": "sha512-lhd/wF+Lk98HZoTCtlVraHtfh5XYijIjalXck7saUtuanSDyLMxnHhSXEDJqHxD7msR8D0uCmqlkwjCV8xvwHw==",
				"dev": true
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "magic-string": {
				"version": "0.25.7",
				"resolved": "https://registry.npmjs.org/magic-string/-/magic-string-0.25.7.tgz",
				"integrity": "sha512-4CrMT5DOHTDk4HYDlzmwu4FVCcIYI8gauveasrdCu2IKIFOJ3f0v/8MDGJCDL9oD2ppz/Av1b0Nj345H9M+XIA==",
				"dev": true,
				"requires": {
				  "sourcemap-codec": "^1.4.4"
				}
			  },
			  "rxjs": {
				"version": "6.5.4",
				"resolved": "https://registry.npmjs.org/rxjs/-/rxjs-6.5.4.tgz",
				"integrity": "sha512-naMQXcgEo3csAEGvw/NydRA0fuS2nDZJiw1YUWFKU7aPPAPGZEsD4Iimit96qwCieH6y614MCLYwdkrWx7z/7Q==",
				"dev": true,
				"requires": {
				  "tslib": "^1.9.0"
				}
			  },
			  "strip-ansi": {
				"version": "6.0.0",
				"resolved": "https://registry.npmjs.org/strip-ansi/-/strip-ansi-6.0.0.tgz",
				"integrity": "sha512-AuvKTrTfQNYNIctbR1K/YGTR1756GycPsg7b9bdV9Duqur4gv6aKqHXah67Z8ImS7WEz5QVcOtlfW2rZEugt6w==",
				"dev": true,
				"requires": {
				  "ansi-regex": "^5.0.0"
				}
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "@nestjs/common": {
			"version": "7.1.1",
			"resolved": "https://registry.npmjs.org/@nestjs/common/-/common-7.1.1.tgz",
			"integrity": "sha512-ZcFbHw2LEkTG9lR84ysNutHnrZPiRYjdjHM7nDCz0zXra98Cu4mKw0wOdoMmtd6WFtpVaSd9JVMXYWvHleFzmw==",
			"requires": {
			  "axios": "0.19.2",
			  "cli-color": "2.0.0",
			  "iterare": "1.2.0",
			  "tslib": "2.0.0",
			  "uuid": "8.0.0"
			},
			"dependencies": {
			  "tslib": {
				"version": "2.0.0",
				"resolved": "https://registry.npmjs.org/tslib/-/tslib-2.0.0.tgz",
				"integrity": "sha512-lTqkx847PI7xEDYJntxZH89L2/aXInsyF2luSafe/+0fHOMjlBNXdH6th7f70qxLDhul7KZK0zC8V5ZIyHl0/g=="
			  },
			  "uuid": {
				"version": "8.0.0",
				"resolved": "https://registry.npmjs.org/uuid/-/uuid-8.0.0.tgz",
				"integrity": "sha512-jOXGuXZAWdsTH7eZLtyXMqUb9EcWMGZNbL9YcGBJl4MH4nrxHmZJhEHvyLFrkxo+28uLb/NYRcStH48fnD0Vzw=="
			  }
			}
		  },
		  "@nestjs/core": {
			"version": "7.1.1",
			"resolved": "https://registry.npmjs.org/@nestjs/core/-/core-7.1.1.tgz",
			"integrity": "sha512-YOMyjWAOtYy1pzz4OsEKSujLKEeY0YML3HPekQPpC2fcgWHBtt0CMJ9RK4385FSltcFW13PgHUzLmcqUgrJTNQ==",
			"requires": {
			  "@nuxtjs/opencollective": "0.2.2",
			  "fast-safe-stringify": "2.0.7",
			  "iterare": "1.2.0",
			  "object-hash": "2.0.3",
			  "path-to-regexp": "3.2.0",
			  "tslib": "2.0.0",
			  "uuid": "8.0.0"
			},
			"dependencies": {
			  "tslib": {
				"version": "2.0.0",
				"resolved": "https://registry.npmjs.org/tslib/-/tslib-2.0.0.tgz",
				"integrity": "sha512-lTqkx847PI7xEDYJntxZH89L2/aXInsyF2luSafe/+0fHOMjlBNXdH6th7f70qxLDhul7KZK0zC8V5ZIyHl0/g=="
			  },
			  "uuid": {
				"version": "8.0.0",
				"resolved": "https://registry.npmjs.org/uuid/-/uuid-8.0.0.tgz",
				"integrity": "sha512-jOXGuXZAWdsTH7eZLtyXMqUb9EcWMGZNbL9YcGBJl4MH4nrxHmZJhEHvyLFrkxo+28uLb/NYRcStH48fnD0Vzw=="
			  }
			}
		  },
		  "@nestjs/platform-express": {
			"version": "7.1.1",
			"resolved": "https://registry.npmjs.org/@nestjs/platform-express/-/platform-express-7.1.1.tgz",
			"integrity": "sha512-g8XNSFTtzsABelQ0uBveCeh6Jf7mUDfGHOXDybaEP//tM+cLU2vbzDYJSHaL68PfGCbhtvzaMiK9ezIWdW9rWg==",
			"requires": {
			  "body-parser": "1.19.0",
			  "cors": "2.8.5",
			  "express": "4.17.1",
			  "multer": "1.4.2",
			  "tslib": "2.0.0"
			},
			"dependencies": {
			  "tslib": {
				"version": "2.0.0",
				"resolved": "https://registry.npmjs.org/tslib/-/tslib-2.0.0.tgz",
				"integrity": "sha512-lTqkx847PI7xEDYJntxZH89L2/aXInsyF2luSafe/+0fHOMjlBNXdH6th7f70qxLDhul7KZK0zC8V5ZIyHl0/g=="
			  }
			}
		  },
		  "@nestjs/schematics": {
			"version": "7.0.0",
			"resolved": "https://registry.npmjs.org/@nestjs/schematics/-/schematics-7.0.0.tgz",
			"integrity": "sha512-R7Yw7wbRk5FNLyeJhvnAys+6ZXymIQWD6vKHOMRbSlR5LV8Q2F0st0EVw/sMmBOtp0EYAwGiXjwNV4ye/znmHA==",
			"dev": true,
			"requires": {
			  "@angular-devkit/core": "9.0.6",
			  "@angular-devkit/schematics": "9.0.6",
			  "fs-extra": "8.1.0"
			}
		  },
		  "@nestjs/testing": {
			"version": "7.1.1",
			"resolved": "https://registry.npmjs.org/@nestjs/testing/-/testing-7.1.1.tgz",
			"integrity": "sha512-J/cdpmUUikYoIwat2LY5eE6ebKDHfcOJXyzDN+7BU8pTINik3Tu6UlguR2C/OnmE6nRCw63yOSuTuFeAWuVa4w==",
			"dev": true,
			"requires": {
			  "optional": "0.1.4",
			  "tslib": "2.0.0"
			},
			"dependencies": {
			  "tslib": {
				"version": "2.0.0",
				"resolved": "https://registry.npmjs.org/tslib/-/tslib-2.0.0.tgz",
				"integrity": "sha512-lTqkx847PI7xEDYJntxZH89L2/aXInsyF2luSafe/+0fHOMjlBNXdH6th7f70qxLDhul7KZK0zC8V5ZIyHl0/g==",
				"dev": true
			  }
			}
		  },
		  "@nuxtjs/opencollective": {
			"version": "0.2.2",
			"resolved": "https://registry.npmjs.org/@nuxtjs/opencollective/-/opencollective-0.2.2.tgz",
			"integrity": "sha512-69gFVDs7mJfNjv9Zs5DFVD+pvBW+k1TaHSOqUWqAyTTfLcKI/EMYQgvEvziRd+zAFtUOoye6MfWh0qvinGISPw==",
			"requires": {
			  "chalk": "^2.4.1",
			  "consola": "^2.3.0",
			  "node-fetch": "^2.3.0"
			}
		  },
		  "@schematics/schematics": {
			"version": "0.901.7",
			"resolved": "https://registry.npmjs.org/@schematics/schematics/-/schematics-0.901.7.tgz",
			"integrity": "sha512-tn45fGbL63EXnKb9Hl85y5Dy6J6cJxVe706aMOZDSzP4z2e4Lem5CCrmA4M9+2RGmyj0UaV/O6/jkmg8W5ZvRQ==",
			"dev": true,
			"requires": {
			  "@angular-devkit/core": "9.1.7",
			  "@angular-devkit/schematics": "9.1.7"
			},
			"dependencies": {
			  "@angular-devkit/core": {
				"version": "9.1.7",
				"resolved": "https://registry.npmjs.org/@angular-devkit/core/-/core-9.1.7.tgz",
				"integrity": "sha512-guvolu9Cl+qYMTtedLZD9wCqustJjdqzJ2psD2C1Sr1LrX9T0mprmDldR/YnhsitThveJEb6sM/0EvqWxoSvKw==",
				"dev": true,
				"requires": {
				  "ajv": "6.12.0",
				  "fast-json-stable-stringify": "2.1.0",
				  "magic-string": "0.25.7",
				  "rxjs": "6.5.4",
				  "source-map": "0.7.3"
				}
			  },
			  "@angular-devkit/schematics": {
				"version": "9.1.7",
				"resolved": "https://registry.npmjs.org/@angular-devkit/schematics/-/schematics-9.1.7.tgz",
				"integrity": "sha512-oeHPJePBcPp/bd94jHQeFUnft93PGF5iJiKV9szxqS8WWC5OMZ5eK7icRY0PwvLyfenspAZxdZcNaqJqPMul5A==",
				"dev": true,
				"requires": {
				  "@angular-devkit/core": "9.1.7",
				  "ora": "4.0.3",
				  "rxjs": "6.5.4"
				}
			  },
			  "ajv": {
				"version": "6.12.0",
				"resolved": "https://registry.npmjs.org/ajv/-/ajv-6.12.0.tgz",
				"integrity": "sha512-D6gFiFA0RRLyUbvijN74DWAjXSFxWKaWP7mldxkVhyhAV3+SWA9HEJPHQ2c9soIeTFJqcSdFDGFgdqs1iUU2Hw==",
				"dev": true,
				"requires": {
				  "fast-deep-equal": "^3.1.1",
				  "fast-json-stable-stringify": "^2.0.0",
				  "json-schema-traverse": "^0.4.1",
				  "uri-js": "^4.2.2"
				}
			  },
			  "ansi-regex": {
				"version": "5.0.0",
				"resolved": "https://registry.npmjs.org/ansi-regex/-/ansi-regex-5.0.0.tgz",
				"integrity": "sha512-bY6fj56OUQ0hU1KjFNDQuJFezqKdrAyFdIevADiqrWHwSlbmBNMHp5ak2f40Pm8JTFyM2mqxkG6ngkHO11f/lg==",
				"dev": true
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "3.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-3.0.0.tgz",
				"integrity": "sha512-4D3B6Wf41KOYRFdszmDqMCGq5VV/uMAB273JILmO+3jAlh8X4qDtdtgCR3fxtbLEMzSx22QdhnDcJvu2u1fVwg==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "fast-deep-equal": {
				"version": "3.1.1",
				"resolved": "https://registry.npmjs.org/fast-deep-equal/-/fast-deep-equal-3.1.1.tgz",
				"integrity": "sha512-8UEa58QDLauDNfpbrX55Q9jrGHThw2ZMdOky5Gl1CDtVeJDPVrG4Jxx1N8jw2gkWaff5UUuX1KJd+9zGe2B+ZA==",
				"dev": true
			  },
			  "fast-json-stable-stringify": {
				"version": "2.1.0",
				"resolved": "https://registry.npmjs.org/fast-json-stable-stringify/-/fast-json-stable-stringify-2.1.0.tgz",
				"integrity": "sha512-lhd/wF+Lk98HZoTCtlVraHtfh5XYijIjalXck7saUtuanSDyLMxnHhSXEDJqHxD7msR8D0uCmqlkwjCV8xvwHw==",
				"dev": true
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "magic-string": {
				"version": "0.25.7",
				"resolved": "https://registry.npmjs.org/magic-string/-/magic-string-0.25.7.tgz",
				"integrity": "sha512-4CrMT5DOHTDk4HYDlzmwu4FVCcIYI8gauveasrdCu2IKIFOJ3f0v/8MDGJCDL9oD2ppz/Av1b0Nj345H9M+XIA==",
				"dev": true,
				"requires": {
				  "sourcemap-codec": "^1.4.4"
				}
			  },
			  "ora": {
				"version": "4.0.3",
				"resolved": "https://registry.npmjs.org/ora/-/ora-4.0.3.tgz",
				"integrity": "sha512-fnDebVFyz309A73cqCipVL1fBZewq4vwgSHfxh43vVy31mbyoQ8sCH3Oeaog/owYOs/lLlGVPCISQonTneg6Pg==",
				"dev": true,
				"requires": {
				  "chalk": "^3.0.0",
				  "cli-cursor": "^3.1.0",
				  "cli-spinners": "^2.2.0",
				  "is-interactive": "^1.0.0",
				  "log-symbols": "^3.0.0",
				  "mute-stream": "0.0.8",
				  "strip-ansi": "^6.0.0",
				  "wcwidth": "^1.0.1"
				}
			  },
			  "rxjs": {
				"version": "6.5.4",
				"resolved": "https://registry.npmjs.org/rxjs/-/rxjs-6.5.4.tgz",
				"integrity": "sha512-naMQXcgEo3csAEGvw/NydRA0fuS2nDZJiw1YUWFKU7aPPAPGZEsD4Iimit96qwCieH6y614MCLYwdkrWx7z/7Q==",
				"dev": true,
				"requires": {
				  "tslib": "^1.9.0"
				}
			  },
			  "strip-ansi": {
				"version": "6.0.0",
				"resolved": "https://registry.npmjs.org/strip-ansi/-/strip-ansi-6.0.0.tgz",
				"integrity": "sha512-AuvKTrTfQNYNIctbR1K/YGTR1756GycPsg7b9bdV9Duqur4gv6aKqHXah67Z8ImS7WEz5QVcOtlfW2rZEugt6w==",
				"dev": true,
				"requires": {
				  "ansi-regex": "^5.0.0"
				}
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "@sinonjs/commons": {
			"version": "1.7.2",
			"resolved": "https://registry.npmjs.org/@sinonjs/commons/-/commons-1.7.2.tgz",
			"integrity": "sha512-+DUO6pnp3udV/v2VfUWgaY5BIE1IfT7lLfeDzPVeMT1XKkaAp9LgSI9x5RtrFQoZ9Oi0PgXQQHPaoKu7dCjVxw==",
			"dev": true,
			"requires": {
			  "type-detect": "4.0.8"
			}
		  },
		  "@sinonjs/fake-timers": {
			"version": "6.0.1",
			"resolved": "https://registry.npmjs.org/@sinonjs/fake-timers/-/fake-timers-6.0.1.tgz",
			"integrity": "sha512-MZPUxrmFubI36XS1DI3qmI0YdN1gks62JtFZvxR67ljjSNCeK6U08Zx4msEWOXuofgqUt6zPHSi1H9fbjR/NRA==",
			"dev": true,
			"requires": {
			  "@sinonjs/commons": "^1.7.0"
			}
		  },
		  "@types/anymatch": {
			"version": "1.3.1",
			"resolved": "https://registry.npmjs.org/@types/anymatch/-/anymatch-1.3.1.tgz",
			"integrity": "sha512-/+CRPXpBDpo2RK9C68N3b2cOvO0Cf5B9aPijHsoDQTHivnGSObdOF2BRQOYjojWTDy6nQvMjmqRXIxH55VjxxA==",
			"dev": true
		  },
		  "@types/babel__core": {
			"version": "7.1.7",
			"resolved": "https://registry.npmjs.org/@types/babel__core/-/babel__core-7.1.7.tgz",
			"integrity": "sha512-RL62NqSFPCDK2FM1pSDH0scHpJvsXtZNiYlMB73DgPBaG1E38ZYVL+ei5EkWRbr+KC4YNiAUNBnRj+bgwpgjMw==",
			"dev": true,
			"requires": {
			  "@babel/parser": "^7.1.0",
			  "@babel/types": "^7.0.0",
			  "@types/babel__generator": "*",
			  "@types/babel__template": "*",
			  "@types/babel__traverse": "*"
			}
		  },
		  "@types/babel__generator": {
			"version": "7.6.1",
			"resolved": "https://registry.npmjs.org/@types/babel__generator/-/babel__generator-7.6.1.tgz",
			"integrity": "sha512-bBKm+2VPJcMRVwNhxKu8W+5/zT7pwNEqeokFOmbvVSqGzFneNxYcEBro9Ac7/N9tlsaPYnZLK8J1LWKkMsLAew==",
			"dev": true,
			"requires": {
			  "@babel/types": "^7.0.0"
			}
		  },
		  "@types/babel__template": {
			"version": "7.0.2",
			"resolved": "https://registry.npmjs.org/@types/babel__template/-/babel__template-7.0.2.tgz",
			"integrity": "sha512-/K6zCpeW7Imzgab2bLkLEbz0+1JlFSrUMdw7KoIIu+IUdu51GWaBZpd3y1VXGVXzynvGa4DaIaxNZHiON3GXUg==",
			"dev": true,
			"requires": {
			  "@babel/parser": "^7.1.0",
			  "@babel/types": "^7.0.0"
			}
		  },
		  "@types/babel__traverse": {
			"version": "7.0.11",
			"resolved": "https://registry.npmjs.org/@types/babel__traverse/-/babel__traverse-7.0.11.tgz",
			"integrity": "sha512-ddHK5icION5U6q11+tV2f9Mo6CZVuT8GJKld2q9LqHSZbvLbH34Kcu2yFGckZut453+eQU6btIA3RihmnRgI+Q==",
			"dev": true,
			"requires": {
			  "@babel/types": "^7.3.0"
			}
		  },
		  "@types/body-parser": {
			"version": "1.19.0",
			"resolved": "https://registry.npmjs.org/@types/body-parser/-/body-parser-1.19.0.tgz",
			"integrity": "sha512-W98JrE0j2K78swW4ukqMleo8R7h/pFETjM2DQ90MF6XK2i4LO4W3gQ71Lt4w3bfm2EvVSyWHplECvB5sK22yFQ==",
			"dev": true,
			"requires": {
			  "@types/connect": "*",
			  "@types/node": "*"
			}
		  },
		  "@types/color-name": {
			"version": "1.1.1",
			"resolved": "https://registry.npmjs.org/@types/color-name/-/color-name-1.1.1.tgz",
			"integrity": "sha512-rr+OQyAjxze7GgWrSaJwydHStIhHq2lvY3BOC2Mj7KnzI7XK0Uw1TOOdI9lDoajEbSWLiYgoo4f1R51erQfhPQ==",
			"dev": true
		  },
		  "@types/connect": {
			"version": "3.4.33",
			"resolved": "https://registry.npmjs.org/@types/connect/-/connect-3.4.33.tgz",
			"integrity": "sha512-2+FrkXY4zllzTNfJth7jOqEHC+enpLeGslEhpnTAkg21GkRrWV4SsAtqchtT4YS9/nODBU2/ZfsBY2X4J/dX7A==",
			"dev": true,
			"requires": {
			  "@types/node": "*"
			}
		  },
		  "@types/cookiejar": {
			"version": "2.1.1",
			"resolved": "https://registry.npmjs.org/@types/cookiejar/-/cookiejar-2.1.1.tgz",
			"integrity": "sha512-aRnpPa7ysx3aNW60hTiCtLHlQaIFsXFCgQlpakNgDNVFzbtusSY8PwjAQgRWfSk0ekNoBjO51eQRB6upA9uuyw==",
			"dev": true
		  },
		  "@types/eslint-visitor-keys": {
			"version": "1.0.0",
			"resolved": "https://registry.npmjs.org/@types/eslint-visitor-keys/-/eslint-visitor-keys-1.0.0.tgz",
			"integrity": "sha512-OCutwjDZ4aFS6PB1UZ988C4YgwlBHJd6wCeQqaLdmadZ/7e+w79+hbMUFC1QXDNCmdyoRfAFdm0RypzwR+Qpag==",
			"dev": true
		  },
		  "@types/express": {
			"version": "4.17.6",
			"resolved": "https://registry.npmjs.org/@types/express/-/express-4.17.6.tgz",
			"integrity": "sha512-n/mr9tZI83kd4azlPG5y997C/M4DNABK9yErhFM6hKdym4kkmd9j0vtsJyjFIwfRBxtrxZtAfGZCNRIBMFLK5w==",
			"dev": true,
			"requires": {
			  "@types/body-parser": "*",
			  "@types/express-serve-static-core": "*",
			  "@types/qs": "*",
			  "@types/serve-static": "*"
			}
		  },
		  "@types/express-serve-static-core": {
			"version": "4.17.7",
			"resolved": "https://registry.npmjs.org/@types/express-serve-static-core/-/express-serve-static-core-4.17.7.tgz",
			"integrity": "sha512-EMgTj/DF9qpgLXyc+Btimg+XoH7A2liE8uKul8qSmMTHCeNYzydDKFdsJskDvw42UsesCnhO63dO0Grbj8J4Dw==",
			"dev": true,
			"requires": {
			  "@types/node": "*",
			  "@types/qs": "*",
			  "@types/range-parser": "*"
			}
		  },
		  "@types/graceful-fs": {
			"version": "4.1.3",
			"resolved": "https://registry.npmjs.org/@types/graceful-fs/-/graceful-fs-4.1.3.tgz",
			"integrity": "sha512-AiHRaEB50LQg0pZmm659vNBb9f4SJ0qrAnteuzhSeAUcJKxoYgEnprg/83kppCnc2zvtCKbdZry1a5pVY3lOTQ==",
			"dev": true,
			"requires": {
			  "@types/node": "*"
			}
		  },
		  "@types/istanbul-lib-coverage": {
			"version": "2.0.2",
			"resolved": "https://registry.npmjs.org/@types/istanbul-lib-coverage/-/istanbul-lib-coverage-2.0.2.tgz",
			"integrity": "sha512-rsZg7eL+Xcxsxk2XlBt9KcG8nOp9iYdKCOikY9x2RFJCyOdNj4MKPQty0e8oZr29vVAzKXr1BmR+kZauti3o1w==",
			"dev": true
		  },
		  "@types/istanbul-lib-report": {
			"version": "3.0.0",
			"resolved": "https://registry.npmjs.org/@types/istanbul-lib-report/-/istanbul-lib-report-3.0.0.tgz",
			"integrity": "sha512-plGgXAPfVKFoYfa9NpYDAkseG+g6Jr294RqeqcqDixSbU34MZVJRi/P+7Y8GDpzkEwLaGZZOpKIEmeVZNtKsrg==",
			"dev": true,
			"requires": {
			  "@types/istanbul-lib-coverage": "*"
			}
		  },
		  "@types/istanbul-reports": {
			"version": "1.1.2",
			"resolved": "https://registry.npmjs.org/@types/istanbul-reports/-/istanbul-reports-1.1.2.tgz",
			"integrity": "sha512-P/W9yOX/3oPZSpaYOCQzGqgCQRXn0FFO/V8bWrCQs+wLmvVVxk6CRBXALEvNs9OHIatlnlFokfhuDo2ug01ciw==",
			"dev": true,
			"requires": {
			  "@types/istanbul-lib-coverage": "*",
			  "@types/istanbul-lib-report": "*"
			}
		  },
		  "@types/jest": {
			"version": "25.2.3",
			"resolved": "https://registry.npmjs.org/@types/jest/-/jest-25.2.3.tgz",
			"integrity": "sha512-JXc1nK/tXHiDhV55dvfzqtmP4S3sy3T3ouV2tkViZgxY/zeUkcpQcQPGRlgF4KmWzWW5oiWYSZwtCB+2RsE4Fw==",
			"dev": true,
			"requires": {
			  "jest-diff": "^25.2.1",
			  "pretty-format": "^25.2.1"
			}
		  },
		  "@types/json-schema": {
			"version": "7.0.4",
			"resolved": "https://registry.npmjs.org/@types/json-schema/-/json-schema-7.0.4.tgz",
			"integrity": "sha512-8+KAKzEvSUdeo+kmqnKrqgeE+LcA0tjYWFY7RPProVYwnqDjukzO+3b6dLD56rYX5TdWejnEOLJYOIeh4CXKuA==",
			"dev": true
		  },
		  "@types/json5": {
			"version": "0.0.29",
			"resolved": "https://registry.npmjs.org/@types/json5/-/json5-0.0.29.tgz",
			"integrity": "sha1-7ihweulOEdK4J7y+UnC86n8+ce4=",
			"dev": true
		  },
		  "@types/mime": {
			"version": "2.0.1",
			"resolved": "https://registry.npmjs.org/@types/mime/-/mime-2.0.1.tgz",
			"integrity": "sha512-FwI9gX75FgVBJ7ywgnq/P7tw+/o1GUbtP0KzbtusLigAOgIgNISRK0ZPl4qertvXSIE8YbsVJueQ90cDt9YYyw==",
			"dev": true
		  },
		  "@types/node": {
			"version": "14.0.6",
			"resolved": "https://registry.npmjs.org/@types/node/-/node-14.0.6.tgz",
			"integrity": "sha512-FbNmu4F67d3oZMWBV6Y4MaPER+0EpE9eIYf2yaHhCWovc1dlXCZkqGX4NLHfVVr6umt20TNBdRzrNJIzIKfdbw==",
			"dev": true
		  },
		  "@types/normalize-package-data": {
			"version": "2.4.0",
			"resolved": "https://registry.npmjs.org/@types/normalize-package-data/-/normalize-package-data-2.4.0.tgz",
			"integrity": "sha512-f5j5b/Gf71L+dbqxIpQ4Z2WlmI/mPJ0fOkGGmFgtb6sAu97EPczzbS3/tJKxmcYDj55OX6ssqwDAWOHIYDRDGA==",
			"dev": true
		  },
		  "@types/prettier": {
			"version": "2.0.0",
			"resolved": "https://registry.npmjs.org/@types/prettier/-/prettier-2.0.0.tgz",
			"integrity": "sha512-/rM+sWiuOZ5dvuVzV37sUuklsbg+JPOP8d+nNFlo2ZtfpzPiPvh1/gc8liWOLBqe+sR+ZM7guPaIcTt6UZTo7Q==",
			"dev": true
		  },
		  "@types/qs": {
			"version": "6.9.2",
			"resolved": "https://registry.npmjs.org/@types/qs/-/qs-6.9.2.tgz",
			"integrity": "sha512-a9bDi4Z3zCZf4Lv1X/vwnvbbDYSNz59h3i3KdyuYYN+YrLjSeJD0dnphdULDfySvUv6Exy/O0K6wX/kQpnPQ+A==",
			"dev": true
		  },
		  "@types/range-parser": {
			"version": "1.2.3",
			"resolved": "https://registry.npmjs.org/@types/range-parser/-/range-parser-1.2.3.tgz",
			"integrity": "sha512-ewFXqrQHlFsgc09MK5jP5iR7vumV/BYayNC6PgJO2LPe8vrnNFyjQjSppfEngITi0qvfKtzFvgKymGheFM9UOA==",
			"dev": true
		  },
		  "@types/serve-static": {
			"version": "1.13.3",
			"resolved": "https://registry.npmjs.org/@types/serve-static/-/serve-static-1.13.3.tgz",
			"integrity": "sha512-oprSwp094zOglVrXdlo/4bAHtKTAxX6VT8FOZlBKrmyLbNvE1zxZyJ6yikMVtHIvwP45+ZQGJn+FdXGKTozq0g==",
			"dev": true,
			"requires": {
			  "@types/express-serve-static-core": "*",
			  "@types/mime": "*"
			}
		  },
		  "@types/source-list-map": {
			"version": "0.1.2",
			"resolved": "https://registry.npmjs.org/@types/source-list-map/-/source-list-map-0.1.2.tgz",
			"integrity": "sha512-K5K+yml8LTo9bWJI/rECfIPrGgxdpeNbj+d53lwN4QjW1MCwlkhUms+gtdzigTeUyBr09+u8BwOIY3MXvHdcsA==",
			"dev": true
		  },
		  "@types/stack-utils": {
			"version": "1.0.1",
			"resolved": "https://registry.npmjs.org/@types/stack-utils/-/stack-utils-1.0.1.tgz",
			"integrity": "sha512-l42BggppR6zLmpfU6fq9HEa2oGPEI8yrSPL3GITjfRInppYFahObbIQOQK3UGxEnyQpltZLaPe75046NOZQikw==",
			"dev": true
		  },
		  "@types/superagent": {
			"version": "4.1.7",
			"resolved": "https://registry.npmjs.org/@types/superagent/-/superagent-4.1.7.tgz",
			"integrity": "sha512-JSwNPgRYjIC4pIeOqLwWwfGj6iP1n5NE6kNBEbGx2V8H78xCPwx7QpNp9plaI30+W3cFEzJO7BIIsXE+dbtaGg==",
			"dev": true,
			"requires": {
			  "@types/cookiejar": "*",
			  "@types/node": "*"
			}
		  },
		  "@types/supertest": {
			"version": "2.0.9",
			"resolved": "https://registry.npmjs.org/@types/supertest/-/supertest-2.0.9.tgz",
			"integrity": "sha512-0BTpWWWAO1+uXaP/oA0KW1eOZv4hc0knhrWowV06Gwwz3kqQxNO98fUFM2e15T+PdPRmOouNFrYvaBgdojPJ3g==",
			"dev": true,
			"requires": {
			  "@types/superagent": "*"
			}
		  },
		  "@types/tapable": {
			"version": "1.0.5",
			"resolved": "https://registry.npmjs.org/@types/tapable/-/tapable-1.0.5.tgz",
			"integrity": "sha512-/gG2M/Imw7cQFp8PGvz/SwocNrmKFjFsm5Pb8HdbHkZ1K8pmuPzOX4VeVoiEecFCVf4CsN1r3/BRvx+6sNqwtQ==",
			"dev": true
		  },
		  "@types/uglify-js": {
			"version": "3.9.2",
			"resolved": "https://registry.npmjs.org/@types/uglify-js/-/uglify-js-3.9.2.tgz",
			"integrity": "sha512-d6dIfpPbF+8B7WiCi2ELY7m0w1joD8cRW4ms88Emdb2w062NeEpbNCeWwVCgzLRpVG+5e74VFSg4rgJ2xXjEiQ==",
			"dev": true,
			"requires": {
			  "source-map": "^0.6.1"
			},
			"dependencies": {
			  "source-map": {
				"version": "0.6.1",
				"resolved": "https://registry.npmjs.org/source-map/-/source-map-0.6.1.tgz",
				"integrity": "sha512-UjgapumWlbMhkBgzT7Ykc5YXUT46F0iKu8SGXq0bcwP5dz/h0Plj6enJqjz1Zbq2l5WaqYnrVbwWOWMyF3F47g==",
				"dev": true
			  }
			}
		  },
		  "@types/webpack": {
			"version": "4.41.13",
			"resolved": "https://registry.npmjs.org/@types/webpack/-/webpack-4.41.13.tgz",
			"integrity": "sha512-RYmIHOWSxnTTa765N6jJBVE45pd2SYNblEYshVDduLw6RhocazNmRzE5/ytvBD8IkDMH6DI+bcrqxh8NILimBA==",
			"dev": true,
			"requires": {
			  "@types/anymatch": "*",
			  "@types/node": "*",
			  "@types/tapable": "*",
			  "@types/uglify-js": "*",
			  "@types/webpack-sources": "*",
			  "source-map": "^0.6.0"
			},
			"dependencies": {
			  "source-map": {
				"version": "0.6.1",
				"resolved": "https://registry.npmjs.org/source-map/-/source-map-0.6.1.tgz",
				"integrity": "sha512-UjgapumWlbMhkBgzT7Ykc5YXUT46F0iKu8SGXq0bcwP5dz/h0Plj6enJqjz1Zbq2l5WaqYnrVbwWOWMyF3F47g==",
				"dev": true
			  }
			}
		  },
		  "@types/webpack-sources": {
			"version": "0.1.7",
			"resolved": "https://registry.npmjs.org/@types/webpack-sources/-/webpack-sources-0.1.7.tgz",
			"integrity": "sha512-XyaHrJILjK1VHVC4aVlKsdNN5KBTwufMb43cQs+flGxtPAf/1Qwl8+Q0tp5BwEGaI8D6XT1L+9bSWXckgkjTLw==",
			"dev": true,
			"requires": {
			  "@types/node": "*",
			  "@types/source-list-map": "*",
			  "source-map": "^0.6.1"
			},
			"dependencies": {
			  "source-map": {
				"version": "0.6.1",
				"resolved": "https://registry.npmjs.org/source-map/-/source-map-0.6.1.tgz",
				"integrity": "sha512-UjgapumWlbMhkBgzT7Ykc5YXUT46F0iKu8SGXq0bcwP5dz/h0Plj6enJqjz1Zbq2l5WaqYnrVbwWOWMyF3F47g==",
				"dev": true
			  }
			}
		  },
		  "@types/yargs": {
			"version": "15.0.5",
			"resolved": "https://registry.npmjs.org/@types/yargs/-/yargs-15.0.5.tgz",
			"integrity": "sha512-Dk/IDOPtOgubt/IaevIUbTgV7doaKkoorvOyYM2CMwuDyP89bekI7H4xLIwunNYiK9jhCkmc6pUrJk3cj2AB9w==",
			"dev": true,
			"requires": {
			  "@types/yargs-parser": "*"
			}
		  },
		  "@types/yargs-parser": {
			"version": "15.0.0",
			"resolved": "https://registry.npmjs.org/@types/yargs-parser/-/yargs-parser-15.0.0.tgz",
			"integrity": "sha512-FA/BWv8t8ZWJ+gEOnLLd8ygxH/2UFbAvgEonyfN6yWGLKc7zVjbpl2Y4CTjid9h2RfgPP6SEt6uHwEOply00yw==",
			"dev": true
		  },
		  "@typescript-eslint/eslint-plugin": {
			"version": "3.0.2",
			"resolved": "https://registry.npmjs.org/@typescript-eslint/eslint-plugin/-/eslint-plugin-3.0.2.tgz",
			"integrity": "sha512-ER3bSS/A/pKQT/hjMGCK8UQzlL0yLjuCZ/G8CDFJFVTfl3X65fvq2lNYqOG8JPTfrPa2RULCdwfOyFjZEMNExQ==",
			"dev": true,
			"requires": {
			  "@typescript-eslint/experimental-utils": "3.0.2",
			  "functional-red-black-tree": "^1.0.1",
			  "regexpp": "^3.0.0",
			  "semver": "^7.3.2",
			  "tsutils": "^3.17.1"
			},
			"dependencies": {
			  "semver": {
				"version": "7.3.2",
				"resolved": "https://registry.npmjs.org/semver/-/semver-7.3.2.tgz",
				"integrity": "sha512-OrOb32TeeambH6UrhtShmF7CRDqhL6/5XpPNp2DuRH6+9QLw/orhp72j87v8Qa1ScDkvrrBNpZcDejAirJmfXQ==",
				"dev": true
			  }
			}
		  },
		  "@typescript-eslint/experimental-utils": {
			"version": "3.0.2",
			"resolved": "https://registry.npmjs.org/@typescript-eslint/experimental-utils/-/experimental-utils-3.0.2.tgz",
			"integrity": "sha512-4Wc4EczvoY183SSEnKgqAfkj1eLtRgBQ04AAeG+m4RhTVyaazxc1uI8IHf0qLmu7xXe9j1nn+UoDJjbmGmuqXQ==",
			"dev": true,
			"requires": {
			  "@types/json-schema": "^7.0.3",
			  "@typescript-eslint/typescript-estree": "3.0.2",
			  "eslint-scope": "^5.0.0",
			  "eslint-utils": "^2.0.0"
			},
			"dependencies": {
			  "eslint-scope": {
				"version": "5.0.0",
				"resolved": "https://registry.npmjs.org/eslint-scope/-/eslint-scope-5.0.0.tgz",
				"integrity": "sha512-oYrhJW7S0bxAFDvWqzvMPRm6pcgcnWc4QnofCAqRTRfQC0JcwenzGglTtsLyIuuWFfkqDG9vz67cnttSd53djw==",
				"dev": true,
				"requires": {
				  "esrecurse": "^4.1.0",
				  "estraverse": "^4.1.1"
				}
			  }
			}
		  },
		  "@typescript-eslint/parser": {
			"version": "3.0.2",
			"resolved": "https://registry.npmjs.org/@typescript-eslint/parser/-/parser-3.0.2.tgz",
			"integrity": "sha512-80Z7s83e8QXHNUspqVlWwb4t5gdz/1bBBmafElbK1wwAwiD/yvJsFyHRxlEpNrt4rdK6eB3p+2WEFkEDHAKk9w==",
			"dev": true,
			"requires": {
			  "@types/eslint-visitor-keys": "^1.0.0",
			  "@typescript-eslint/experimental-utils": "3.0.2",
			  "@typescript-eslint/typescript-estree": "3.0.2",
			  "eslint-visitor-keys": "^1.1.0"
			}
		  },
		  "@typescript-eslint/typescript-estree": {
			"version": "3.0.2",
			"resolved": "https://registry.npmjs.org/@typescript-eslint/typescript-estree/-/typescript-estree-3.0.2.tgz",
			"integrity": "sha512-cs84mxgC9zQ6viV8MEcigfIKQmKtBkZNDYf8Gru2M+MhnA6z9q0NFMZm2IEzKqAwN8lY5mFVd1Z8DiHj6zQ3Tw==",
			"dev": true,
			"requires": {
			  "debug": "^4.1.1",
			  "eslint-visitor-keys": "^1.1.0",
			  "glob": "^7.1.6",
			  "is-glob": "^4.0.1",
			  "lodash": "^4.17.15",
			  "semver": "^7.3.2",
			  "tsutils": "^3.17.1"
			},
			"dependencies": {
			  "debug": {
				"version": "4.1.1",
				"resolved": "https://registry.npmjs.org/debug/-/debug-4.1.1.tgz",
				"integrity": "sha512-pYAIzeRo8J6KPEaJ0VWOh5Pzkbw/RetuzehGM7QRRX5he4fPHx2rdKMB256ehJCkX+XRQm16eZLqLNS8RSZXZw==",
				"dev": true,
				"requires": {
				  "ms": "^2.1.1"
				}
			  },
			  "glob": {
				"version": "7.1.6",
				"resolved": "https://registry.npmjs.org/glob/-/glob-7.1.6.tgz",
				"integrity": "sha512-LwaxwyZ72Lk7vZINtNNrywX0ZuLyStrdDtabefZKAY5ZGJhVtgdznluResxNmPitE0SAO+O26sWTHeKSI2wMBA==",
				"dev": true,
				"requires": {
				  "fs.realpath": "^1.0.0",
				  "inflight": "^1.0.4",
				  "inherits": "2",
				  "minimatch": "^3.0.4",
				  "once": "^1.3.0",
				  "path-is-absolute": "^1.0.0"
				}
			  },
			  "ms": {
				"version": "2.1.2",
				"resolved": "https://registry.npmjs.org/ms/-/ms-2.1.2.tgz",
				"integrity": "sha512-sGkPx+VjMtmA6MX27oA4FBFELFCZZ4S4XqeGOXCv68tT+jb3vk/RyaKWP0PTKyWtmLSM0b+adUTEvbs1PEaH2w==",
				"dev": true
			  },
			  "semver": {
				"version": "7.3.2",
				"resolved": "https://registry.npmjs.org/semver/-/semver-7.3.2.tgz",
				"integrity": "sha512-OrOb32TeeambH6UrhtShmF7CRDqhL6/5XpPNp2DuRH6+9QLw/orhp72j87v8Qa1ScDkvrrBNpZcDejAirJmfXQ==",
				"dev": true
			  }
			}
		  },
		  "@webassemblyjs/ast": {
			"version": "1.9.0",
			"resolved": "https://registry.npmjs.org/@webassemblyjs/ast/-/ast-1.9.0.tgz",
			"integrity": "sha512-C6wW5L+b7ogSDVqymbkkvuW9kruN//YisMED04xzeBBqjHa2FYnmvOlS6Xj68xWQRgWvI9cIglsjFowH/RJyEA==",
			"dev": true,
			"requires": {
			  "@webassemblyjs/helper-module-context": "1.9.0",
			  "@webassemblyjs/helper-wasm-bytecode": "1.9.0",
			  "@webassemblyjs/wast-parser": "1.9.0"
			}
		  },
		  "@webassemblyjs/floating-point-hex-parser": {
			"version": "1.9.0",
			"resolved": "https://registry.npmjs.org/@webassemblyjs/floating-point-hex-parser/-/floating-point-hex-parser-1.9.0.tgz",
			"integrity": "sha512-TG5qcFsS8QB4g4MhrxK5TqfdNe7Ey/7YL/xN+36rRjl/BlGE/NcBvJcqsRgCP6Z92mRE+7N50pRIi8SmKUbcQA==",
			"dev": true
		  },
		  "@webassemblyjs/helper-api-error": {
			"version": "1.9.0",
			"resolved": "https://registry.npmjs.org/@webassemblyjs/helper-api-error/-/helper-api-error-1.9.0.tgz",
			"integrity": "sha512-NcMLjoFMXpsASZFxJ5h2HZRcEhDkvnNFOAKneP5RbKRzaWJN36NC4jqQHKwStIhGXu5mUWlUUk7ygdtrO8lbmw==",
			"dev": true
		  },
		  "@webassemblyjs/helper-buffer": {
			"version": "1.9.0",
			"resolved": "https://registry.npmjs.org/@webassemblyjs/helper-buffer/-/helper-buffer-1.9.0.tgz",
			"integrity": "sha512-qZol43oqhq6yBPx7YM3m9Bv7WMV9Eevj6kMi6InKOuZxhw+q9hOkvq5e/PpKSiLfyetpaBnogSbNCfBwyB00CA==",
			"dev": true
		  },
		  "@webassemblyjs/helper-code-frame": {
			"version": "1.9.0",
			"resolved": "https://registry.npmjs.org/@webassemblyjs/helper-code-frame/-/helper-code-frame-1.9.0.tgz",
			"integrity": "sha512-ERCYdJBkD9Vu4vtjUYe8LZruWuNIToYq/ME22igL+2vj2dQ2OOujIZr3MEFvfEaqKoVqpsFKAGsRdBSBjrIvZA==",
			"dev": true,
			"requires": {
			  "@webassemblyjs/wast-printer": "1.9.0"
			}
		  },
		  "@webassemblyjs/helper-fsm": {
			"version": "1.9.0",
			"resolved": "https://registry.npmjs.org/@webassemblyjs/helper-fsm/-/helper-fsm-1.9.0.tgz",
			"integrity": "sha512-OPRowhGbshCb5PxJ8LocpdX9Kl0uB4XsAjl6jH/dWKlk/mzsANvhwbiULsaiqT5GZGT9qinTICdj6PLuM5gslw==",
			"dev": true
		  },
		  "@webassemblyjs/helper-module-context": {
			"version": "1.9.0",
			"resolved": "https://registry.npmjs.org/@webassemblyjs/helper-module-context/-/helper-module-context-1.9.0.tgz",
			"integrity": "sha512-MJCW8iGC08tMk2enck1aPW+BE5Cw8/7ph/VGZxwyvGbJwjktKkDK7vy7gAmMDx88D7mhDTCNKAW5tED+gZ0W8g==",
			"dev": true,
			"requires": {
			  "@webassemblyjs/ast": "1.9.0"
			}
		  },
		  "@webassemblyjs/helper-wasm-bytecode": {
			"version": "1.9.0",
			"resolved": "https://registry.npmjs.org/@webassemblyjs/helper-wasm-bytecode/-/helper-wasm-bytecode-1.9.0.tgz",
			"integrity": "sha512-R7FStIzyNcd7xKxCZH5lE0Bqy+hGTwS3LJjuv1ZVxd9O7eHCedSdrId/hMOd20I+v8wDXEn+bjfKDLzTepoaUw==",
			"dev": true
		  },
		  "@webassemblyjs/helper-wasm-section": {
			"version": "1.9.0",
			"resolved": "https://registry.npmjs.org/@webassemblyjs/helper-wasm-section/-/helper-wasm-section-1.9.0.tgz",
			"integrity": "sha512-XnMB8l3ek4tvrKUUku+IVaXNHz2YsJyOOmz+MMkZvh8h1uSJpSen6vYnw3IoQ7WwEuAhL8Efjms1ZWjqh2agvw==",
			"dev": true,
			"requires": {
			  "@webassemblyjs/ast": "1.9.0",
			  "@webassemblyjs/helper-buffer": "1.9.0",
			  "@webassemblyjs/helper-wasm-bytecode": "1.9.0",
			  "@webassemblyjs/wasm-gen": "1.9.0"
			}
		  },
		  "@webassemblyjs/ieee754": {
			"version": "1.9.0",
			"resolved": "https://registry.npmjs.org/@webassemblyjs/ieee754/-/ieee754-1.9.0.tgz",
			"integrity": "sha512-dcX8JuYU/gvymzIHc9DgxTzUUTLexWwt8uCTWP3otys596io0L5aW02Gb1RjYpx2+0Jus1h4ZFqjla7umFniTg==",
			"dev": true,
			"requires": {
			  "@xtuc/ieee754": "^1.2.0"
			}
		  },
		  "@webassemblyjs/leb128": {
			"version": "1.9.0",
			"resolved": "https://registry.npmjs.org/@webassemblyjs/leb128/-/leb128-1.9.0.tgz",
			"integrity": "sha512-ENVzM5VwV1ojs9jam6vPys97B/S65YQtv/aanqnU7D8aSoHFX8GyhGg0CMfyKNIHBuAVjy3tlzd5QMMINa7wpw==",
			"dev": true,
			"requires": {
			  "@xtuc/long": "4.2.2"
			}
		  },
		  "@webassemblyjs/utf8": {
			"version": "1.9.0",
			"resolved": "https://registry.npmjs.org/@webassemblyjs/utf8/-/utf8-1.9.0.tgz",
			"integrity": "sha512-GZbQlWtopBTP0u7cHrEx+73yZKrQoBMpwkGEIqlacljhXCkVM1kMQge/Mf+csMJAjEdSwhOyLAS0AoR3AG5P8w==",
			"dev": true
		  },
		  "@webassemblyjs/wasm-edit": {
			"version": "1.9.0",
			"resolved": "https://registry.npmjs.org/@webassemblyjs/wasm-edit/-/wasm-edit-1.9.0.tgz",
			"integrity": "sha512-FgHzBm80uwz5M8WKnMTn6j/sVbqilPdQXTWraSjBwFXSYGirpkSWE2R9Qvz9tNiTKQvoKILpCuTjBKzOIm0nxw==",
			"dev": true,
			"requires": {
			  "@webassemblyjs/ast": "1.9.0",
			  "@webassemblyjs/helper-buffer": "1.9.0",
			  "@webassemblyjs/helper-wasm-bytecode": "1.9.0",
			  "@webassemblyjs/helper-wasm-section": "1.9.0",
			  "@webassemblyjs/wasm-gen": "1.9.0",
			  "@webassemblyjs/wasm-opt": "1.9.0",
			  "@webassemblyjs/wasm-parser": "1.9.0",
			  "@webassemblyjs/wast-printer": "1.9.0"
			}
		  },
		  "@webassemblyjs/wasm-gen": {
			"version": "1.9.0",
			"resolved": "https://registry.npmjs.org/@webassemblyjs/wasm-gen/-/wasm-gen-1.9.0.tgz",
			"integrity": "sha512-cPE3o44YzOOHvlsb4+E9qSqjc9Qf9Na1OO/BHFy4OI91XDE14MjFN4lTMezzaIWdPqHnsTodGGNP+iRSYfGkjA==",
			"dev": true,
			"requires": {
			  "@webassemblyjs/ast": "1.9.0",
			  "@webassemblyjs/helper-wasm-bytecode": "1.9.0",
			  "@webassemblyjs/ieee754": "1.9.0",
			  "@webassemblyjs/leb128": "1.9.0",
			  "@webassemblyjs/utf8": "1.9.0"
			}
		  },
		  "@webassemblyjs/wasm-opt": {
			"version": "1.9.0",
			"resolved": "https://registry.npmjs.org/@webassemblyjs/wasm-opt/-/wasm-opt-1.9.0.tgz",
			"integrity": "sha512-Qkjgm6Anhm+OMbIL0iokO7meajkzQD71ioelnfPEj6r4eOFuqm4YC3VBPqXjFyyNwowzbMD+hizmprP/Fwkl2A==",
			"dev": true,
			"requires": {
			  "@webassemblyjs/ast": "1.9.0",
			  "@webassemblyjs/helper-buffer": "1.9.0",
			  "@webassemblyjs/wasm-gen": "1.9.0",
			  "@webassemblyjs/wasm-parser": "1.9.0"
			}
		  },
		  "@webassemblyjs/wasm-parser": {
			"version": "1.9.0",
			"resolved": "https://registry.npmjs.org/@webassemblyjs/wasm-parser/-/wasm-parser-1.9.0.tgz",
			"integrity": "sha512-9+wkMowR2AmdSWQzsPEjFU7njh8HTO5MqO8vjwEHuM+AMHioNqSBONRdr0NQQ3dVQrzp0s8lTcYqzUdb7YgELA==",
			"dev": true,
			"requires": {
			  "@webassemblyjs/ast": "1.9.0",
			  "@webassemblyjs/helper-api-error": "1.9.0",
			  "@webassemblyjs/helper-wasm-bytecode": "1.9.0",
			  "@webassemblyjs/ieee754": "1.9.0",
			  "@webassemblyjs/leb128": "1.9.0",
			  "@webassemblyjs/utf8": "1.9.0"
			}
		  },
		  "@webassemblyjs/wast-parser": {
			"version": "1.9.0",
			"resolved": "https://registry.npmjs.org/@webassemblyjs/wast-parser/-/wast-parser-1.9.0.tgz",
			"integrity": "sha512-qsqSAP3QQ3LyZjNC/0jBJ/ToSxfYJ8kYyuiGvtn/8MK89VrNEfwj7BPQzJVHi0jGTRK2dGdJ5PRqhtjzoww+bw==",
			"dev": true,
			"requires": {
			  "@webassemblyjs/ast": "1.9.0",
			  "@webassemblyjs/floating-point-hex-parser": "1.9.0",
			  "@webassemblyjs/helper-api-error": "1.9.0",
			  "@webassemblyjs/helper-code-frame": "1.9.0",
			  "@webassemblyjs/helper-fsm": "1.9.0",
			  "@xtuc/long": "4.2.2"
			}
		  },
		  "@webassemblyjs/wast-printer": {
			"version": "1.9.0",
			"resolved": "https://registry.npmjs.org/@webassemblyjs/wast-printer/-/wast-printer-1.9.0.tgz",
			"integrity": "sha512-2J0nE95rHXHyQ24cWjMKJ1tqB/ds8z/cyeOZxJhcb+rW+SQASVjuznUSmdz5GpVJTzU8JkhYut0D3siFDD6wsA==",
			"dev": true,
			"requires": {
			  "@webassemblyjs/ast": "1.9.0",
			  "@webassemblyjs/wast-parser": "1.9.0",
			  "@xtuc/long": "4.2.2"
			}
		  },
		  "@xtuc/ieee754": {
			"version": "1.2.0",
			"resolved": "https://registry.npmjs.org/@xtuc/ieee754/-/ieee754-1.2.0.tgz",
			"integrity": "sha512-DX8nKgqcGwsc0eJSqYt5lwP4DH5FlHnmuWWBRy7X0NcaGR0ZtuyeESgMwTYVEtxmsNGY+qit4QYT/MIYTOTPeA==",
			"dev": true
		  },
		  "@xtuc/long": {
			"version": "4.2.2",
			"resolved": "https://registry.npmjs.org/@xtuc/long/-/long-4.2.2.tgz",
			"integrity": "sha512-NuHqBY1PB/D8xU6s/thBgOAiAP7HOYDQ32+BFZILJ8ivkUkAHQnWfn6WhL79Owj1qmUnoN/YPhktdIoucipkAQ==",
			"dev": true
		  },
		  "abab": {
			"version": "2.0.3",
			"resolved": "https://registry.npmjs.org/abab/-/abab-2.0.3.tgz",
			"integrity": "sha512-tsFzPpcttalNjFBCFMqsKYQcWxxen1pgJR56by//QwvJc4/OUS3kPOOttx2tSIfjsylB0pYu7f5D3K1RCxUnUg==",
			"dev": true
		  },
		  "accepts": {
			"version": "1.3.7",
			"resolved": "https://registry.npmjs.org/accepts/-/accepts-1.3.7.tgz",
			"integrity": "sha512-Il80Qs2WjYlJIBNzNkK6KYqlVMTbZLXgHx2oT0pU/fjRHyEp+PEfEPY0R3WCwAGVOtauxh1hOxNgIf5bv7dQpA==",
			"requires": {
			  "mime-types": "~2.1.24",
			  "negotiator": "0.6.2"
			}
		  },
		  "acorn": {
			"version": "6.4.1",
			"resolved": "https://registry.npmjs.org/acorn/-/acorn-6.4.1.tgz",
			"integrity": "sha512-ZVA9k326Nwrj3Cj9jlh3wGFutC2ZornPNARZwsNYqQYgN0EsV2d53w5RN/co65Ohn4sUAUtb1rSUAOD6XN9idA==",
			"dev": true
		  },
		  "acorn-globals": {
			"version": "6.0.0",
			"resolved": "https://registry.npmjs.org/acorn-globals/-/acorn-globals-6.0.0.tgz",
			"integrity": "sha512-ZQl7LOWaF5ePqqcX4hLuv/bLXYQNfNWw2c0/yX/TsPRKamzHcTGQnlCjHT3TsmkOUVEPS3crCxiPfdzE/Trlhg==",
			"dev": true,
			"requires": {
			  "acorn": "^7.1.1",
			  "acorn-walk": "^7.1.1"
			},
			"dependencies": {
			  "acorn": {
				"version": "7.2.0",
				"resolved": "https://registry.npmjs.org/acorn/-/acorn-7.2.0.tgz",
				"integrity": "sha512-apwXVmYVpQ34m/i71vrApRrRKCWQnZZF1+npOD0WV5xZFfwWOmKGQ2RWlfdy9vWITsenisM8M0Qeq8agcFHNiQ==",
				"dev": true
			  }
			}
		  },
		  "acorn-jsx": {
			"version": "5.2.0",
			"resolved": "https://registry.npmjs.org/acorn-jsx/-/acorn-jsx-5.2.0.tgz",
			"integrity": "sha512-HiUX/+K2YpkpJ+SzBffkM/AQ2YE03S0U1kjTLVpoJdhZMOWy8qvXVN9JdLqv2QsaQ6MPYQIuNmwD8zOiYUofLQ==",
			"dev": true
		  },
		  "acorn-walk": {
			"version": "7.1.1",
			"resolved": "https://registry.npmjs.org/acorn-walk/-/acorn-walk-7.1.1.tgz",
			"integrity": "sha512-wdlPY2tm/9XBr7QkKlq0WQVgiuGTX6YWPyRyBviSoScBuLfTVQhvwg6wJ369GJ/1nPfTLMfnrFIfjqVg6d+jQQ==",
			"dev": true
		  },
		  "ajv": {
			"version": "6.10.2",
			"resolved": "https://registry.npmjs.org/ajv/-/ajv-6.10.2.tgz",
			"integrity": "sha512-TXtUUEYHuaTEbLZWIKUr5pmBuhDLy+8KYtPYdcV8qC+pOZL+NKqYwvWSRrVXHn+ZmRRAu8vJTAznH7Oag6RVRw==",
			"dev": true,
			"requires": {
			  "fast-deep-equal": "^2.0.1",
			  "fast-json-stable-stringify": "^2.0.0",
			  "json-schema-traverse": "^0.4.1",
			  "uri-js": "^4.2.2"
			}
		  },
		  "ajv-errors": {
			"version": "1.0.1",
			"resolved": "https://registry.npmjs.org/ajv-errors/-/ajv-errors-1.0.1.tgz",
			"integrity": "sha512-DCRfO/4nQ+89p/RK43i8Ezd41EqdGIU4ld7nGF8OQ14oc/we5rEntLCUa7+jrn3nn83BosfwZA0wb4pon2o8iQ==",
			"dev": true
		  },
		  "ajv-keywords": {
			"version": "3.4.1",
			"resolved": "https://registry.npmjs.org/ajv-keywords/-/ajv-keywords-3.4.1.tgz",
			"integrity": "sha512-RO1ibKvd27e6FEShVFfPALuHI3WjSVNeK5FIsmme/LYRNxjKuNj+Dt7bucLa6NdSv3JcVTyMlm9kGR84z1XpaQ==",
			"dev": true
		  },
		  "ansi-escapes": {
			"version": "4.3.1",
			"resolved": "https://registry.npmjs.org/ansi-escapes/-/ansi-escapes-4.3.1.tgz",
			"integrity": "sha512-JWF7ocqNrp8u9oqpgV+wH5ftbt+cfvv+PTjOvKLT3AdYly/LmORARfEVT1iyjwN+4MqE5UmVKoAdIBqeoCHgLA==",
			"dev": true,
			"requires": {
			  "type-fest": "^0.11.0"
			}
		  },
		  "ansi-regex": {
			"version": "2.1.1",
			"resolved": "https://registry.npmjs.org/ansi-regex/-/ansi-regex-2.1.1.tgz",
			"integrity": "sha1-w7M6te42DYbg5ijwRorn7yfWVN8="
		  },
		  "ansi-styles": {
			"version": "3.2.1",
			"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-3.2.1.tgz",
			"integrity": "sha512-VT0ZI6kZRdTh8YyJw3SMbYm/u+NqfsAxEpWO0Pf9sq8/e94WxxOpPKx9FR1FlyCtOVDNOQ+8ntlqFxiRc+r5qA==",
			"requires": {
			  "color-convert": "^1.9.0"
			}
		  },
		  "anymatch": {
			"version": "3.1.1",
			"resolved": "https://registry.npmjs.org/anymatch/-/anymatch-3.1.1.tgz",
			"integrity": "sha512-mM8522psRCqzV+6LhomX5wgp25YVibjh8Wj23I5RPkPppSVSjyKD2A2mBJmWGa+KN7f2D6LNh9jkBCeyLktzjg==",
			"dev": true,
			"requires": {
			  "normalize-path": "^3.0.0",
			  "picomatch": "^2.0.4"
			}
		  },
		  "append-field": {
			"version": "1.0.0",
			"resolved": "https://registry.npmjs.org/append-field/-/append-field-1.0.0.tgz",
			"integrity": "sha1-HjRA6RXwsSA9I3SOeO3XubW0PlY="
		  },
		  "aproba": {
			"version": "1.2.0",
			"resolved": "https://registry.npmjs.org/aproba/-/aproba-1.2.0.tgz",
			"integrity": "sha512-Y9J6ZjXtoYh8RnXVCMOU/ttDmk1aBjunq9vO0ta5x85WDQiQfUF9sIPBITdbiiIVcBo03Hi3jMxigBtsddlXRw==",
			"dev": true
		  },
		  "arg": {
			"version": "4.1.3",
			"resolved": "https://registry.npmjs.org/arg/-/arg-4.1.3.tgz",
			"integrity": "sha512-58S9QDqG0Xx27YwPSt9fJxivjYl432YCwfDMfZ+71RAqUrZef7LrKQZ3LHLOwCS4FLNBplP533Zx895SeOCHvA==",
			"dev": true
		  },
		  "argparse": {
			"version": "1.0.10",
			"resolved": "https://registry.npmjs.org/argparse/-/argparse-1.0.10.tgz",
			"integrity": "sha512-o5Roy6tNG4SL/FOkCAN6RzjiakZS25RLYFrcMttJqbdd8BWrnA+fGz57iN5Pb06pvBGvl5gQ0B48dJlslXvoTg==",
			"dev": true,
			"requires": {
			  "sprintf-js": "~1.0.2"
			}
		  },
		  "arr-diff": {
			"version": "4.0.0",
			"resolved": "https://registry.npmjs.org/arr-diff/-/arr-diff-4.0.0.tgz",
			"integrity": "sha1-1kYQdP6/7HHn4VI1dhoyml3HxSA=",
			"dev": true
		  },
		  "arr-flatten": {
			"version": "1.1.0",
			"resolved": "https://registry.npmjs.org/arr-flatten/-/arr-flatten-1.1.0.tgz",
			"integrity": "sha512-L3hKV5R/p5o81R7O02IGnwpDmkp6E982XhtbuwSe3O4qOtMMMtodicASA1Cny2U+aCXcNpml+m4dPsvsJ3jatg==",
			"dev": true
		  },
		  "arr-union": {
			"version": "3.1.0",
			"resolved": "https://registry.npmjs.org/arr-union/-/arr-union-3.1.0.tgz",
			"integrity": "sha1-45sJrqne+Gao8gbiiK9jkZuuOcQ=",
			"dev": true
		  },
		  "array-flatten": {
			"version": "1.1.1",
			"resolved": "https://registry.npmjs.org/array-flatten/-/array-flatten-1.1.1.tgz",
			"integrity": "sha1-ml9pkFGx5wczKPKgCJaLZOopVdI="
		  },
		  "array-includes": {
			"version": "3.1.1",
			"resolved": "https://registry.npmjs.org/array-includes/-/array-includes-3.1.1.tgz",
			"integrity": "sha512-c2VXaCHl7zPsvpkFsw4nxvFie4fh1ur9bpcgsVkIjqn0H/Xwdg+7fv3n2r/isyS8EBj5b06M9kHyZuIr4El6WQ==",
			"dev": true,
			"requires": {
			  "define-properties": "^1.1.3",
			  "es-abstract": "^1.17.0",
			  "is-string": "^1.0.5"
			}
		  },
		  "array-unique": {
			"version": "0.3.2",
			"resolved": "https://registry.npmjs.org/array-unique/-/array-unique-0.3.2.tgz",
			"integrity": "sha1-qJS3XUvE9s1nnvMkSp/Y9Gri1Cg=",
			"dev": true
		  },
		  "array.prototype.flat": {
			"version": "1.2.3",
			"resolved": "https://registry.npmjs.org/array.prototype.flat/-/array.prototype.flat-1.2.3.tgz",
			"integrity": "sha512-gBlRZV0VSmfPIeWfuuy56XZMvbVfbEUnOXUvt3F/eUUUSyzlgLxhEX4YAEpxNAogRGehPSnfXyPtYyKAhkzQhQ==",
			"dev": true,
			"requires": {
			  "define-properties": "^1.1.3",
			  "es-abstract": "^1.17.0-next.1"
			}
		  },
		  "asn1": {
			"version": "0.2.4",
			"resolved": "https://registry.npmjs.org/asn1/-/asn1-0.2.4.tgz",
			"integrity": "sha512-jxwzQpLQjSmWXgwaCZE9Nz+glAG01yF1QnWgbhGwHI5A6FRIEY6IVqtHhIepHqI7/kyEyQEagBC5mBEFlIYvdg==",
			"dev": true,
			"requires": {
			  "safer-buffer": "~2.1.0"
			}
		  },
		  "asn1.js": {
			"version": "4.10.1",
			"resolved": "https://registry.npmjs.org/asn1.js/-/asn1.js-4.10.1.tgz",
			"integrity": "sha512-p32cOF5q0Zqs9uBiONKYLm6BClCoBCM5O9JfeUSlnQLBTxYdTK+pW+nXflm8UkKd2UYlEbYz5qEi0JuZR9ckSw==",
			"dev": true,
			"requires": {
			  "bn.js": "^4.0.0",
			  "inherits": "^2.0.1",
			  "minimalistic-assert": "^1.0.0"
			},
			"dependencies": {
			  "bn.js": {
				"version": "4.11.9",
				"resolved": "https://registry.npmjs.org/bn.js/-/bn.js-4.11.9.tgz",
				"integrity": "sha512-E6QoYqCKZfgatHTdHzs1RRKP7ip4vvm+EyRUeE2RF0NblwVvb0p6jSVeNTOFxPn26QXN2o6SMfNxKp6kU8zQaw==",
				"dev": true
			  }
			}
		  },
		  "assert": {
			"version": "1.5.0",
			"resolved": "https://registry.npmjs.org/assert/-/assert-1.5.0.tgz",
			"integrity": "sha512-EDsgawzwoun2CZkCgtxJbv392v4nbk9XDD06zI+kQYoBM/3RBWLlEyJARDOmhAAosBjWACEkKL6S+lIZtcAubA==",
			"dev": true,
			"requires": {
			  "object-assign": "^4.1.1",
			  "util": "0.10.3"
			},
			"dependencies": {
			  "inherits": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/inherits/-/inherits-2.0.1.tgz",
				"integrity": "sha1-sX0I0ya0Qj5Wjv9xn5GwscvfafE=",
				"dev": true
			  },
			  "util": {
				"version": "0.10.3",
				"resolved": "https://registry.npmjs.org/util/-/util-0.10.3.tgz",
				"integrity": "sha1-evsa/lCAUkZInj23/g7TeTNqwPk=",
				"dev": true,
				"requires": {
				  "inherits": "2.0.1"
				}
			  }
			}
		  },
		  "assert-plus": {
			"version": "1.0.0",
			"resolved": "https://registry.npmjs.org/assert-plus/-/assert-plus-1.0.0.tgz",
			"integrity": "sha1-8S4PPF13sLHN2RRpQuTpbB5N1SU=",
			"dev": true
		  },
		  "assign-symbols": {
			"version": "1.0.0",
			"resolved": "https://registry.npmjs.org/assign-symbols/-/assign-symbols-1.0.0.tgz",
			"integrity": "sha1-WWZ/QfrdTyDMvCu5a41Pf3jsA2c=",
			"dev": true
		  },
		  "astral-regex": {
			"version": "1.0.0",
			"resolved": "https://registry.npmjs.org/astral-regex/-/astral-regex-1.0.0.tgz",
			"integrity": "sha512-+Ryf6g3BKoRc7jfp7ad8tM4TtMiaWvbF/1/sQcZPkkS7ag3D5nMBCe2UfOTONtAkaG0tO0ij3C5Lwmf1EiyjHg==",
			"dev": true
		  },
		  "async-each": {
			"version": "1.0.3",
			"resolved": "https://registry.npmjs.org/async-each/-/async-each-1.0.3.tgz",
			"integrity": "sha512-z/WhQ5FPySLdvREByI2vZiTWwCnF0moMJ1hK9YQwDTHKh6I7/uSckMetoRGb5UBZPC1z0jlw+n/XCgjeH7y1AQ==",
			"dev": true,
			"optional": true
		  },
		  "asynckit": {
			"version": "0.4.0",
			"resolved": "https://registry.npmjs.org/asynckit/-/asynckit-0.4.0.tgz",
			"integrity": "sha1-x57Zf380y48robyXkLzDZkdLS3k=",
			"dev": true
		  },
		  "atob": {
			"version": "2.1.2",
			"resolved": "https://registry.npmjs.org/atob/-/atob-2.1.2.tgz",
			"integrity": "sha512-Wm6ukoaOGJi/73p/cl2GvLjTI5JM1k/O14isD73YML8StrH/7/lRFgmg8nICZgD3bZZvjwCGxtMOD3wWNAu8cg==",
			"dev": true
		  },
		  "aws-sign2": {
			"version": "0.7.0",
			"resolved": "https://registry.npmjs.org/aws-sign2/-/aws-sign2-0.7.0.tgz",
			"integrity": "sha1-tG6JCTSpWR8tL2+G1+ap8bP+dqg=",
			"dev": true
		  },
		  "aws4": {
			"version": "1.9.1",
			"resolved": "https://registry.npmjs.org/aws4/-/aws4-1.9.1.tgz",
			"integrity": "sha512-wMHVg2EOHaMRxbzgFJ9gtjOOCrI80OHLG14rxi28XwOW8ux6IiEbRCGGGqCtdAIg4FQCbW20k9RsT4y3gJlFug==",
			"dev": true
		  },
		  "axios": {
			"version": "0.19.2",
			"resolved": "https://registry.npmjs.org/axios/-/axios-0.19.2.tgz",
			"integrity": "sha512-fjgm5MvRHLhx+osE2xoekY70AhARk3a6hkN+3Io1jc00jtquGvxYlKlsFUhmUET0V5te6CcZI7lcv2Ym61mjHA==",
			"requires": {
			  "follow-redirects": "1.5.10"
			}
		  },
		  "babel-jest": {
			"version": "26.0.1",
			"resolved": "https://registry.npmjs.org/babel-jest/-/babel-jest-26.0.1.tgz",
			"integrity": "sha512-Z4GGmSNQ8pX3WS1O+6v3fo41YItJJZsVxG5gIQ+HuB/iuAQBJxMTHTwz292vuYws1LnHfwSRgoqI+nxdy/pcvw==",
			"dev": true,
			"requires": {
			  "@jest/transform": "^26.0.1",
			  "@jest/types": "^26.0.1",
			  "@types/babel__core": "^7.1.7",
			  "babel-plugin-istanbul": "^6.0.0",
			  "babel-preset-jest": "^26.0.0",
			  "chalk": "^4.0.0",
			  "graceful-fs": "^4.2.4",
			  "slash": "^3.0.0"
			},
			"dependencies": {
			  "@jest/types": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/@jest/types/-/types-26.0.1.tgz",
				"integrity": "sha512-IbtjvqI9+eS1qFnOIEL7ggWmT+iK/U+Vde9cGWtYb/b6XgKb3X44ZAe/z9YZzoAAZ/E92m0DqrilF934IGNnQA==",
				"dev": true,
				"requires": {
				  "@types/istanbul-lib-coverage": "^2.0.0",
				  "@types/istanbul-reports": "^1.1.1",
				  "@types/yargs": "^15.0.0",
				  "chalk": "^4.0.0"
				}
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-4.0.0.tgz",
				"integrity": "sha512-N9oWFcegS0sFr9oh1oz2d7Npos6vNoWW9HvtCg5N1KRFpUhaAhvTv5Y58g880fZaEYSNm3qDz8SU1UrGvp+n7A==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "graceful-fs": {
				"version": "4.2.4",
				"resolved": "https://registry.npmjs.org/graceful-fs/-/graceful-fs-4.2.4.tgz",
				"integrity": "sha512-WjKPNJF79dtJAVniUlGGWHYGz2jWxT6VhN/4m1NdkbZ2nOsEF+cI1Edgql5zCRhs/VsQYRvrXctxktVXZUkixw==",
				"dev": true
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "babel-plugin-istanbul": {
			"version": "6.0.0",
			"resolved": "https://registry.npmjs.org/babel-plugin-istanbul/-/babel-plugin-istanbul-6.0.0.tgz",
			"integrity": "sha512-AF55rZXpe7trmEylbaE1Gv54wn6rwU03aptvRoVIGP8YykoSxqdVLV1TfwflBCE/QtHmqtP8SWlTENqbK8GCSQ==",
			"dev": true,
			"requires": {
			  "@babel/helper-plugin-utils": "^7.0.0",
			  "@istanbuljs/load-nyc-config": "^1.0.0",
			  "@istanbuljs/schema": "^0.1.2",
			  "istanbul-lib-instrument": "^4.0.0",
			  "test-exclude": "^6.0.0"
			}
		  },
		  "babel-plugin-jest-hoist": {
			"version": "26.0.0",
			"resolved": "https://registry.npmjs.org/babel-plugin-jest-hoist/-/babel-plugin-jest-hoist-26.0.0.tgz",
			"integrity": "sha512-+AuoehOrjt9irZL7DOt2+4ZaTM6dlu1s5TTS46JBa0/qem4dy7VNW3tMb96qeEqcIh20LD73TVNtmVEeymTG7w==",
			"dev": true,
			"requires": {
			  "@babel/template": "^7.3.3",
			  "@babel/types": "^7.3.3",
			  "@types/babel__traverse": "^7.0.6"
			}
		  },
		  "babel-preset-current-node-syntax": {
			"version": "0.1.2",
			"resolved": "https://registry.npmjs.org/babel-preset-current-node-syntax/-/babel-preset-current-node-syntax-0.1.2.tgz",
			"integrity": "sha512-u/8cS+dEiK1SFILbOC8/rUI3ml9lboKuuMvZ/4aQnQmhecQAgPw5ew066C1ObnEAUmlx7dv/s2z52psWEtLNiw==",
			"dev": true,
			"requires": {
			  "@babel/plugin-syntax-async-generators": "^7.8.4",
			  "@babel/plugin-syntax-bigint": "^7.8.3",
			  "@babel/plugin-syntax-class-properties": "^7.8.3",
			  "@babel/plugin-syntax-json-strings": "^7.8.3",
			  "@babel/plugin-syntax-logical-assignment-operators": "^7.8.3",
			  "@babel/plugin-syntax-nullish-coalescing-operator": "^7.8.3",
			  "@babel/plugin-syntax-numeric-separator": "^7.8.3",
			  "@babel/plugin-syntax-object-rest-spread": "^7.8.3",
			  "@babel/plugin-syntax-optional-catch-binding": "^7.8.3",
			  "@babel/plugin-syntax-optional-chaining": "^7.8.3"
			}
		  },
		  "babel-preset-jest": {
			"version": "26.0.0",
			"resolved": "https://registry.npmjs.org/babel-preset-jest/-/babel-preset-jest-26.0.0.tgz",
			"integrity": "sha512-9ce+DatAa31DpR4Uir8g4Ahxs5K4W4L8refzt+qHWQANb6LhGcAEfIFgLUwk67oya2cCUd6t4eUMtO/z64ocNw==",
			"dev": true,
			"requires": {
			  "babel-plugin-jest-hoist": "^26.0.0",
			  "babel-preset-current-node-syntax": "^0.1.2"
			}
		  },
		  "balanced-match": {
			"version": "1.0.0",
			"resolved": "https://registry.npmjs.org/balanced-match/-/balanced-match-1.0.0.tgz",
			"integrity": "sha1-ibTRmasr7kneFk6gK4nORi1xt2c="
		  },
		  "base": {
			"version": "0.11.2",
			"resolved": "https://registry.npmjs.org/base/-/base-0.11.2.tgz",
			"integrity": "sha512-5T6P4xPgpp0YDFvSWwEZ4NoE3aM4QBQXDzmVbraCkFj8zHM+mba8SyqB5DbZWyR7mYHo6Y7BdQo3MoA4m0TeQg==",
			"dev": true,
			"requires": {
			  "cache-base": "^1.0.1",
			  "class-utils": "^0.3.5",
			  "component-emitter": "^1.2.1",
			  "define-property": "^1.0.0",
			  "isobject": "^3.0.1",
			  "mixin-deep": "^1.2.0",
			  "pascalcase": "^0.1.1"
			},
			"dependencies": {
			  "define-property": {
				"version": "1.0.0",
				"resolved": "https://registry.npmjs.org/define-property/-/define-property-1.0.0.tgz",
				"integrity": "sha1-dp66rz9KY6rTr56NMEybvnm/sOY=",
				"dev": true,
				"requires": {
				  "is-descriptor": "^1.0.0"
				}
			  },
			  "is-accessor-descriptor": {
				"version": "1.0.0",
				"resolved": "https://registry.npmjs.org/is-accessor-descriptor/-/is-accessor-descriptor-1.0.0.tgz",
				"integrity": "sha512-m5hnHTkcVsPfqx3AKlyttIPb7J+XykHvJP2B9bZDjlhLIoEq4XoK64Vg7boZlVWYK6LUY94dYPEE7Lh0ZkZKcQ==",
				"dev": true,
				"requires": {
				  "kind-of": "^6.0.0"
				}
			  },
			  "is-data-descriptor": {
				"version": "1.0.0",
				"resolved": "https://registry.npmjs.org/is-data-descriptor/-/is-data-descriptor-1.0.0.tgz",
				"integrity": "sha512-jbRXy1FmtAoCjQkVmIVYwuuqDFUbaOeDjmed1tOGPrsMhtJA4rD9tkgA0F1qJ3gRFRXcHYVkdeaP50Q5rE/jLQ==",
				"dev": true,
				"requires": {
				  "kind-of": "^6.0.0"
				}
			  },
			  "is-descriptor": {
				"version": "1.0.2",
				"resolved": "https://registry.npmjs.org/is-descriptor/-/is-descriptor-1.0.2.tgz",
				"integrity": "sha512-2eis5WqQGV7peooDyLmNEPUrps9+SXX5c9pL3xEB+4e9HnGuDa7mB7kHxHw4CbqS9k1T2hOH3miL8n8WtiYVtg==",
				"dev": true,
				"requires": {
				  "is-accessor-descriptor": "^1.0.0",
				  "is-data-descriptor": "^1.0.0",
				  "kind-of": "^6.0.2"
				}
			  }
			}
		  },
		  "base64-js": {
			"version": "1.3.1",
			"resolved": "https://registry.npmjs.org/base64-js/-/base64-js-1.3.1.tgz",
			"integrity": "sha512-mLQ4i2QO1ytvGWFWmcngKO//JXAQueZvwEKtjgQFM4jIK0kU+ytMfplL8j+n5mspOfjHwoAg+9yhb7BwAHm36g==",
			"dev": true
		  },
		  "bcrypt-pbkdf": {
			"version": "1.0.2",
			"resolved": "https://registry.npmjs.org/bcrypt-pbkdf/-/bcrypt-pbkdf-1.0.2.tgz",
			"integrity": "sha1-pDAdOJtqQ/m2f/PKEaP2Y342Dp4=",
			"dev": true,
			"requires": {
			  "tweetnacl": "^0.14.3"
			}
		  },
		  "big.js": {
			"version": "5.2.2",
			"resolved": "https://registry.npmjs.org/big.js/-/big.js-5.2.2.tgz",
			"integrity": "sha512-vyL2OymJxmarO8gxMr0mhChsO9QGwhynfuu4+MHTAW6czfq9humCB7rKpUjDd9YUiDPU4mzpyupFSvOClAwbmQ==",
			"dev": true
		  },
		  "binary-extensions": {
			"version": "2.0.0",
			"resolved": "https://registry.npmjs.org/binary-extensions/-/binary-extensions-2.0.0.tgz",
			"integrity": "sha512-Phlt0plgpIIBOGTT/ehfFnbNlfsDEiqmzE2KRXoX1bLIlir4X/MR+zSyBEkL05ffWgnRSf/DXv+WrUAVr93/ow==",
			"dev": true
		  },
		  "bindings": {
			"version": "1.5.0",
			"resolved": "https://registry.npmjs.org/bindings/-/bindings-1.5.0.tgz",
			"integrity": "sha512-p2q/t/mhvuOj/UeLlV6566GD/guowlr0hHxClI0W9m7MWYkL1F0hLo+0Aexs9HSPCtR1SXQ0TD3MMKrXZajbiQ==",
			"dev": true,
			"optional": true,
			"requires": {
			  "file-uri-to-path": "1.0.0"
			}
		  },
		  "bluebird": {
			"version": "3.7.2",
			"resolved": "https://registry.npmjs.org/bluebird/-/bluebird-3.7.2.tgz",
			"integrity": "sha512-XpNj6GDQzdfW+r2Wnn7xiSAd7TM3jzkxGXBGTtWKuSXv1xUV+azxAm8jdWZN06QTQk+2N2XB9jRDkvbmQmcRtg==",
			"dev": true
		  },
		  "bn.js": {
			"version": "5.1.2",
			"resolved": "https://registry.npmjs.org/bn.js/-/bn.js-5.1.2.tgz",
			"integrity": "sha512-40rZaf3bUNKTVYu9sIeeEGOg7g14Yvnj9kH7b50EiwX0Q7A6umbvfI5tvHaOERH0XigqKkfLkFQxzb4e6CIXnA==",
			"dev": true
		  },
		  "body-parser": {
			"version": "1.19.0",
			"resolved": "https://registry.npmjs.org/body-parser/-/body-parser-1.19.0.tgz",
			"integrity": "sha512-dhEPs72UPbDnAQJ9ZKMNTP6ptJaionhP5cBb541nXPlW60Jepo9RV/a4fX4XWW9CuFNK22krhrj1+rgzifNCsw==",
			"requires": {
			  "bytes": "3.1.0",
			  "content-type": "~1.0.4",
			  "debug": "2.6.9",
			  "depd": "~1.1.2",
			  "http-errors": "1.7.2",
			  "iconv-lite": "0.4.24",
			  "on-finished": "~2.3.0",
			  "qs": "6.7.0",
			  "raw-body": "2.4.0",
			  "type-is": "~1.6.17"
			},
			"dependencies": {
			  "debug": {
				"version": "2.6.9",
				"resolved": "https://registry.npmjs.org/debug/-/debug-2.6.9.tgz",
				"integrity": "sha512-bC7ElrdJaJnPbAP+1EotYvqZsb3ecl5wi6Bfi6BJTUcNowp6cvspg0jXznRTKDjm/E7AdgFBVeAPVMNcKGsHMA==",
				"requires": {
				  "ms": "2.0.0"
				}
			  }
			}
		  },
		  "brace-expansion": {
			"version": "1.1.11",
			"resolved": "https://registry.npmjs.org/brace-expansion/-/brace-expansion-1.1.11.tgz",
			"integrity": "sha512-iCuPHDFgrHX7H2vEI/5xpz07zSHB00TpugqhmYtVmMO6518mCuRMoOYFldEBl0g187ufozdaHgWKcYFb61qGiA==",
			"requires": {
			  "balanced-match": "^1.0.0",
			  "concat-map": "0.0.1"
			}
		  },
		  "braces": {
			"version": "3.0.2",
			"resolved": "https://registry.npmjs.org/braces/-/braces-3.0.2.tgz",
			"integrity": "sha512-b8um+L1RzM3WDSzvhm6gIz1yfTbBt6YTlcEKAvsmqCZZFw46z626lVj9j1yEPW33H5H+lBQpZMP1k8l+78Ha0A==",
			"dev": true,
			"requires": {
			  "fill-range": "^7.0.1"
			}
		  },
		  "brorand": {
			"version": "1.1.0",
			"resolved": "https://registry.npmjs.org/brorand/-/brorand-1.1.0.tgz",
			"integrity": "sha1-EsJe/kCkXjwyPrhnWgoM5XsiNx8=",
			"dev": true
		  },
		  "browser-process-hrtime": {
			"version": "1.0.0",
			"resolved": "https://registry.npmjs.org/browser-process-hrtime/-/browser-process-hrtime-1.0.0.tgz",
			"integrity": "sha512-9o5UecI3GhkpM6DrXr69PblIuWxPKk9Y0jHBRhdocZ2y7YECBFCsHm79Pr3OyR2AvjhDkabFJaDJMYRazHgsow==",
			"dev": true
		  },
		  "browserify-aes": {
			"version": "1.2.0",
			"resolved": "https://registry.npmjs.org/browserify-aes/-/browserify-aes-1.2.0.tgz",
			"integrity": "sha512-+7CHXqGuspUn/Sl5aO7Ea0xWGAtETPXNSAjHo48JfLdPWcMng33Xe4znFvQweqc/uzk5zSOI3H52CYnjCfb5hA==",
			"dev": true,
			"requires": {
			  "buffer-xor": "^1.0.3",
			  "cipher-base": "^1.0.0",
			  "create-hash": "^1.1.0",
			  "evp_bytestokey": "^1.0.3",
			  "inherits": "^2.0.1",
			  "safe-buffer": "^5.0.1"
			}
		  },
		  "browserify-cipher": {
			"version": "1.0.1",
			"resolved": "https://registry.npmjs.org/browserify-cipher/-/browserify-cipher-1.0.1.tgz",
			"integrity": "sha512-sPhkz0ARKbf4rRQt2hTpAHqn47X3llLkUGn+xEJzLjwY8LRs2p0v7ljvI5EyoRO/mexrNunNECisZs+gw2zz1w==",
			"dev": true,
			"requires": {
			  "browserify-aes": "^1.0.4",
			  "browserify-des": "^1.0.0",
			  "evp_bytestokey": "^1.0.0"
			}
		  },
		  "browserify-des": {
			"version": "1.0.2",
			"resolved": "https://registry.npmjs.org/browserify-des/-/browserify-des-1.0.2.tgz",
			"integrity": "sha512-BioO1xf3hFwz4kc6iBhI3ieDFompMhrMlnDFC4/0/vd5MokpuAc3R+LYbwTA9A5Yc9pq9UYPqffKpW2ObuwX5A==",
			"dev": true,
			"requires": {
			  "cipher-base": "^1.0.1",
			  "des.js": "^1.0.0",
			  "inherits": "^2.0.1",
			  "safe-buffer": "^5.1.2"
			}
		  },
		  "browserify-rsa": {
			"version": "4.0.1",
			"resolved": "https://registry.npmjs.org/browserify-rsa/-/browserify-rsa-4.0.1.tgz",
			"integrity": "sha1-IeCr+vbyApzy+vsTNWenAdQTVSQ=",
			"dev": true,
			"requires": {
			  "bn.js": "^4.1.0",
			  "randombytes": "^2.0.1"
			},
			"dependencies": {
			  "bn.js": {
				"version": "4.11.9",
				"resolved": "https://registry.npmjs.org/bn.js/-/bn.js-4.11.9.tgz",
				"integrity": "sha512-E6QoYqCKZfgatHTdHzs1RRKP7ip4vvm+EyRUeE2RF0NblwVvb0p6jSVeNTOFxPn26QXN2o6SMfNxKp6kU8zQaw==",
				"dev": true
			  }
			}
		  },
		  "browserify-sign": {
			"version": "4.2.0",
			"resolved": "https://registry.npmjs.org/browserify-sign/-/browserify-sign-4.2.0.tgz",
			"integrity": "sha512-hEZC1KEeYuoHRqhGhTy6gWrpJA3ZDjFWv0DE61643ZnOXAKJb3u7yWcrU0mMc9SwAqK1n7myPGndkp0dFG7NFA==",
			"dev": true,
			"requires": {
			  "bn.js": "^5.1.1",
			  "browserify-rsa": "^4.0.1",
			  "create-hash": "^1.2.0",
			  "create-hmac": "^1.1.7",
			  "elliptic": "^6.5.2",
			  "inherits": "^2.0.4",
			  "parse-asn1": "^5.1.5",
			  "readable-stream": "^3.6.0",
			  "safe-buffer": "^5.2.0"
			},
			"dependencies": {
			  "inherits": {
				"version": "2.0.4",
				"resolved": "https://registry.npmjs.org/inherits/-/inherits-2.0.4.tgz",
				"integrity": "sha512-k/vGaX4/Yla3WzyMCvTQOXYeIHvqOKtnqBduzTHpzpQZzAskKMhZ2K+EnBiSM9zGSoIFeMpXKxa4dYeZIQqewQ==",
				"dev": true
			  },
			  "readable-stream": {
				"version": "3.6.0",
				"resolved": "https://registry.npmjs.org/readable-stream/-/readable-stream-3.6.0.tgz",
				"integrity": "sha512-BViHy7LKeTz4oNnkcLJ+lVSL6vpiFeX6/d3oSH8zCW7UxP2onchk+vTGB143xuFjHS3deTgkKoXXymXqymiIdA==",
				"dev": true,
				"requires": {
				  "inherits": "^2.0.3",
				  "string_decoder": "^1.1.1",
				  "util-deprecate": "^1.0.1"
				}
			  },
			  "safe-buffer": {
				"version": "5.2.1",
				"resolved": "https://registry.npmjs.org/safe-buffer/-/safe-buffer-5.2.1.tgz",
				"integrity": "sha512-rp3So07KcdmmKbGvgaNxQSJr7bGVSVk5S9Eq1F+ppbRo70+YeaDxkw5Dd8NPN+GD6bjnYm2VuPuCXmpuYvmCXQ==",
				"dev": true
			  },
			  "string_decoder": {
				"version": "1.3.0",
				"resolved": "https://registry.npmjs.org/string_decoder/-/string_decoder-1.3.0.tgz",
				"integrity": "sha512-hkRX8U1WjJFd8LsDJ2yQ/wWWxaopEsABU1XfkM8A+j0+85JAGppt16cr1Whg6KIbb4okU6Mql6BOj+uup/wKeA==",
				"dev": true,
				"requires": {
				  "safe-buffer": "~5.2.0"
				}
			  }
			}
		  },
		  "browserify-zlib": {
			"version": "0.2.0",
			"resolved": "https://registry.npmjs.org/browserify-zlib/-/browserify-zlib-0.2.0.tgz",
			"integrity": "sha512-Z942RysHXmJrhqk88FmKBVq/v5tqmSkDz7p54G/MGyjMnCFFnC79XWNbg+Vta8W6Wb2qtSZTSxIGkJrRpCFEiA==",
			"dev": true,
			"requires": {
			  "pako": "~1.0.5"
			}
		  },
		  "bs-logger": {
			"version": "0.2.6",
			"resolved": "https://registry.npmjs.org/bs-logger/-/bs-logger-0.2.6.tgz",
			"integrity": "sha512-pd8DCoxmbgc7hyPKOvxtqNcjYoOsABPQdcCUjGp3d42VR2CX1ORhk2A87oqqu5R1kk+76nsxZupkmyd+MVtCog==",
			"dev": true,
			"requires": {
			  "fast-json-stable-stringify": "2.x"
			}
		  },
		  "bser": {
			"version": "2.1.1",
			"resolved": "https://registry.npmjs.org/bser/-/bser-2.1.1.tgz",
			"integrity": "sha512-gQxTNE/GAfIIrmHLUE3oJyp5FO6HRBfhjnw4/wMmA63ZGDJnWBmgY/lyQBpnDUkGmAhbSe39tx2d/iTOAfglwQ==",
			"dev": true,
			"requires": {
			  "node-int64": "^0.4.0"
			}
		  },
		  "buffer": {
			"version": "4.9.2",
			"resolved": "https://registry.npmjs.org/buffer/-/buffer-4.9.2.tgz",
			"integrity": "sha512-xq+q3SRMOxGivLhBNaUdC64hDTQwejJ+H0T/NB1XMtTVEwNTrfFF3gAxiyW0Bu/xWEGhjVKgUcMhCrUy2+uCWg==",
			"dev": true,
			"requires": {
			  "base64-js": "^1.0.2",
			  "ieee754": "^1.1.4",
			  "isarray": "^1.0.0"
			},
			"dependencies": {
			  "isarray": {
				"version": "1.0.0",
				"resolved": "https://registry.npmjs.org/isarray/-/isarray-1.0.0.tgz",
				"integrity": "sha1-u5NdSFgsuhaMBoNJV6VKPgcSTxE=",
				"dev": true
			  }
			}
		  },
		  "buffer-from": {
			"version": "1.1.1",
			"resolved": "https://registry.npmjs.org/buffer-from/-/buffer-from-1.1.1.tgz",
			"integrity": "sha512-MQcXEUbCKtEo7bhqEs6560Hyd4XaovZlO/k9V3hjVUF/zwW7KBVdSK4gIt/bzwS9MbR5qob+F5jusZsb0YQK2A=="
		  },
		  "buffer-xor": {
			"version": "1.0.3",
			"resolved": "https://registry.npmjs.org/buffer-xor/-/buffer-xor-1.0.3.tgz",
			"integrity": "sha1-JuYe0UIvtw3ULm42cp7VHYVf6Nk=",
			"dev": true
		  },
		  "builtin-status-codes": {
			"version": "3.0.0",
			"resolved": "https://registry.npmjs.org/builtin-status-codes/-/builtin-status-codes-3.0.0.tgz",
			"integrity": "sha1-hZgoeOIbmOHGZCXgPQF0eI9Wnug=",
			"dev": true
		  },
		  "busboy": {
			"version": "0.2.14",
			"resolved": "https://registry.npmjs.org/busboy/-/busboy-0.2.14.tgz",
			"integrity": "sha1-bCpiLvz0fFe7vh4qnDetNseSVFM=",
			"requires": {
			  "dicer": "0.2.5",
			  "readable-stream": "1.1.x"
			}
		  },
		  "bytes": {
			"version": "3.1.0",
			"resolved": "https://registry.npmjs.org/bytes/-/bytes-3.1.0.tgz",
			"integrity": "sha512-zauLjrfCG+xvoyaqLoV8bLVXXNGC4JqlxFCutSDWA6fJrTo2ZuvLYTqZ7aHBLZSMOopbzwv8f+wZcVzfVTI2Dg=="
		  },
		  "cacache": {
			"version": "12.0.4",
			"resolved": "https://registry.npmjs.org/cacache/-/cacache-12.0.4.tgz",
			"integrity": "sha512-a0tMB40oefvuInr4Cwb3GerbL9xTj1D5yg0T5xrjGCGyfvbxseIXX7BAO/u/hIXdafzOI5JC3wDwHyf24buOAQ==",
			"dev": true,
			"requires": {
			  "bluebird": "^3.5.5",
			  "chownr": "^1.1.1",
			  "figgy-pudding": "^3.5.1",
			  "glob": "^7.1.4",
			  "graceful-fs": "^4.1.15",
			  "infer-owner": "^1.0.3",
			  "lru-cache": "^5.1.1",
			  "mississippi": "^3.0.0",
			  "mkdirp": "^0.5.1",
			  "move-concurrently": "^1.0.1",
			  "promise-inflight": "^1.0.1",
			  "rimraf": "^2.6.3",
			  "ssri": "^6.0.1",
			  "unique-filename": "^1.1.1",
			  "y18n": "^4.0.0"
			},
			"dependencies": {
			  "rimraf": {
				"version": "2.7.1",
				"resolved": "https://registry.npmjs.org/rimraf/-/rimraf-2.7.1.tgz",
				"integrity": "sha512-uWjbaKIK3T1OSVptzX7Nl6PvQ3qAGtKEtVRjRuazjfL3Bx5eI409VZSqgND+4UNnmzLVdPj9FqFJNPqBZFve4w==",
				"dev": true,
				"requires": {
				  "glob": "^7.1.3"
				}
			  }
			}
		  },
		  "cache-base": {
			"version": "1.0.1",
			"resolved": "https://registry.npmjs.org/cache-base/-/cache-base-1.0.1.tgz",
			"integrity": "sha512-AKcdTnFSWATd5/GCPRxr2ChwIJ85CeyrEyjRHlKxQ56d4XJMGym0uAiKn0xbLOGOl3+yRpOTi484dVCEc5AUzQ==",
			"dev": true,
			"requires": {
			  "collection-visit": "^1.0.0",
			  "component-emitter": "^1.2.1",
			  "get-value": "^2.0.6",
			  "has-value": "^1.0.0",
			  "isobject": "^3.0.1",
			  "set-value": "^2.0.0",
			  "to-object-path": "^0.3.0",
			  "union-value": "^1.0.0",
			  "unset-value": "^1.0.0"
			}
		  },
		  "callsites": {
			"version": "3.1.0",
			"resolved": "https://registry.npmjs.org/callsites/-/callsites-3.1.0.tgz",
			"integrity": "sha512-P8BjAsXvZS+VIDUI11hHCQEv74YT67YUi5JJFNWIqL235sBmjX4+qx9Muvls5ivyNENctx46xQLQ3aTuE7ssaQ==",
			"dev": true
		  },
		  "camelcase": {
			"version": "5.3.1",
			"resolved": "https://registry.npmjs.org/camelcase/-/camelcase-5.3.1.tgz",
			"integrity": "sha512-L28STB170nwWS63UjtlEOE3dldQApaJXZkOI1uMFfzf3rRuPegHaHesyee+YxQ+W6SvRDQV6UrdOdRiR153wJg==",
			"dev": true
		  },
		  "capture-exit": {
			"version": "2.0.0",
			"resolved": "https://registry.npmjs.org/capture-exit/-/capture-exit-2.0.0.tgz",
			"integrity": "sha512-PiT/hQmTonHhl/HFGN+Lx3JJUznrVYJ3+AQsnthneZbvW7x+f08Tk7yLJTLEOUvBTbduLeeBkxEaYXUOUrRq6g==",
			"dev": true,
			"requires": {
			  "rsvp": "^4.8.4"
			}
		  },
		  "caseless": {
			"version": "0.12.0",
			"resolved": "https://registry.npmjs.org/caseless/-/caseless-0.12.0.tgz",
			"integrity": "sha1-G2gcIf+EAzyCZUMJBolCDRhxUdw=",
			"dev": true
		  },
		  "chalk": {
			"version": "2.4.2",
			"resolved": "https://registry.npmjs.org/chalk/-/chalk-2.4.2.tgz",
			"integrity": "sha512-Mti+f9lpJNcwF4tWV8/OrTTtF1gZi+f8FqlyAdouralcFWFQWF2+NgCHShjkCb+IFBLq9buZwE1xckQU4peSuQ==",
			"requires": {
			  "ansi-styles": "^3.2.1",
			  "escape-string-regexp": "^1.0.5",
			  "supports-color": "^5.3.0"
			}
		  },
		  "char-regex": {
			"version": "1.0.2",
			"resolved": "https://registry.npmjs.org/char-regex/-/char-regex-1.0.2.tgz",
			"integrity": "sha512-kWWXztvZ5SBQV+eRgKFeh8q5sLuZY2+8WUIzlxWVTg+oGwY14qylx1KbKzHd8P6ZYkAg0xyIDU9JMHhyJMZ1jw==",
			"dev": true
		  },
		  "chardet": {
			"version": "0.7.0",
			"resolved": "https://registry.npmjs.org/chardet/-/chardet-0.7.0.tgz",
			"integrity": "sha512-mT8iDcrh03qDGRRmoA2hmBJnxpllMR+0/0qlzjqZES6NdiWDcZkCNAk4rPFZ9Q85r27unkiNNg8ZOiwZXBHwcA==",
			"dev": true
		  },
		  "chokidar": {
			"version": "3.4.0",
			"resolved": "https://registry.npmjs.org/chokidar/-/chokidar-3.4.0.tgz",
			"integrity": "sha512-aXAaho2VJtisB/1fg1+3nlLJqGOuewTzQpd/Tz0yTg2R0e4IGtshYvtjowyEumcBv2z+y4+kc75Mz7j5xJskcQ==",
			"dev": true,
			"requires": {
			  "anymatch": "~3.1.1",
			  "braces": "~3.0.2",
			  "fsevents": "~2.1.2",
			  "glob-parent": "~5.1.0",
			  "is-binary-path": "~2.1.0",
			  "is-glob": "~4.0.1",
			  "normalize-path": "~3.0.0",
			  "readdirp": "~3.4.0"
			}
		  },
		  "chownr": {
			"version": "1.1.4",
			"resolved": "https://registry.npmjs.org/chownr/-/chownr-1.1.4.tgz",
			"integrity": "sha512-jJ0bqzaylmJtVnNgzTeSOs8DPavpbYgEr/b0YL8/2GO3xJEhInFmhKMUnEJQjZumK7KXGFhUy89PrsJWlakBVg==",
			"dev": true
		  },
		  "chrome-trace-event": {
			"version": "1.0.2",
			"resolved": "https://registry.npmjs.org/chrome-trace-event/-/chrome-trace-event-1.0.2.tgz",
			"integrity": "sha512-9e/zx1jw7B4CO+c/RXoCsfg/x1AfUBioy4owYH0bJprEYAx5hRFLRhWBqHAG57D0ZM4H7vxbP7bPe0VwhQRYDQ==",
			"dev": true,
			"requires": {
			  "tslib": "^1.9.0"
			}
		  },
		  "ci-info": {
			"version": "2.0.0",
			"resolved": "https://registry.npmjs.org/ci-info/-/ci-info-2.0.0.tgz",
			"integrity": "sha512-5tK7EtrZ0N+OLFMthtqOj4fI2Jeb88C4CAZPu25LDVUgXJ0A3Js4PMGqrn0JU1W0Mh1/Z8wZzYPxqUrXeBboCQ==",
			"dev": true
		  },
		  "cipher-base": {
			"version": "1.0.4",
			"resolved": "https://registry.npmjs.org/cipher-base/-/cipher-base-1.0.4.tgz",
			"integrity": "sha512-Kkht5ye6ZGmwv40uUDZztayT2ThLQGfnj/T71N/XzeZeo3nf8foyW7zGTsPYkEya3m5f3cAypH+qe7YOrM1U2Q==",
			"dev": true,
			"requires": {
			  "inherits": "^2.0.1",
			  "safe-buffer": "^5.0.1"
			}
		  },
		  "class-utils": {
			"version": "0.3.6",
			"resolved": "https://registry.npmjs.org/class-utils/-/class-utils-0.3.6.tgz",
			"integrity": "sha512-qOhPa/Fj7s6TY8H8esGu5QNpMMQxz79h+urzrNYN6mn+9BnxlDGf5QZ+XeCDsxSjPqsSR56XOZOJmpeurnLMeg==",
			"dev": true,
			"requires": {
			  "arr-union": "^3.1.0",
			  "define-property": "^0.2.5",
			  "isobject": "^3.0.0",
			  "static-extend": "^0.1.1"
			},
			"dependencies": {
			  "define-property": {
				"version": "0.2.5",
				"resolved": "https://registry.npmjs.org/define-property/-/define-property-0.2.5.tgz",
				"integrity": "sha1-w1se+RjsPJkPmlvFe+BKrOxcgRY=",
				"dev": true,
				"requires": {
				  "is-descriptor": "^0.1.0"
				}
			  }
			}
		  },
		  "cli-color": {
			"version": "2.0.0",
			"resolved": "https://registry.npmjs.org/cli-color/-/cli-color-2.0.0.tgz",
			"integrity": "sha512-a0VZ8LeraW0jTuCkuAGMNufareGHhyZU9z8OGsW0gXd1hZGi1SRuNRXdbGkraBBKnhyUhyebFWnRbp+dIn0f0A==",
			"requires": {
			  "ansi-regex": "^2.1.1",
			  "d": "^1.0.1",
			  "es5-ext": "^0.10.51",
			  "es6-iterator": "^2.0.3",
			  "memoizee": "^0.4.14",
			  "timers-ext": "^0.1.7"
			}
		  },
		  "cli-cursor": {
			"version": "3.1.0",
			"resolved": "https://registry.npmjs.org/cli-cursor/-/cli-cursor-3.1.0.tgz",
			"integrity": "sha512-I/zHAwsKf9FqGoXM4WWRACob9+SNukZTd94DWF57E4toouRulbCxcUh6RKUEOQlYTHJnzkPMySvPNaaSLNfLZw==",
			"dev": true,
			"requires": {
			  "restore-cursor": "^3.1.0"
			}
		  },
		  "cli-spinners": {
			"version": "2.2.0",
			"resolved": "https://registry.npmjs.org/cli-spinners/-/cli-spinners-2.2.0.tgz",
			"integrity": "sha512-tgU3fKwzYjiLEQgPMD9Jt+JjHVL9kW93FiIMX/l7rivvOD4/LL0Mf7gda3+4U2KJBloybwgj5KEoQgGRioMiKQ==",
			"dev": true
		  },
		  "cli-table3": {
			"version": "0.5.1",
			"resolved": "https://registry.npmjs.org/cli-table3/-/cli-table3-0.5.1.tgz",
			"integrity": "sha512-7Qg2Jrep1S/+Q3EceiZtQcDPWxhAvBw+ERf1162v4sikJrvojMHFqXt8QIVha8UlH9rgU0BeWPytZ9/TzYqlUw==",
			"dev": true,
			"requires": {
			  "colors": "^1.1.2",
			  "object-assign": "^4.1.0",
			  "string-width": "^2.1.1"
			},
			"dependencies": {
			  "ansi-regex": {
				"version": "3.0.0",
				"resolved": "https://registry.npmjs.org/ansi-regex/-/ansi-regex-3.0.0.tgz",
				"integrity": "sha1-7QMXwyIGT3lGbAKWa922Bas32Zg=",
				"dev": true
			  },
			  "is-fullwidth-code-point": {
				"version": "2.0.0",
				"resolved": "https://registry.npmjs.org/is-fullwidth-code-point/-/is-fullwidth-code-point-2.0.0.tgz",
				"integrity": "sha1-o7MKXE8ZkYMWeqq5O+764937ZU8=",
				"dev": true
			  },
			  "string-width": {
				"version": "2.1.1",
				"resolved": "https://registry.npmjs.org/string-width/-/string-width-2.1.1.tgz",
				"integrity": "sha512-nOqH59deCq9SRHlxq1Aw85Jnt4w6KvLKqWVik6oA9ZklXLNIOlqg4F2yrT1MVaTjAqvVwdfeZ7w7aCvJD7ugkw==",
				"dev": true,
				"requires": {
				  "is-fullwidth-code-point": "^2.0.0",
				  "strip-ansi": "^4.0.0"
				}
			  },
			  "strip-ansi": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/strip-ansi/-/strip-ansi-4.0.0.tgz",
				"integrity": "sha1-qEeQIusaw2iocTibY1JixQXuNo8=",
				"dev": true,
				"requires": {
				  "ansi-regex": "^3.0.0"
				}
			  }
			}
		  },
		  "cli-width": {
			"version": "2.2.1",
			"resolved": "https://registry.npmjs.org/cli-width/-/cli-width-2.2.1.tgz",
			"integrity": "sha512-GRMWDxpOB6Dgk2E5Uo+3eEBvtOOlimMmpbFiKuLFnQzYDavtLFY3K5ona41jgN/WdRZtG7utuVSVTL4HbZHGkw==",
			"dev": true
		  },
		  "cliui": {
			"version": "6.0.0",
			"resolved": "https://registry.npmjs.org/cliui/-/cliui-6.0.0.tgz",
			"integrity": "sha512-t6wbgtoCXvAzst7QgXxJYqPt0usEfbgQdftEPbLL/cvv6HPE5VgvqCuAIDR0NgU52ds6rFwqrgakNLrHEjCbrQ==",
			"dev": true,
			"requires": {
			  "string-width": "^4.2.0",
			  "strip-ansi": "^6.0.0",
			  "wrap-ansi": "^6.2.0"
			},
			"dependencies": {
			  "ansi-regex": {
				"version": "5.0.0",
				"resolved": "https://registry.npmjs.org/ansi-regex/-/ansi-regex-5.0.0.tgz",
				"integrity": "sha512-bY6fj56OUQ0hU1KjFNDQuJFezqKdrAyFdIevADiqrWHwSlbmBNMHp5ak2f40Pm8JTFyM2mqxkG6ngkHO11f/lg==",
				"dev": true
			  },
			  "strip-ansi": {
				"version": "6.0.0",
				"resolved": "https://registry.npmjs.org/strip-ansi/-/strip-ansi-6.0.0.tgz",
				"integrity": "sha512-AuvKTrTfQNYNIctbR1K/YGTR1756GycPsg7b9bdV9Duqur4gv6aKqHXah67Z8ImS7WEz5QVcOtlfW2rZEugt6w==",
				"dev": true,
				"requires": {
				  "ansi-regex": "^5.0.0"
				}
			  }
			}
		  },
		  "clone": {
			"version": "1.0.4",
			"resolved": "https://registry.npmjs.org/clone/-/clone-1.0.4.tgz",
			"integrity": "sha1-2jCcwmPfFZlMaIypAheco8fNfH4=",
			"dev": true
		  },
		  "co": {
			"version": "4.6.0",
			"resolved": "https://registry.npmjs.org/co/-/co-4.6.0.tgz",
			"integrity": "sha1-bqa989hTrlTMuOR7+gvz+QMfsYQ=",
			"dev": true
		  },
		  "collect-v8-coverage": {
			"version": "1.0.1",
			"resolved": "https://registry.npmjs.org/collect-v8-coverage/-/collect-v8-coverage-1.0.1.tgz",
			"integrity": "sha512-iBPtljfCNcTKNAto0KEtDfZ3qzjJvqE3aTGZsbhjSBlorqpXJlaWWtPO35D+ZImoC3KWejX64o+yPGxhWSTzfg==",
			"dev": true
		  },
		  "collection-visit": {
			"version": "1.0.0",
			"resolved": "https://registry.npmjs.org/collection-visit/-/collection-visit-1.0.0.tgz",
			"integrity": "sha1-S8A3PBZLwykbTTaMgpzxqApZ3KA=",
			"dev": true,
			"requires": {
			  "map-visit": "^1.0.0",
			  "object-visit": "^1.0.0"
			}
		  },
		  "color-convert": {
			"version": "1.9.3",
			"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-1.9.3.tgz",
			"integrity": "sha512-QfAUtd+vFdAtFQcC8CCyYt1fYWxSqAiK2cSD6zDB8N3cpsEBAvRxp9zOGg6G/SHHJYAT88/az/IuDGALsNVbGg==",
			"requires": {
			  "color-name": "1.1.3"
			}
		  },
		  "color-name": {
			"version": "1.1.3",
			"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.3.tgz",
			"integrity": "sha1-p9BVi9icQveV3UIyj3QIMcpTvCU="
		  },
		  "colors": {
			"version": "1.4.0",
			"resolved": "https://registry.npmjs.org/colors/-/colors-1.4.0.tgz",
			"integrity": "sha512-a+UqTh4kgZg/SlGvfbzDHpgRu7AAQOmmqRHJnxhRZICKFUT91brVhNNt58CMWU9PsBbv3PDCZUHbVxuDiH2mtA==",
			"dev": true,
			"optional": true
		  },
		  "combined-stream": {
			"version": "1.0.8",
			"resolved": "https://registry.npmjs.org/combined-stream/-/combined-stream-1.0.8.tgz",
			"integrity": "sha512-FQN4MRfuJeHf7cBbBMJFXhKSDq+2kAArBlmRBvcvFE5BB1HZKXtSFASDhdlz9zOYwxh8lDdnvmMOe/+5cdoEdg==",
			"dev": true,
			"requires": {
			  "delayed-stream": "~1.0.0"
			}
		  },
		  "commander": {
			"version": "4.1.1",
			"resolved": "https://registry.npmjs.org/commander/-/commander-4.1.1.tgz",
			"integrity": "sha512-NOKm8xhkzAjzFx8B2v5OAHT+u5pRQc2UCa2Vq9jYL/31o2wi9mxBA7LIFs3sV5VSC49z6pEhfbMULvShKj26WA==",
			"dev": true
		  },
		  "commondir": {
			"version": "1.0.1",
			"resolved": "https://registry.npmjs.org/commondir/-/commondir-1.0.1.tgz",
			"integrity": "sha1-3dgA2gxmEnOTzKWVDqloo6rxJTs=",
			"dev": true
		  },
		  "component-emitter": {
			"version": "1.3.0",
			"resolved": "https://registry.npmjs.org/component-emitter/-/component-emitter-1.3.0.tgz",
			"integrity": "sha512-Rd3se6QB+sO1TwqZjscQrurpEPIfO0/yYnSin6Q/rD3mOutHvUrCAhJub3r90uNb+SESBuE0QYoB90YdfatsRg==",
			"dev": true
		  },
		  "concat-map": {
			"version": "0.0.1",
			"resolved": "https://registry.npmjs.org/concat-map/-/concat-map-0.0.1.tgz",
			"integrity": "sha1-2Klr13/Wjfd5OnMDajug1UBdR3s="
		  },
		  "concat-stream": {
			"version": "1.6.2",
			"resolved": "https://registry.npmjs.org/concat-stream/-/concat-stream-1.6.2.tgz",
			"integrity": "sha512-27HBghJxjiZtIk3Ycvn/4kbJk/1uZuJFfuPEns6LaEvpvG1f0hTea8lilrouyo9mVc2GWdcEZ8OLoGmSADlrCw==",
			"requires": {
			  "buffer-from": "^1.0.0",
			  "inherits": "^2.0.3",
			  "readable-stream": "^2.2.2",
			  "typedarray": "^0.0.6"
			},
			"dependencies": {
			  "isarray": {
				"version": "1.0.0",
				"resolved": "https://registry.npmjs.org/isarray/-/isarray-1.0.0.tgz",
				"integrity": "sha1-u5NdSFgsuhaMBoNJV6VKPgcSTxE="
			  },
			  "readable-stream": {
				"version": "2.3.7",
				"resolved": "https://registry.npmjs.org/readable-stream/-/readable-stream-2.3.7.tgz",
				"integrity": "sha512-Ebho8K4jIbHAxnuxi7o42OrZgF/ZTNcsZj6nRKyUmkhLFq8CHItp/fy6hQZuZmP/n3yZ9VBUbp4zz/mX8hmYPw==",
				"requires": {
				  "core-util-is": "~1.0.0",
				  "inherits": "~2.0.3",
				  "isarray": "~1.0.0",
				  "process-nextick-args": "~2.0.0",
				  "safe-buffer": "~5.1.1",
				  "string_decoder": "~1.1.1",
				  "util-deprecate": "~1.0.1"
				}
			  },
			  "string_decoder": {
				"version": "1.1.1",
				"resolved": "https://registry.npmjs.org/string_decoder/-/string_decoder-1.1.1.tgz",
				"integrity": "sha512-n/ShnvDi6FHbbVfviro+WojiFzv+s8MPMHBczVePfUpDJLwoLT0ht1l4YwBCbi8pJAveEEdnkHyPyTP/mzRfwg==",
				"requires": {
				  "safe-buffer": "~5.1.0"
				}
			  }
			}
		  },
		  "consola": {
			"version": "2.12.2",
			"resolved": "https://registry.npmjs.org/consola/-/consola-2.12.2.tgz",
			"integrity": "sha512-c9mzemrAk57s3UIjepn8KKkuEH5fauMdot5kFSJUnqHcnApVS9Db8Rbv5AZ1Iz6lXzaGe9z1crQXhJtGX4h/Og=="
		  },
		  "console-browserify": {
			"version": "1.2.0",
			"resolved": "https://registry.npmjs.org/console-browserify/-/console-browserify-1.2.0.tgz",
			"integrity": "sha512-ZMkYO/LkF17QvCPqM0gxw8yUzigAOZOSWSHg91FH6orS7vcEj5dVZTidN2fQ14yBSdg97RqhSNwLUXInd52OTA==",
			"dev": true
		  },
		  "constants-browserify": {
			"version": "1.0.0",
			"resolved": "https://registry.npmjs.org/constants-browserify/-/constants-browserify-1.0.0.tgz",
			"integrity": "sha1-wguW2MYXdIqvHBYCF2DNJ/y4y3U=",
			"dev": true
		  },
		  "contains-path": {
			"version": "0.1.0",
			"resolved": "https://registry.npmjs.org/contains-path/-/contains-path-0.1.0.tgz",
			"integrity": "sha1-/ozxhP9mcLa67wGp1IYaXL7EEgo=",
			"dev": true
		  },
		  "content-disposition": {
			"version": "0.5.3",
			"resolved": "https://registry.npmjs.org/content-disposition/-/content-disposition-0.5.3.tgz",
			"integrity": "sha512-ExO0774ikEObIAEV9kDo50o+79VCUdEB6n6lzKgGwupcVeRlhrj3qGAfwq8G6uBJjkqLrhT0qEYFcWng8z1z0g==",
			"requires": {
			  "safe-buffer": "5.1.2"
			}
		  },
		  "content-type": {
			"version": "1.0.4",
			"resolved": "https://registry.npmjs.org/content-type/-/content-type-1.0.4.tgz",
			"integrity": "sha512-hIP3EEPs8tB9AT1L+NUqtwOAps4mk2Zob89MWXMHjHWg9milF/j4osnnQLXBCBFBk/tvIG/tUc9mOUJiPBhPXA=="
		  },
		  "convert-source-map": {
			"version": "1.7.0",
			"resolved": "https://registry.npmjs.org/convert-source-map/-/convert-source-map-1.7.0.tgz",
			"integrity": "sha512-4FJkXzKXEDB1snCFZlLP4gpC3JILicCpGbzG9f9G7tGqGCzETQ2hWPrcinA9oU4wtf2biUaEH5065UnMeR33oA==",
			"dev": true,
			"requires": {
			  "safe-buffer": "~5.1.1"
			}
		  },
		  "cookie": {
			"version": "0.4.0",
			"resolved": "https://registry.npmjs.org/cookie/-/cookie-0.4.0.tgz",
			"integrity": "sha512-+Hp8fLp57wnUSt0tY0tHEXh4voZRDnoIrZPqlo3DPiI4y9lwg/jqx+1Om94/W6ZaPDOUbnjOt/99w66zk+l1Xg=="
		  },
		  "cookie-signature": {
			"version": "1.0.6",
			"resolved": "https://registry.npmjs.org/cookie-signature/-/cookie-signature-1.0.6.tgz",
			"integrity": "sha1-4wOogrNCzD7oylE6eZmXNNqzriw="
		  },
		  "cookiejar": {
			"version": "2.1.2",
			"resolved": "https://registry.npmjs.org/cookiejar/-/cookiejar-2.1.2.tgz",
			"integrity": "sha512-Mw+adcfzPxcPeI+0WlvRrr/3lGVO0bD75SxX6811cxSh1Wbxx7xZBGK1eVtDf6si8rg2lhnUjsVLMFMfbRIuwA==",
			"dev": true
		  },
		  "copy-concurrently": {
			"version": "1.0.5",
			"resolved": "https://registry.npmjs.org/copy-concurrently/-/copy-concurrently-1.0.5.tgz",
			"integrity": "sha512-f2domd9fsVDFtaFcbaRZuYXwtdmnzqbADSwhSWYxYB/Q8zsdUUFMXVRwXGDMWmbEzAn1kdRrtI1T/KTFOL4X2A==",
			"dev": true,
			"requires": {
			  "aproba": "^1.1.1",
			  "fs-write-stream-atomic": "^1.0.8",
			  "iferr": "^0.1.5",
			  "mkdirp": "^0.5.1",
			  "rimraf": "^2.5.4",
			  "run-queue": "^1.0.0"
			},
			"dependencies": {
			  "rimraf": {
				"version": "2.7.1",
				"resolved": "https://registry.npmjs.org/rimraf/-/rimraf-2.7.1.tgz",
				"integrity": "sha512-uWjbaKIK3T1OSVptzX7Nl6PvQ3qAGtKEtVRjRuazjfL3Bx5eI409VZSqgND+4UNnmzLVdPj9FqFJNPqBZFve4w==",
				"dev": true,
				"requires": {
				  "glob": "^7.1.3"
				}
			  }
			}
		  },
		  "copy-descriptor": {
			"version": "0.1.1",
			"resolved": "https://registry.npmjs.org/copy-descriptor/-/copy-descriptor-0.1.1.tgz",
			"integrity": "sha1-Z29us8OZl8LuGsOpJP1hJHSPV40=",
			"dev": true
		  },
		  "core-util-is": {
			"version": "1.0.2",
			"resolved": "https://registry.npmjs.org/core-util-is/-/core-util-is-1.0.2.tgz",
			"integrity": "sha1-tf1UIgqivFq1eqtxQMlAdUUDwac="
		  },
		  "cors": {
			"version": "2.8.5",
			"resolved": "https://registry.npmjs.org/cors/-/cors-2.8.5.tgz",
			"integrity": "sha512-KIHbLJqu73RGr/hnbrO9uBeixNGuvSQjul/jdFvS/KFSIH1hWVd1ng7zOHx+YrEfInLG7q4n6GHQ9cDtxv/P6g==",
			"requires": {
			  "object-assign": "^4",
			  "vary": "^1"
			}
		  },
		  "create-ecdh": {
			"version": "4.0.3",
			"resolved": "https://registry.npmjs.org/create-ecdh/-/create-ecdh-4.0.3.tgz",
			"integrity": "sha512-GbEHQPMOswGpKXM9kCWVrremUcBmjteUaQ01T9rkKCPDXfUHX0IoP9LpHYo2NPFampa4e+/pFDc3jQdxrxQLaw==",
			"dev": true,
			"requires": {
			  "bn.js": "^4.1.0",
			  "elliptic": "^6.0.0"
			},
			"dependencies": {
			  "bn.js": {
				"version": "4.11.9",
				"resolved": "https://registry.npmjs.org/bn.js/-/bn.js-4.11.9.tgz",
				"integrity": "sha512-E6QoYqCKZfgatHTdHzs1RRKP7ip4vvm+EyRUeE2RF0NblwVvb0p6jSVeNTOFxPn26QXN2o6SMfNxKp6kU8zQaw==",
				"dev": true
			  }
			}
		  },
		  "create-hash": {
			"version": "1.2.0",
			"resolved": "https://registry.npmjs.org/create-hash/-/create-hash-1.2.0.tgz",
			"integrity": "sha512-z00bCGNHDG8mHAkP7CtT1qVu+bFQUPjYq/4Iv3C3kWjTFV10zIjfSoeqXo9Asws8gwSHDGj/hl2u4OGIjapeCg==",
			"dev": true,
			"requires": {
			  "cipher-base": "^1.0.1",
			  "inherits": "^2.0.1",
			  "md5.js": "^1.3.4",
			  "ripemd160": "^2.0.1",
			  "sha.js": "^2.4.0"
			}
		  },
		  "create-hmac": {
			"version": "1.1.7",
			"resolved": "https://registry.npmjs.org/create-hmac/-/create-hmac-1.1.7.tgz",
			"integrity": "sha512-MJG9liiZ+ogc4TzUwuvbER1JRdgvUFSB5+VR/g5h82fGaIRWMWddtKBHi7/sVhfjQZ6SehlyhvQYrcYkaUIpLg==",
			"dev": true,
			"requires": {
			  "cipher-base": "^1.0.3",
			  "create-hash": "^1.1.0",
			  "inherits": "^2.0.1",
			  "ripemd160": "^2.0.0",
			  "safe-buffer": "^5.0.1",
			  "sha.js": "^2.4.8"
			}
		  },
		  "cross-spawn": {
			"version": "6.0.5",
			"resolved": "https://registry.npmjs.org/cross-spawn/-/cross-spawn-6.0.5.tgz",
			"integrity": "sha512-eTVLrBSt7fjbDygz805pMnstIs2VTBNkRm0qxZd+M7A5XDdxVRWO5MxGBXZhjY4cqLYLdtrGqRf8mBPmzwSpWQ==",
			"dev": true,
			"requires": {
			  "nice-try": "^1.0.4",
			  "path-key": "^2.0.1",
			  "semver": "^5.5.0",
			  "shebang-command": "^1.2.0",
			  "which": "^1.2.9"
			}
		  },
		  "crypto-browserify": {
			"version": "3.12.0",
			"resolved": "https://registry.npmjs.org/crypto-browserify/-/crypto-browserify-3.12.0.tgz",
			"integrity": "sha512-fz4spIh+znjO2VjL+IdhEpRJ3YN6sMzITSBijk6FK2UvTqruSQW+/cCZTSNsMiZNvUeq0CqurF+dAbyiGOY6Wg==",
			"dev": true,
			"requires": {
			  "browserify-cipher": "^1.0.0",
			  "browserify-sign": "^4.0.0",
			  "create-ecdh": "^4.0.0",
			  "create-hash": "^1.1.0",
			  "create-hmac": "^1.1.0",
			  "diffie-hellman": "^5.0.0",
			  "inherits": "^2.0.1",
			  "pbkdf2": "^3.0.3",
			  "public-encrypt": "^4.0.0",
			  "randombytes": "^2.0.0",
			  "randomfill": "^1.0.3"
			}
		  },
		  "cssom": {
			"version": "0.4.4",
			"resolved": "https://registry.npmjs.org/cssom/-/cssom-0.4.4.tgz",
			"integrity": "sha512-p3pvU7r1MyyqbTk+WbNJIgJjG2VmTIaB10rI93LzVPrmDJKkzKYMtxxyAvQXR/NS6otuzveI7+7BBq3SjBS2mw==",
			"dev": true
		  },
		  "cssstyle": {
			"version": "2.3.0",
			"resolved": "https://registry.npmjs.org/cssstyle/-/cssstyle-2.3.0.tgz",
			"integrity": "sha512-AZL67abkUzIuvcHqk7c09cezpGNcxUxU4Ioi/05xHk4DQeTkWmGYftIE6ctU6AEt+Gn4n1lDStOtj7FKycP71A==",
			"dev": true,
			"requires": {
			  "cssom": "~0.3.6"
			},
			"dependencies": {
			  "cssom": {
				"version": "0.3.8",
				"resolved": "https://registry.npmjs.org/cssom/-/cssom-0.3.8.tgz",
				"integrity": "sha512-b0tGHbfegbhPJpxpiBPU2sCkigAqtM9O121le6bbOlgyV+NyGyCmVfJ6QW9eRjz8CpNfWEOYBIMIGRYkLwsIYg==",
				"dev": true
			  }
			}
		  },
		  "cyclist": {
			"version": "1.0.1",
			"resolved": "https://registry.npmjs.org/cyclist/-/cyclist-1.0.1.tgz",
			"integrity": "sha1-WW6WmP0MgOEgOMK4LW6xs1tiJNk=",
			"dev": true
		  },
		  "d": {
			"version": "1.0.1",
			"resolved": "https://registry.npmjs.org/d/-/d-1.0.1.tgz",
			"integrity": "sha512-m62ShEObQ39CfralilEQRjH6oAMtNCV1xJyEx5LpRYUVN+EviphDgUc/F3hnYbADmkiNs67Y+3ylmlG7Lnu+FA==",
			"requires": {
			  "es5-ext": "^0.10.50",
			  "type": "^1.0.1"
			}
		  },
		  "dashdash": {
			"version": "1.14.1",
			"resolved": "https://registry.npmjs.org/dashdash/-/dashdash-1.14.1.tgz",
			"integrity": "sha1-hTz6D3y+L+1d4gMmuN1YEDX24vA=",
			"dev": true,
			"requires": {
			  "assert-plus": "^1.0.0"
			}
		  },
		  "data-urls": {
			"version": "2.0.0",
			"resolved": "https://registry.npmjs.org/data-urls/-/data-urls-2.0.0.tgz",
			"integrity": "sha512-X5eWTSXO/BJmpdIKCRuKUgSCgAN0OwliVK3yPKbwIWU1Tdw5BRajxlzMidvh+gwko9AfQ9zIj52pzF91Q3YAvQ==",
			"dev": true,
			"requires": {
			  "abab": "^2.0.3",
			  "whatwg-mimetype": "^2.3.0",
			  "whatwg-url": "^8.0.0"
			}
		  },
		  "debug": {
			"version": "3.1.0",
			"resolved": "https://registry.npmjs.org/debug/-/debug-3.1.0.tgz",
			"integrity": "sha512-OX8XqP7/1a9cqkxYw2yXss15f26NKWBpDXQd0/uK/KPqdQhxbPa994hnzjcE2VqQpDslf55723cKPUOGSmMY3g==",
			"requires": {
			  "ms": "2.0.0"
			}
		  },
		  "decamelize": {
			"version": "1.2.0",
			"resolved": "https://registry.npmjs.org/decamelize/-/decamelize-1.2.0.tgz",
			"integrity": "sha1-9lNNFRSCabIDUue+4m9QH5oZEpA=",
			"dev": true
		  },
		  "decimal.js": {
			"version": "10.2.0",
			"resolved": "https://registry.npmjs.org/decimal.js/-/decimal.js-10.2.0.tgz",
			"integrity": "sha512-vDPw+rDgn3bZe1+F/pyEwb1oMG2XTlRVgAa6B4KccTEpYgF8w6eQllVbQcfIJnZyvzFtFpxnpGtx8dd7DJp/Rw==",
			"dev": true
		  },
		  "decode-uri-component": {
			"version": "0.2.0",
			"resolved": "https://registry.npmjs.org/decode-uri-component/-/decode-uri-component-0.2.0.tgz",
			"integrity": "sha1-6zkTMzRYd1y4TNGh+uBiEGu4dUU=",
			"dev": true
		  },
		  "deep-is": {
			"version": "0.1.3",
			"resolved": "https://registry.npmjs.org/deep-is/-/deep-is-0.1.3.tgz",
			"integrity": "sha1-s2nW+128E+7PUk+RsHD+7cNXzzQ=",
			"dev": true
		  },
		  "deepmerge": {
			"version": "4.2.2",
			"resolved": "https://registry.npmjs.org/deepmerge/-/deepmerge-4.2.2.tgz",
			"integrity": "sha512-FJ3UgI4gIl+PHZm53knsuSFpE+nESMr7M4v9QcgB7S63Kj/6WqMiFQJpBBYz1Pt+66bZpP3Q7Lye0Oo9MPKEdg==",
			"dev": true
		  },
		  "defaults": {
			"version": "1.0.3",
			"resolved": "https://registry.npmjs.org/defaults/-/defaults-1.0.3.tgz",
			"integrity": "sha1-xlYFHpgX2f8I7YgUd/P+QBnz730=",
			"dev": true,
			"requires": {
			  "clone": "^1.0.2"
			}
		  },
		  "define-properties": {
			"version": "1.1.3",
			"resolved": "https://registry.npmjs.org/define-properties/-/define-properties-1.1.3.tgz",
			"integrity": "sha512-3MqfYKj2lLzdMSf8ZIZE/V+Zuy+BgD6f164e8K2w7dgnpKArBDerGYpM46IYYcjnkdPNMjPk9A6VFB8+3SKlXQ==",
			"dev": true,
			"requires": {
			  "object-keys": "^1.0.12"
			}
		  },
		  "define-property": {
			"version": "2.0.2",
			"resolved": "https://registry.npmjs.org/define-property/-/define-property-2.0.2.tgz",
			"integrity": "sha512-jwK2UV4cnPpbcG7+VRARKTZPUWowwXA8bzH5NP6ud0oeAxyYPuGZUAC7hMugpCdz4BeSZl2Dl9k66CHJ/46ZYQ==",
			"dev": true,
			"requires": {
			  "is-descriptor": "^1.0.2",
			  "isobject": "^3.0.1"
			},
			"dependencies": {
			  "is-accessor-descriptor": {
				"version": "1.0.0",
				"resolved": "https://registry.npmjs.org/is-accessor-descriptor/-/is-accessor-descriptor-1.0.0.tgz",
				"integrity": "sha512-m5hnHTkcVsPfqx3AKlyttIPb7J+XykHvJP2B9bZDjlhLIoEq4XoK64Vg7boZlVWYK6LUY94dYPEE7Lh0ZkZKcQ==",
				"dev": true,
				"requires": {
				  "kind-of": "^6.0.0"
				}
			  },
			  "is-data-descriptor": {
				"version": "1.0.0",
				"resolved": "https://registry.npmjs.org/is-data-descriptor/-/is-data-descriptor-1.0.0.tgz",
				"integrity": "sha512-jbRXy1FmtAoCjQkVmIVYwuuqDFUbaOeDjmed1tOGPrsMhtJA4rD9tkgA0F1qJ3gRFRXcHYVkdeaP50Q5rE/jLQ==",
				"dev": true,
				"requires": {
				  "kind-of": "^6.0.0"
				}
			  },
			  "is-descriptor": {
				"version": "1.0.2",
				"resolved": "https://registry.npmjs.org/is-descriptor/-/is-descriptor-1.0.2.tgz",
				"integrity": "sha512-2eis5WqQGV7peooDyLmNEPUrps9+SXX5c9pL3xEB+4e9HnGuDa7mB7kHxHw4CbqS9k1T2hOH3miL8n8WtiYVtg==",
				"dev": true,
				"requires": {
				  "is-accessor-descriptor": "^1.0.0",
				  "is-data-descriptor": "^1.0.0",
				  "kind-of": "^6.0.2"
				}
			  }
			}
		  },
		  "delayed-stream": {
			"version": "1.0.0",
			"resolved": "https://registry.npmjs.org/delayed-stream/-/delayed-stream-1.0.0.tgz",
			"integrity": "sha1-3zrhmayt+31ECqrgsp4icrJOxhk=",
			"dev": true
		  },
		  "depd": {
			"version": "1.1.2",
			"resolved": "https://registry.npmjs.org/depd/-/depd-1.1.2.tgz",
			"integrity": "sha1-m81S4UwJd2PnSbJ0xDRu0uVgtak="
		  },
		  "des.js": {
			"version": "1.0.1",
			"resolved": "https://registry.npmjs.org/des.js/-/des.js-1.0.1.tgz",
			"integrity": "sha512-Q0I4pfFrv2VPd34/vfLrFOoRmlYj3OV50i7fskps1jZWK1kApMWWT9G6RRUeYedLcBDIhnSDaUvJMb3AhUlaEA==",
			"dev": true,
			"requires": {
			  "inherits": "^2.0.1",
			  "minimalistic-assert": "^1.0.0"
			}
		  },
		  "destroy": {
			"version": "1.0.4",
			"resolved": "https://registry.npmjs.org/destroy/-/destroy-1.0.4.tgz",
			"integrity": "sha1-l4hXRCxEdJ5CBmE+N5RiBYJqvYA="
		  },
		  "detect-newline": {
			"version": "3.1.0",
			"resolved": "https://registry.npmjs.org/detect-newline/-/detect-newline-3.1.0.tgz",
			"integrity": "sha512-TLz+x/vEXm/Y7P7wn1EJFNLxYpUD4TgMosxY6fAVJUnJMbupHBOncxyWUG9OpTaH9EBD7uFI5LfEgmMOc54DsA==",
			"dev": true
		  },
		  "dicer": {
			"version": "0.2.5",
			"resolved": "https://registry.npmjs.org/dicer/-/dicer-0.2.5.tgz",
			"integrity": "sha1-WZbAhrszIYyBLAkL3cCc0S+stw8=",
			"requires": {
			  "readable-stream": "1.1.x",
			  "streamsearch": "0.1.2"
			}
		  },
		  "diff": {
			"version": "4.0.2",
			"resolved": "https://registry.npmjs.org/diff/-/diff-4.0.2.tgz",
			"integrity": "sha512-58lmxKSA4BNyLz+HHMUzlOEpg09FV+ev6ZMe3vJihgdxzgcwZ8VoEEPmALCZG9LmqfVoNMMKpttIYTVG6uDY7A==",
			"dev": true
		  },
		  "diff-sequences": {
			"version": "25.2.6",
			"resolved": "https://registry.npmjs.org/diff-sequences/-/diff-sequences-25.2.6.tgz",
			"integrity": "sha512-Hq8o7+6GaZeoFjtpgvRBUknSXNeJiCx7V9Fr94ZMljNiCr9n9L8H8aJqgWOQiDDGdyn29fRNcDdRVJ5fdyihfg==",
			"dev": true
		  },
		  "diffie-hellman": {
			"version": "5.0.3",
			"resolved": "https://registry.npmjs.org/diffie-hellman/-/diffie-hellman-5.0.3.tgz",
			"integrity": "sha512-kqag/Nl+f3GwyK25fhUMYj81BUOrZ9IuJsjIcDE5icNM9FJHAVm3VcUDxdLPoQtTuUylWm6ZIknYJwwaPxsUzg==",
			"dev": true,
			"requires": {
			  "bn.js": "^4.1.0",
			  "miller-rabin": "^4.0.0",
			  "randombytes": "^2.0.0"
			},
			"dependencies": {
			  "bn.js": {
				"version": "4.11.9",
				"resolved": "https://registry.npmjs.org/bn.js/-/bn.js-4.11.9.tgz",
				"integrity": "sha512-E6QoYqCKZfgatHTdHzs1RRKP7ip4vvm+EyRUeE2RF0NblwVvb0p6jSVeNTOFxPn26QXN2o6SMfNxKp6kU8zQaw==",
				"dev": true
			  }
			}
		  },
		  "doctrine": {
			"version": "3.0.0",
			"resolved": "https://registry.npmjs.org/doctrine/-/doctrine-3.0.0.tgz",
			"integrity": "sha512-yS+Q5i3hBf7GBkd4KG8a7eBNNWNGLTaEwwYWUijIYM7zrlYDM0BFXHjjPWlWZ1Rg7UaddZeIDmi9jF3HmqiQ2w==",
			"dev": true,
			"requires": {
			  "esutils": "^2.0.2"
			}
		  },
		  "domain-browser": {
			"version": "1.2.0",
			"resolved": "https://registry.npmjs.org/domain-browser/-/domain-browser-1.2.0.tgz",
			"integrity": "sha512-jnjyiM6eRyZl2H+W8Q/zLMA481hzi0eszAaBUzIVnmYVDBbnLxVNnfu1HgEBvCbL+71FrxMl3E6lpKH7Ge3OXA==",
			"dev": true
		  },
		  "domexception": {
			"version": "2.0.1",
			"resolved": "https://registry.npmjs.org/domexception/-/domexception-2.0.1.tgz",
			"integrity": "sha512-yxJ2mFy/sibVQlu5qHjOkf9J3K6zgmCxgJ94u2EdvDOV09H+32LtRswEcUsmUWN72pVLOEnTSRaIVVzVQgS0dg==",
			"dev": true,
			"requires": {
			  "webidl-conversions": "^5.0.0"
			},
			"dependencies": {
			  "webidl-conversions": {
				"version": "5.0.0",
				"resolved": "https://registry.npmjs.org/webidl-conversions/-/webidl-conversions-5.0.0.tgz",
				"integrity": "sha512-VlZwKPCkYKxQgeSbH5EyngOmRp7Ww7I9rQLERETtf5ofd9pGeswWiOtogpEO850jziPRarreGxn5QIiTqpb2wA==",
				"dev": true
			  }
			}
		  },
		  "duplexify": {
			"version": "3.7.1",
			"resolved": "https://registry.npmjs.org/duplexify/-/duplexify-3.7.1.tgz",
			"integrity": "sha512-07z8uv2wMyS51kKhD1KsdXJg5WQ6t93RneqRxUHnskXVtlYYkLqM0gqStQZ3pj073g687jPCHrqNfCzawLYh5g==",
			"dev": true,
			"requires": {
			  "end-of-stream": "^1.0.0",
			  "inherits": "^2.0.1",
			  "readable-stream": "^2.0.0",
			  "stream-shift": "^1.0.0"
			},
			"dependencies": {
			  "isarray": {
				"version": "1.0.0",
				"resolved": "https://registry.npmjs.org/isarray/-/isarray-1.0.0.tgz",
				"integrity": "sha1-u5NdSFgsuhaMBoNJV6VKPgcSTxE=",
				"dev": true
			  },
			  "readable-stream": {
				"version": "2.3.7",
				"resolved": "https://registry.npmjs.org/readable-stream/-/readable-stream-2.3.7.tgz",
				"integrity": "sha512-Ebho8K4jIbHAxnuxi7o42OrZgF/ZTNcsZj6nRKyUmkhLFq8CHItp/fy6hQZuZmP/n3yZ9VBUbp4zz/mX8hmYPw==",
				"dev": true,
				"requires": {
				  "core-util-is": "~1.0.0",
				  "inherits": "~2.0.3",
				  "isarray": "~1.0.0",
				  "process-nextick-args": "~2.0.0",
				  "safe-buffer": "~5.1.1",
				  "string_decoder": "~1.1.1",
				  "util-deprecate": "~1.0.1"
				}
			  },
			  "string_decoder": {
				"version": "1.1.1",
				"resolved": "https://registry.npmjs.org/string_decoder/-/string_decoder-1.1.1.tgz",
				"integrity": "sha512-n/ShnvDi6FHbbVfviro+WojiFzv+s8MPMHBczVePfUpDJLwoLT0ht1l4YwBCbi8pJAveEEdnkHyPyTP/mzRfwg==",
				"dev": true,
				"requires": {
				  "safe-buffer": "~5.1.0"
				}
			  }
			}
		  },
		  "ecc-jsbn": {
			"version": "0.1.2",
			"resolved": "https://registry.npmjs.org/ecc-jsbn/-/ecc-jsbn-0.1.2.tgz",
			"integrity": "sha1-OoOpBOVDUyh4dMVkt1SThoSamMk=",
			"dev": true,
			"requires": {
			  "jsbn": "~0.1.0",
			  "safer-buffer": "^2.1.0"
			}
		  },
		  "ee-first": {
			"version": "1.1.1",
			"resolved": "https://registry.npmjs.org/ee-first/-/ee-first-1.1.1.tgz",
			"integrity": "sha1-WQxhFWsK4vTwJVcyoViyZrxWsh0="
		  },
		  "elliptic": {
			"version": "6.5.3",
			"resolved": "https://registry.npmjs.org/elliptic/-/elliptic-6.5.3.tgz",
			"integrity": "sha512-IMqzv5wNQf+E6aHeIqATs0tOLeOTwj1QKbRcS3jBbYkl5oLAserA8yJTT7/VyHUYG91PRmPyeQDObKLPpeS4dw==",
			"dev": true,
			"requires": {
			  "bn.js": "^4.4.0",
			  "brorand": "^1.0.1",
			  "hash.js": "^1.0.0",
			  "hmac-drbg": "^1.0.0",
			  "inherits": "^2.0.1",
			  "minimalistic-assert": "^1.0.0",
			  "minimalistic-crypto-utils": "^1.0.0"
			},
			"dependencies": {
			  "bn.js": {
				"version": "4.11.9",
				"resolved": "https://registry.npmjs.org/bn.js/-/bn.js-4.11.9.tgz",
				"integrity": "sha512-E6QoYqCKZfgatHTdHzs1RRKP7ip4vvm+EyRUeE2RF0NblwVvb0p6jSVeNTOFxPn26QXN2o6SMfNxKp6kU8zQaw==",
				"dev": true
			  }
			}
		  },
		  "emoji-regex": {
			"version": "8.0.0",
			"resolved": "https://registry.npmjs.org/emoji-regex/-/emoji-regex-8.0.0.tgz",
			"integrity": "sha512-MSjYzcWNOA0ewAHpz0MxpYFvwg6yjy1NG3xteoqz644VCo/RPgnr1/GGt+ic3iJTzQ8Eu3TdM14SawnVUmGE6A==",
			"dev": true
		  },
		  "emojis-list": {
			"version": "3.0.0",
			"resolved": "https://registry.npmjs.org/emojis-list/-/emojis-list-3.0.0.tgz",
			"integrity": "sha512-/kyM18EfinwXZbno9FyUGeFh87KC8HRQBQGildHZbEuRyWFOmv1U10o9BBp8XVZDVNNuQKyIGIu5ZYAAXJ0V2Q==",
			"dev": true
		  },
		  "encodeurl": {
			"version": "1.0.2",
			"resolved": "https://registry.npmjs.org/encodeurl/-/encodeurl-1.0.2.tgz",
			"integrity": "sha1-rT/0yG7C0CkyL1oCw6mmBslbP1k="
		  },
		  "end-of-stream": {
			"version": "1.4.4",
			"resolved": "https://registry.npmjs.org/end-of-stream/-/end-of-stream-1.4.4.tgz",
			"integrity": "sha512-+uw1inIHVPQoaVuHzRyXd21icM+cnt4CzD5rW+NC1wjOUSTOs+Te7FOv7AhN7vS9x/oIyhLP5PR1H+phQAHu5Q==",
			"dev": true,
			"requires": {
			  "once": "^1.4.0"
			}
		  },
		  "enhanced-resolve": {
			"version": "4.1.1",
			"resolved": "https://registry.npmjs.org/enhanced-resolve/-/enhanced-resolve-4.1.1.tgz",
			"integrity": "sha512-98p2zE+rL7/g/DzMHMTF4zZlCgeVdJ7yr6xzEpJRYwFYrGi9ANdn5DnJURg6RpBkyk60XYDnWIv51VfIhfNGuA==",
			"dev": true,
			"requires": {
			  "graceful-fs": "^4.1.2",
			  "memory-fs": "^0.5.0",
			  "tapable": "^1.0.0"
			}
		  },
		  "errno": {
			"version": "0.1.7",
			"resolved": "https://registry.npmjs.org/errno/-/errno-0.1.7.tgz",
			"integrity": "sha512-MfrRBDWzIWifgq6tJj60gkAwtLNb6sQPlcFrSOflcP1aFmmruKQ2wRnze/8V6kgyz7H3FF8Npzv78mZ7XLLflg==",
			"dev": true,
			"requires": {
			  "prr": "~1.0.1"
			}
		  },
		  "error-ex": {
			"version": "1.3.2",
			"resolved": "https://registry.npmjs.org/error-ex/-/error-ex-1.3.2.tgz",
			"integrity": "sha512-7dFHNmqeFSEt2ZBsCriorKnn3Z2pj+fd9kmI6QoWw4//DL+icEBfc0U7qJCisqrTsKTjw4fNFy2pW9OqStD84g==",
			"dev": true,
			"requires": {
			  "is-arrayish": "^0.2.1"
			}
		  },
		  "es-abstract": {
			"version": "1.17.5",
			"resolved": "https://registry.npmjs.org/es-abstract/-/es-abstract-1.17.5.tgz",
			"integrity": "sha512-BR9auzDbySxOcfog0tLECW8l28eRGpDpU3Dm3Hp4q/N+VtLTmyj4EUN088XZWQDW/hzj6sYRDXeOFsaAODKvpg==",
			"dev": true,
			"requires": {
			  "es-to-primitive": "^1.2.1",
			  "function-bind": "^1.1.1",
			  "has": "^1.0.3",
			  "has-symbols": "^1.0.1",
			  "is-callable": "^1.1.5",
			  "is-regex": "^1.0.5",
			  "object-inspect": "^1.7.0",
			  "object-keys": "^1.1.1",
			  "object.assign": "^4.1.0",
			  "string.prototype.trimleft": "^2.1.1",
			  "string.prototype.trimright": "^2.1.1"
			}
		  },
		  "es-to-primitive": {
			"version": "1.2.1",
			"resolved": "https://registry.npmjs.org/es-to-primitive/-/es-to-primitive-1.2.1.tgz",
			"integrity": "sha512-QCOllgZJtaUo9miYBcLChTUaHNjJF3PYs1VidD7AwiEj1kYxKeQTctLAezAOH5ZKRH0g2IgPn6KwB4IT8iRpvA==",
			"dev": true,
			"requires": {
			  "is-callable": "^1.1.4",
			  "is-date-object": "^1.0.1",
			  "is-symbol": "^1.0.2"
			}
		  },
		  "es5-ext": {
			"version": "0.10.53",
			"resolved": "https://registry.npmjs.org/es5-ext/-/es5-ext-0.10.53.tgz",
			"integrity": "sha512-Xs2Stw6NiNHWypzRTY1MtaG/uJlwCk8kH81920ma8mvN8Xq1gsfhZvpkImLQArw8AHnv8MT2I45J3c0R8slE+Q==",
			"requires": {
			  "es6-iterator": "~2.0.3",
			  "es6-symbol": "~3.1.3",
			  "next-tick": "~1.0.0"
			}
		  },
		  "es6-iterator": {
			"version": "2.0.3",
			"resolved": "https://registry.npmjs.org/es6-iterator/-/es6-iterator-2.0.3.tgz",
			"integrity": "sha1-p96IkUGgWpSwhUQDstCg+/qY87c=",
			"requires": {
			  "d": "1",
			  "es5-ext": "^0.10.35",
			  "es6-symbol": "^3.1.1"
			}
		  },
		  "es6-symbol": {
			"version": "3.1.3",
			"resolved": "https://registry.npmjs.org/es6-symbol/-/es6-symbol-3.1.3.tgz",
			"integrity": "sha512-NJ6Yn3FuDinBaBRWl/q5X/s4koRHBrgKAu+yGI6JCBeiu3qrcbJhwT2GeR/EXVfylRk8dpQVJoLEFhK+Mu31NA==",
			"requires": {
			  "d": "^1.0.1",
			  "ext": "^1.1.2"
			}
		  },
		  "es6-weak-map": {
			"version": "2.0.3",
			"resolved": "https://registry.npmjs.org/es6-weak-map/-/es6-weak-map-2.0.3.tgz",
			"integrity": "sha512-p5um32HOTO1kP+w7PRnB+5lQ43Z6muuMuIMffvDN8ZB4GcnjLBV6zGStpbASIMk4DCAvEaamhe2zhyCb/QXXsA==",
			"requires": {
			  "d": "1",
			  "es5-ext": "^0.10.46",
			  "es6-iterator": "^2.0.3",
			  "es6-symbol": "^3.1.1"
			}
		  },
		  "escape-html": {
			"version": "1.0.3",
			"resolved": "https://registry.npmjs.org/escape-html/-/escape-html-1.0.3.tgz",
			"integrity": "sha1-Aljq5NPQwJdN4cFpGI7wBR0dGYg="
		  },
		  "escape-string-regexp": {
			"version": "1.0.5",
			"resolved": "https://registry.npmjs.org/escape-string-regexp/-/escape-string-regexp-1.0.5.tgz",
			"integrity": "sha1-G2HAViGQqN/2rjuyzwIAyhMLhtQ="
		  },
		  "escodegen": {
			"version": "1.14.1",
			"resolved": "https://registry.npmjs.org/escodegen/-/escodegen-1.14.1.tgz",
			"integrity": "sha512-Bmt7NcRySdIfNPfU2ZoXDrrXsG9ZjvDxcAlMfDUgRBjLOWTuIACXPBFJH7Z+cLb40JeQco5toikyc9t9P8E9SQ==",
			"dev": true,
			"requires": {
			  "esprima": "^4.0.1",
			  "estraverse": "^4.2.0",
			  "esutils": "^2.0.2",
			  "optionator": "^0.8.1",
			  "source-map": "~0.6.1"
			},
			"dependencies": {
			  "levn": {
				"version": "0.3.0",
				"resolved": "https://registry.npmjs.org/levn/-/levn-0.3.0.tgz",
				"integrity": "sha1-OwmSTt+fCDwEkP3UwLxEIeBHZO4=",
				"dev": true,
				"requires": {
				  "prelude-ls": "~1.1.2",
				  "type-check": "~0.3.2"
				}
			  },
			  "optionator": {
				"version": "0.8.3",
				"resolved": "https://registry.npmjs.org/optionator/-/optionator-0.8.3.tgz",
				"integrity": "sha512-+IW9pACdk3XWmmTXG8m3upGUJst5XRGzxMRjXzAuJ1XnIFNvfhjjIuYkDvysnPQ7qzqVzLt78BCruntqRhWQbA==",
				"dev": true,
				"requires": {
				  "deep-is": "~0.1.3",
				  "fast-levenshtein": "~2.0.6",
				  "levn": "~0.3.0",
				  "prelude-ls": "~1.1.2",
				  "type-check": "~0.3.2",
				  "word-wrap": "~1.2.3"
				}
			  },
			  "prelude-ls": {
				"version": "1.1.2",
				"resolved": "https://registry.npmjs.org/prelude-ls/-/prelude-ls-1.1.2.tgz",
				"integrity": "sha1-IZMqVJ9eUv/ZqCf1cOBL5iqX2lQ=",
				"dev": true
			  },
			  "source-map": {
				"version": "0.6.1",
				"resolved": "https://registry.npmjs.org/source-map/-/source-map-0.6.1.tgz",
				"integrity": "sha512-UjgapumWlbMhkBgzT7Ykc5YXUT46F0iKu8SGXq0bcwP5dz/h0Plj6enJqjz1Zbq2l5WaqYnrVbwWOWMyF3F47g==",
				"dev": true,
				"optional": true
			  },
			  "type-check": {
				"version": "0.3.2",
				"resolved": "https://registry.npmjs.org/type-check/-/type-check-0.3.2.tgz",
				"integrity": "sha1-WITKtRLPHTVeP7eE8wgEsrUg23I=",
				"dev": true,
				"requires": {
				  "prelude-ls": "~1.1.2"
				}
			  }
			}
		  },
		  "eslint": {
			"version": "7.1.0",
			"resolved": "https://registry.npmjs.org/eslint/-/eslint-7.1.0.tgz",
			"integrity": "sha512-DfS3b8iHMK5z/YLSme8K5cge168I8j8o1uiVmFCgnnjxZQbCGyraF8bMl7Ju4yfBmCuxD7shOF7eqGkcuIHfsA==",
			"dev": true,
			"requires": {
			  "@babel/code-frame": "^7.0.0",
			  "ajv": "^6.10.0",
			  "chalk": "^4.0.0",
			  "cross-spawn": "^7.0.2",
			  "debug": "^4.0.1",
			  "doctrine": "^3.0.0",
			  "eslint-scope": "^5.0.0",
			  "eslint-utils": "^2.0.0",
			  "eslint-visitor-keys": "^1.1.0",
			  "espree": "^7.0.0",
			  "esquery": "^1.2.0",
			  "esutils": "^2.0.2",
			  "file-entry-cache": "^5.0.1",
			  "functional-red-black-tree": "^1.0.1",
			  "glob-parent": "^5.0.0",
			  "globals": "^12.1.0",
			  "ignore": "^4.0.6",
			  "import-fresh": "^3.0.0",
			  "imurmurhash": "^0.1.4",
			  "inquirer": "^7.0.0",
			  "is-glob": "^4.0.0",
			  "js-yaml": "^3.13.1",
			  "json-stable-stringify-without-jsonify": "^1.0.1",
			  "levn": "^0.4.1",
			  "lodash": "^4.17.14",
			  "minimatch": "^3.0.4",
			  "natural-compare": "^1.4.0",
			  "optionator": "^0.9.1",
			  "progress": "^2.0.0",
			  "regexpp": "^3.1.0",
			  "semver": "^7.2.1",
			  "strip-ansi": "^6.0.0",
			  "strip-json-comments": "^3.1.0",
			  "table": "^5.2.3",
			  "text-table": "^0.2.0",
			  "v8-compile-cache": "^2.0.3"
			},
			"dependencies": {
			  "ansi-regex": {
				"version": "5.0.0",
				"resolved": "https://registry.npmjs.org/ansi-regex/-/ansi-regex-5.0.0.tgz",
				"integrity": "sha512-bY6fj56OUQ0hU1KjFNDQuJFezqKdrAyFdIevADiqrWHwSlbmBNMHp5ak2f40Pm8JTFyM2mqxkG6ngkHO11f/lg==",
				"dev": true
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-4.0.0.tgz",
				"integrity": "sha512-N9oWFcegS0sFr9oh1oz2d7Npos6vNoWW9HvtCg5N1KRFpUhaAhvTv5Y58g880fZaEYSNm3qDz8SU1UrGvp+n7A==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "cross-spawn": {
				"version": "7.0.3",
				"resolved": "https://registry.npmjs.org/cross-spawn/-/cross-spawn-7.0.3.tgz",
				"integrity": "sha512-iRDPJKUPVEND7dHPO8rkbOnPpyDygcDFtWjpeWNCgy8WP2rXcxXL8TskReQl6OrB2G7+UJrags1q15Fudc7G6w==",
				"dev": true,
				"requires": {
				  "path-key": "^3.1.0",
				  "shebang-command": "^2.0.0",
				  "which": "^2.0.1"
				}
			  },
			  "debug": {
				"version": "4.1.1",
				"resolved": "https://registry.npmjs.org/debug/-/debug-4.1.1.tgz",
				"integrity": "sha512-pYAIzeRo8J6KPEaJ0VWOh5Pzkbw/RetuzehGM7QRRX5he4fPHx2rdKMB256ehJCkX+XRQm16eZLqLNS8RSZXZw==",
				"dev": true,
				"requires": {
				  "ms": "^2.1.1"
				}
			  },
			  "eslint-scope": {
				"version": "5.0.0",
				"resolved": "https://registry.npmjs.org/eslint-scope/-/eslint-scope-5.0.0.tgz",
				"integrity": "sha512-oYrhJW7S0bxAFDvWqzvMPRm6pcgcnWc4QnofCAqRTRfQC0JcwenzGglTtsLyIuuWFfkqDG9vz67cnttSd53djw==",
				"dev": true,
				"requires": {
				  "esrecurse": "^4.1.0",
				  "estraverse": "^4.1.1"
				}
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "ms": {
				"version": "2.1.2",
				"resolved": "https://registry.npmjs.org/ms/-/ms-2.1.2.tgz",
				"integrity": "sha512-sGkPx+VjMtmA6MX27oA4FBFELFCZZ4S4XqeGOXCv68tT+jb3vk/RyaKWP0PTKyWtmLSM0b+adUTEvbs1PEaH2w==",
				"dev": true
			  },
			  "path-key": {
				"version": "3.1.1",
				"resolved": "https://registry.npmjs.org/path-key/-/path-key-3.1.1.tgz",
				"integrity": "sha512-ojmeN0qd+y0jszEtoY48r0Peq5dwMEkIlCOu6Q5f41lfkswXuKtYrhgoTpLnyIcHm24Uhqx+5Tqm2InSwLhE6Q==",
				"dev": true
			  },
			  "semver": {
				"version": "7.3.2",
				"resolved": "https://registry.npmjs.org/semver/-/semver-7.3.2.tgz",
				"integrity": "sha512-OrOb32TeeambH6UrhtShmF7CRDqhL6/5XpPNp2DuRH6+9QLw/orhp72j87v8Qa1ScDkvrrBNpZcDejAirJmfXQ==",
				"dev": true
			  },
			  "shebang-command": {
				"version": "2.0.0",
				"resolved": "https://registry.npmjs.org/shebang-command/-/shebang-command-2.0.0.tgz",
				"integrity": "sha512-kHxr2zZpYtdmrN1qDjrrX/Z1rR1kG8Dx+gkpK1G4eXmvXswmcE1hTWBWYUzlraYw1/yZp6YuDY77YtvbN0dmDA==",
				"dev": true,
				"requires": {
				  "shebang-regex": "^3.0.0"
				}
			  },
			  "shebang-regex": {
				"version": "3.0.0",
				"resolved": "https://registry.npmjs.org/shebang-regex/-/shebang-regex-3.0.0.tgz",
				"integrity": "sha512-7++dFhtcx3353uBaq8DDR4NuxBetBzC7ZQOhmTQInHEd6bSrXdiEyzCvG07Z44UYdLShWUyXt5M/yhz8ekcb1A==",
				"dev": true
			  },
			  "strip-ansi": {
				"version": "6.0.0",
				"resolved": "https://registry.npmjs.org/strip-ansi/-/strip-ansi-6.0.0.tgz",
				"integrity": "sha512-AuvKTrTfQNYNIctbR1K/YGTR1756GycPsg7b9bdV9Duqur4gv6aKqHXah67Z8ImS7WEz5QVcOtlfW2rZEugt6w==",
				"dev": true,
				"requires": {
				  "ansi-regex": "^5.0.0"
				}
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  },
			  "which": {
				"version": "2.0.2",
				"resolved": "https://registry.npmjs.org/which/-/which-2.0.2.tgz",
				"integrity": "sha512-BLI3Tl1TW3Pvl70l3yq3Y64i+awpwXqsGBYWkkqMtnbXgrMD+yj7rhW0kuEDxzJaYXGjEW5ogapKNMEKNMjibA==",
				"dev": true,
				"requires": {
				  "isexe": "^2.0.0"
				}
			  }
			}
		  },
		  "eslint-config-prettier": {
			"version": "6.11.0",
			"resolved": "https://registry.npmjs.org/eslint-config-prettier/-/eslint-config-prettier-6.11.0.tgz",
			"integrity": "sha512-oB8cpLWSAjOVFEJhhyMZh6NOEOtBVziaqdDQ86+qhDHFbZXoRTM7pNSvFRfW/W/L/LrQ38C99J5CGuRBBzBsdA==",
			"dev": true,
			"requires": {
			  "get-stdin": "^6.0.0"
			}
		  },
		  "eslint-import-resolver-node": {
			"version": "0.3.3",
			"resolved": "https://registry.npmjs.org/eslint-import-resolver-node/-/eslint-import-resolver-node-0.3.3.tgz",
			"integrity": "sha512-b8crLDo0M5RSe5YG8Pu2DYBj71tSB6OvXkfzwbJU2w7y8P4/yo0MyF8jU26IEuEuHF2K5/gcAJE3LhQGqBBbVg==",
			"dev": true,
			"requires": {
			  "debug": "^2.6.9",
			  "resolve": "^1.13.1"
			},
			"dependencies": {
			  "debug": {
				"version": "2.6.9",
				"resolved": "https://registry.npmjs.org/debug/-/debug-2.6.9.tgz",
				"integrity": "sha512-bC7ElrdJaJnPbAP+1EotYvqZsb3ecl5wi6Bfi6BJTUcNowp6cvspg0jXznRTKDjm/E7AdgFBVeAPVMNcKGsHMA==",
				"dev": true,
				"requires": {
				  "ms": "2.0.0"
				}
			  }
			}
		  },
		  "eslint-module-utils": {
			"version": "2.6.0",
			"resolved": "https://registry.npmjs.org/eslint-module-utils/-/eslint-module-utils-2.6.0.tgz",
			"integrity": "sha512-6j9xxegbqe8/kZY8cYpcp0xhbK0EgJlg3g9mib3/miLaExuuwc3n5UEfSnU6hWMbT0FAYVvDbL9RrRgpUeQIvA==",
			"dev": true,
			"requires": {
			  "debug": "^2.6.9",
			  "pkg-dir": "^2.0.0"
			},
			"dependencies": {
			  "debug": {
				"version": "2.6.9",
				"resolved": "https://registry.npmjs.org/debug/-/debug-2.6.9.tgz",
				"integrity": "sha512-bC7ElrdJaJnPbAP+1EotYvqZsb3ecl5wi6Bfi6BJTUcNowp6cvspg0jXznRTKDjm/E7AdgFBVeAPVMNcKGsHMA==",
				"dev": true,
				"requires": {
				  "ms": "2.0.0"
				}
			  },
			  "find-up": {
				"version": "2.1.0",
				"resolved": "https://registry.npmjs.org/find-up/-/find-up-2.1.0.tgz",
				"integrity": "sha1-RdG35QbHF93UgndaK3eSCjwMV6c=",
				"dev": true,
				"requires": {
				  "locate-path": "^2.0.0"
				}
			  },
			  "locate-path": {
				"version": "2.0.0",
				"resolved": "https://registry.npmjs.org/locate-path/-/locate-path-2.0.0.tgz",
				"integrity": "sha1-K1aLJl7slExtnA3pw9u7ygNUzY4=",
				"dev": true,
				"requires": {
				  "p-locate": "^2.0.0",
				  "path-exists": "^3.0.0"
				}
			  },
			  "p-limit": {
				"version": "1.3.0",
				"resolved": "https://registry.npmjs.org/p-limit/-/p-limit-1.3.0.tgz",
				"integrity": "sha512-vvcXsLAJ9Dr5rQOPk7toZQZJApBl2K4J6dANSsEuh6QI41JYcsS/qhTGa9ErIUUgK3WNQoJYvylxvjqmiqEA9Q==",
				"dev": true,
				"requires": {
				  "p-try": "^1.0.0"
				}
			  },
			  "p-locate": {
				"version": "2.0.0",
				"resolved": "https://registry.npmjs.org/p-locate/-/p-locate-2.0.0.tgz",
				"integrity": "sha1-IKAQOyIqcMj9OcwuWAaA893l7EM=",
				"dev": true,
				"requires": {
				  "p-limit": "^1.1.0"
				}
			  },
			  "p-try": {
				"version": "1.0.0",
				"resolved": "https://registry.npmjs.org/p-try/-/p-try-1.0.0.tgz",
				"integrity": "sha1-y8ec26+P1CKOE/Yh8rGiN8GyB7M=",
				"dev": true
			  },
			  "pkg-dir": {
				"version": "2.0.0",
				"resolved": "https://registry.npmjs.org/pkg-dir/-/pkg-dir-2.0.0.tgz",
				"integrity": "sha1-9tXREJ4Z1j7fQo4L1X4Sd3YVM0s=",
				"dev": true,
				"requires": {
				  "find-up": "^2.1.0"
				}
			  }
			}
		  },
		  "eslint-plugin-import": {
			"version": "2.20.2",
			"resolved": "https://registry.npmjs.org/eslint-plugin-import/-/eslint-plugin-import-2.20.2.tgz",
			"integrity": "sha512-FObidqpXrR8OnCh4iNsxy+WACztJLXAHBO5hK79T1Hc77PgQZkyDGA5Ag9xAvRpglvLNxhH/zSmZ70/pZ31dHg==",
			"dev": true,
			"requires": {
			  "array-includes": "^3.0.3",
			  "array.prototype.flat": "^1.2.1",
			  "contains-path": "^0.1.0",
			  "debug": "^2.6.9",
			  "doctrine": "1.5.0",
			  "eslint-import-resolver-node": "^0.3.2",
			  "eslint-module-utils": "^2.4.1",
			  "has": "^1.0.3",
			  "minimatch": "^3.0.4",
			  "object.values": "^1.1.0",
			  "read-pkg-up": "^2.0.0",
			  "resolve": "^1.12.0"
			},
			"dependencies": {
			  "debug": {
				"version": "2.6.9",
				"resolved": "https://registry.npmjs.org/debug/-/debug-2.6.9.tgz",
				"integrity": "sha512-bC7ElrdJaJnPbAP+1EotYvqZsb3ecl5wi6Bfi6BJTUcNowp6cvspg0jXznRTKDjm/E7AdgFBVeAPVMNcKGsHMA==",
				"dev": true,
				"requires": {
				  "ms": "2.0.0"
				}
			  },
			  "doctrine": {
				"version": "1.5.0",
				"resolved": "https://registry.npmjs.org/doctrine/-/doctrine-1.5.0.tgz",
				"integrity": "sha1-N53Ocw9hZvds76TmcHoVmwLFpvo=",
				"dev": true,
				"requires": {
				  "esutils": "^2.0.2",
				  "isarray": "^1.0.0"
				}
			  },
			  "isarray": {
				"version": "1.0.0",
				"resolved": "https://registry.npmjs.org/isarray/-/isarray-1.0.0.tgz",
				"integrity": "sha1-u5NdSFgsuhaMBoNJV6VKPgcSTxE=",
				"dev": true
			  }
			}
		  },
		  "eslint-scope": {
			"version": "4.0.3",
			"resolved": "https://registry.npmjs.org/eslint-scope/-/eslint-scope-4.0.3.tgz",
			"integrity": "sha512-p7VutNr1O/QrxysMo3E45FjYDTeXBy0iTltPFNSqKAIfjDSXC+4dj+qfyuD8bfAXrW/y6lW3O76VaYNPKfpKrg==",
			"dev": true,
			"requires": {
			  "esrecurse": "^4.1.0",
			  "estraverse": "^4.1.1"
			}
		  },
		  "eslint-utils": {
			"version": "2.0.0",
			"resolved": "https://registry.npmjs.org/eslint-utils/-/eslint-utils-2.0.0.tgz",
			"integrity": "sha512-0HCPuJv+7Wv1bACm8y5/ECVfYdfsAm9xmVb7saeFlxjPYALefjhbYoCkBjPdPzGH8wWyTpAez82Fh3VKYEZ8OA==",
			"dev": true,
			"requires": {
			  "eslint-visitor-keys": "^1.1.0"
			}
		  },
		  "eslint-visitor-keys": {
			"version": "1.1.0",
			"resolved": "https://registry.npmjs.org/eslint-visitor-keys/-/eslint-visitor-keys-1.1.0.tgz",
			"integrity": "sha512-8y9YjtM1JBJU/A9Kc+SbaOV4y29sSWckBwMHa+FGtVj5gN/sbnKDf6xJUl+8g7FAij9LVaP8C24DUiH/f/2Z9A==",
			"dev": true
		  },
		  "espree": {
			"version": "7.0.0",
			"resolved": "https://registry.npmjs.org/espree/-/espree-7.0.0.tgz",
			"integrity": "sha512-/r2XEx5Mw4pgKdyb7GNLQNsu++asx/dltf/CI8RFi9oGHxmQFgvLbc5Op4U6i8Oaj+kdslhJtVlEZeAqH5qOTw==",
			"dev": true,
			"requires": {
			  "acorn": "^7.1.1",
			  "acorn-jsx": "^5.2.0",
			  "eslint-visitor-keys": "^1.1.0"
			},
			"dependencies": {
			  "acorn": {
				"version": "7.2.0",
				"resolved": "https://registry.npmjs.org/acorn/-/acorn-7.2.0.tgz",
				"integrity": "sha512-apwXVmYVpQ34m/i71vrApRrRKCWQnZZF1+npOD0WV5xZFfwWOmKGQ2RWlfdy9vWITsenisM8M0Qeq8agcFHNiQ==",
				"dev": true
			  }
			}
		  },
		  "esprima": {
			"version": "4.0.1",
			"resolved": "https://registry.npmjs.org/esprima/-/esprima-4.0.1.tgz",
			"integrity": "sha512-eGuFFw7Upda+g4p+QHvnW0RyTX/SVeJBDM/gCtMARO0cLuT2HcEKnTPvhjV6aGeqrCB/sbNop0Kszm0jsaWU4A==",
			"dev": true
		  },
		  "esquery": {
			"version": "1.3.1",
			"resolved": "https://registry.npmjs.org/esquery/-/esquery-1.3.1.tgz",
			"integrity": "sha512-olpvt9QG0vniUBZspVRN6lwB7hOZoTRtT+jzR+tS4ffYx2mzbw+z0XCOk44aaLYKApNX5nMm+E+P6o25ip/DHQ==",
			"dev": true,
			"requires": {
			  "estraverse": "^5.1.0"
			},
			"dependencies": {
			  "estraverse": {
				"version": "5.1.0",
				"resolved": "https://registry.npmjs.org/estraverse/-/estraverse-5.1.0.tgz",
				"integrity": "sha512-FyohXK+R0vE+y1nHLoBM7ZTyqRpqAlhdZHCWIWEviFLiGB8b04H6bQs8G+XTthacvT8VuwvteiP7RJSxMs8UEw==",
				"dev": true
			  }
			}
		  },
		  "esrecurse": {
			"version": "4.2.1",
			"resolved": "https://registry.npmjs.org/esrecurse/-/esrecurse-4.2.1.tgz",
			"integrity": "sha512-64RBB++fIOAXPw3P9cy89qfMlvZEXZkqqJkjqqXIvzP5ezRZjW+lPWjw35UX/3EhUPFYbg5ER4JYgDw4007/DQ==",
			"dev": true,
			"requires": {
			  "estraverse": "^4.1.0"
			}
		  },
		  "estraverse": {
			"version": "4.3.0",
			"resolved": "https://registry.npmjs.org/estraverse/-/estraverse-4.3.0.tgz",
			"integrity": "sha512-39nnKffWz8xN1BU/2c79n9nB9HDzo0niYUqx6xyqUnyoAnQyyWpOTdZEeiCch8BBu515t4wp9ZmgVfVhn9EBpw==",
			"dev": true
		  },
		  "esutils": {
			"version": "2.0.3",
			"resolved": "https://registry.npmjs.org/esutils/-/esutils-2.0.3.tgz",
			"integrity": "sha512-kVscqXk4OCp68SZ0dkgEKVi6/8ij300KBWTJq32P/dYeWTSwK41WyTxalN1eRmA5Z9UU/LX9D7FWSmV9SAYx6g==",
			"dev": true
		  },
		  "etag": {
			"version": "1.8.1",
			"resolved": "https://registry.npmjs.org/etag/-/etag-1.8.1.tgz",
			"integrity": "sha1-Qa4u62XvpiJorr/qg6x9eSmbCIc="
		  },
		  "event-emitter": {
			"version": "0.3.5",
			"resolved": "https://registry.npmjs.org/event-emitter/-/event-emitter-0.3.5.tgz",
			"integrity": "sha1-34xp7vFkeSPHFXuc6DhAYQsCzDk=",
			"requires": {
			  "d": "1",
			  "es5-ext": "~0.10.14"
			}
		  },
		  "events": {
			"version": "3.1.0",
			"resolved": "https://registry.npmjs.org/events/-/events-3.1.0.tgz",
			"integrity": "sha512-Rv+u8MLHNOdMjTAFeT3nCjHn2aGlx435FP/sDHNaRhDEMwyI/aB22Kj2qIN8R0cw3z28psEQLYwxVKLsKrMgWg==",
			"dev": true
		  },
		  "evp_bytestokey": {
			"version": "1.0.3",
			"resolved": "https://registry.npmjs.org/evp_bytestokey/-/evp_bytestokey-1.0.3.tgz",
			"integrity": "sha512-/f2Go4TognH/KvCISP7OUsHn85hT9nUkxxA9BEWxFn+Oj9o8ZNLm/40hdlgSLyuOimsrTKLUMEorQexp/aPQeA==",
			"dev": true,
			"requires": {
			  "md5.js": "^1.3.4",
			  "safe-buffer": "^5.1.1"
			}
		  },
		  "exec-sh": {
			"version": "0.3.4",
			"resolved": "https://registry.npmjs.org/exec-sh/-/exec-sh-0.3.4.tgz",
			"integrity": "sha512-sEFIkc61v75sWeOe72qyrqg2Qg0OuLESziUDk/O/z2qgS15y2gWVFrI6f2Qn/qw/0/NCfCEsmNA4zOjkwEZT1A==",
			"dev": true
		  },
		  "execa": {
			"version": "1.0.0",
			"resolved": "https://registry.npmjs.org/execa/-/execa-1.0.0.tgz",
			"integrity": "sha512-adbxcyWV46qiHyvSp50TKt05tB4tK3HcmF7/nxfAdhnox83seTDbwnaqKO4sXRy7roHAIFqJP/Rw/AuEbX61LA==",
			"dev": true,
			"requires": {
			  "cross-spawn": "^6.0.0",
			  "get-stream": "^4.0.0",
			  "is-stream": "^1.1.0",
			  "npm-run-path": "^2.0.0",
			  "p-finally": "^1.0.0",
			  "signal-exit": "^3.0.0",
			  "strip-eof": "^1.0.0"
			}
		  },
		  "exit": {
			"version": "0.1.2",
			"resolved": "https://registry.npmjs.org/exit/-/exit-0.1.2.tgz",
			"integrity": "sha1-BjJjj42HfMghB9MKD/8aF8uhzQw=",
			"dev": true
		  },
		  "expand-brackets": {
			"version": "2.1.4",
			"resolved": "https://registry.npmjs.org/expand-brackets/-/expand-brackets-2.1.4.tgz",
			"integrity": "sha1-t3c14xXOMPa27/D4OwQVGiJEliI=",
			"dev": true,
			"requires": {
			  "debug": "^2.3.3",
			  "define-property": "^0.2.5",
			  "extend-shallow": "^2.0.1",
			  "posix-character-classes": "^0.1.0",
			  "regex-not": "^1.0.0",
			  "snapdragon": "^0.8.1",
			  "to-regex": "^3.0.1"
			},
			"dependencies": {
			  "debug": {
				"version": "2.6.9",
				"resolved": "https://registry.npmjs.org/debug/-/debug-2.6.9.tgz",
				"integrity": "sha512-bC7ElrdJaJnPbAP+1EotYvqZsb3ecl5wi6Bfi6BJTUcNowp6cvspg0jXznRTKDjm/E7AdgFBVeAPVMNcKGsHMA==",
				"dev": true,
				"requires": {
				  "ms": "2.0.0"
				}
			  },
			  "define-property": {
				"version": "0.2.5",
				"resolved": "https://registry.npmjs.org/define-property/-/define-property-0.2.5.tgz",
				"integrity": "sha1-w1se+RjsPJkPmlvFe+BKrOxcgRY=",
				"dev": true,
				"requires": {
				  "is-descriptor": "^0.1.0"
				}
			  },
			  "extend-shallow": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/extend-shallow/-/extend-shallow-2.0.1.tgz",
				"integrity": "sha1-Ua99YUrZqfYQ6huvu5idaxxWiQ8=",
				"dev": true,
				"requires": {
				  "is-extendable": "^0.1.0"
				}
			  }
			}
		  },
		  "expect": {
			"version": "26.0.1",
			"resolved": "https://registry.npmjs.org/expect/-/expect-26.0.1.tgz",
			"integrity": "sha512-QcCy4nygHeqmbw564YxNbHTJlXh47dVID2BUP52cZFpLU9zHViMFK6h07cC1wf7GYCTIigTdAXhVua8Yl1FkKg==",
			"dev": true,
			"requires": {
			  "@jest/types": "^26.0.1",
			  "ansi-styles": "^4.0.0",
			  "jest-get-type": "^26.0.0",
			  "jest-matcher-utils": "^26.0.1",
			  "jest-message-util": "^26.0.1",
			  "jest-regex-util": "^26.0.0"
			},
			"dependencies": {
			  "@jest/types": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/@jest/types/-/types-26.0.1.tgz",
				"integrity": "sha512-IbtjvqI9+eS1qFnOIEL7ggWmT+iK/U+Vde9cGWtYb/b6XgKb3X44ZAe/z9YZzoAAZ/E92m0DqrilF934IGNnQA==",
				"dev": true,
				"requires": {
				  "@types/istanbul-lib-coverage": "^2.0.0",
				  "@types/istanbul-reports": "^1.1.1",
				  "@types/yargs": "^15.0.0",
				  "chalk": "^4.0.0"
				}
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-4.0.0.tgz",
				"integrity": "sha512-N9oWFcegS0sFr9oh1oz2d7Npos6vNoWW9HvtCg5N1KRFpUhaAhvTv5Y58g880fZaEYSNm3qDz8SU1UrGvp+n7A==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "jest-get-type": {
				"version": "26.0.0",
				"resolved": "https://registry.npmjs.org/jest-get-type/-/jest-get-type-26.0.0.tgz",
				"integrity": "sha512-zRc1OAPnnws1EVfykXOj19zo2EMw5Hi6HLbFCSjpuJiXtOWAYIjNsHVSbpQ8bDX7L5BGYGI8m+HmKdjHYFF0kg==",
				"dev": true
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "express": {
			"version": "4.17.1",
			"resolved": "https://registry.npmjs.org/express/-/express-4.17.1.tgz",
			"integrity": "sha512-mHJ9O79RqluphRrcw2X/GTh3k9tVv8YcoyY4Kkh4WDMUYKRZUq0h1o0w2rrrxBqM7VoeUVqgb27xlEMXTnYt4g==",
			"requires": {
			  "accepts": "~1.3.7",
			  "array-flatten": "1.1.1",
			  "body-parser": "1.19.0",
			  "content-disposition": "0.5.3",
			  "content-type": "~1.0.4",
			  "cookie": "0.4.0",
			  "cookie-signature": "1.0.6",
			  "debug": "2.6.9",
			  "depd": "~1.1.2",
			  "encodeurl": "~1.0.2",
			  "escape-html": "~1.0.3",
			  "etag": "~1.8.1",
			  "finalhandler": "~1.1.2",
			  "fresh": "0.5.2",
			  "merge-descriptors": "1.0.1",
			  "methods": "~1.1.2",
			  "on-finished": "~2.3.0",
			  "parseurl": "~1.3.3",
			  "path-to-regexp": "0.1.7",
			  "proxy-addr": "~2.0.5",
			  "qs": "6.7.0",
			  "range-parser": "~1.2.1",
			  "safe-buffer": "5.1.2",
			  "send": "0.17.1",
			  "serve-static": "1.14.1",
			  "setprototypeof": "1.1.1",
			  "statuses": "~1.5.0",
			  "type-is": "~1.6.18",
			  "utils-merge": "1.0.1",
			  "vary": "~1.1.2"
			},
			"dependencies": {
			  "debug": {
				"version": "2.6.9",
				"resolved": "https://registry.npmjs.org/debug/-/debug-2.6.9.tgz",
				"integrity": "sha512-bC7ElrdJaJnPbAP+1EotYvqZsb3ecl5wi6Bfi6BJTUcNowp6cvspg0jXznRTKDjm/E7AdgFBVeAPVMNcKGsHMA==",
				"requires": {
				  "ms": "2.0.0"
				}
			  },
			  "path-to-regexp": {
				"version": "0.1.7",
				"resolved": "https://registry.npmjs.org/path-to-regexp/-/path-to-regexp-0.1.7.tgz",
				"integrity": "sha1-32BBeABfUi8V60SQ5yR6G/qmf4w="
			  }
			}
		  },
		  "ext": {
			"version": "1.4.0",
			"resolved": "https://registry.npmjs.org/ext/-/ext-1.4.0.tgz",
			"integrity": "sha512-Key5NIsUxdqKg3vIsdw9dSuXpPCQ297y6wBjL30edxwPgt2E44WcWBZey/ZvUc6sERLTxKdyCu4gZFmUbk1Q7A==",
			"requires": {
			  "type": "^2.0.0"
			},
			"dependencies": {
			  "type": {
				"version": "2.0.0",
				"resolved": "https://registry.npmjs.org/type/-/type-2.0.0.tgz",
				"integrity": "sha512-KBt58xCHry4Cejnc2ISQAF7QY+ORngsWfxezO68+12hKV6lQY8P/psIkcbjeHWn7MqcgciWJyCCevFMJdIXpow=="
			  }
			}
		  },
		  "extend": {
			"version": "3.0.2",
			"resolved": "https://registry.npmjs.org/extend/-/extend-3.0.2.tgz",
			"integrity": "sha512-fjquC59cD7CyW6urNXK0FBufkZcoiGG80wTuPujX590cB5Ttln20E2UB4S/WARVqhXffZl2LNgS+gQdPIIim/g==",
			"dev": true
		  },
		  "extend-shallow": {
			"version": "3.0.2",
			"resolved": "https://registry.npmjs.org/extend-shallow/-/extend-shallow-3.0.2.tgz",
			"integrity": "sha1-Jqcarwc7OfshJxcnRhMcJwQCjbg=",
			"dev": true,
			"requires": {
			  "assign-symbols": "^1.0.0",
			  "is-extendable": "^1.0.1"
			},
			"dependencies": {
			  "is-extendable": {
				"version": "1.0.1",
				"resolved": "https://registry.npmjs.org/is-extendable/-/is-extendable-1.0.1.tgz",
				"integrity": "sha512-arnXMxT1hhoKo9k1LZdmlNyJdDDfy2v0fXjFlmok4+i8ul/6WlbVge9bhM74OpNPQPMGUToDtz+KXa1PneJxOA==",
				"dev": true,
				"requires": {
				  "is-plain-object": "^2.0.4"
				}
			  }
			}
		  },
		  "external-editor": {
			"version": "3.1.0",
			"resolved": "https://registry.npmjs.org/external-editor/-/external-editor-3.1.0.tgz",
			"integrity": "sha512-hMQ4CX1p1izmuLYyZqLMO/qGNw10wSv9QDCPfzXfyFrOaCSSoRfqE1Kf1s5an66J5JZC62NewG+mK49jOCtQew==",
			"dev": true,
			"requires": {
			  "chardet": "^0.7.0",
			  "iconv-lite": "^0.4.24",
			  "tmp": "^0.0.33"
			}
		  },
		  "extglob": {
			"version": "2.0.4",
			"resolved": "https://registry.npmjs.org/extglob/-/extglob-2.0.4.tgz",
			"integrity": "sha512-Nmb6QXkELsuBr24CJSkilo6UHHgbekK5UiZgfE6UHD3Eb27YC6oD+bhcT+tJ6cl8dmsgdQxnWlcry8ksBIBLpw==",
			"dev": true,
			"requires": {
			  "array-unique": "^0.3.2",
			  "define-property": "^1.0.0",
			  "expand-brackets": "^2.1.4",
			  "extend-shallow": "^2.0.1",
			  "fragment-cache": "^0.2.1",
			  "regex-not": "^1.0.0",
			  "snapdragon": "^0.8.1",
			  "to-regex": "^3.0.1"
			},
			"dependencies": {
			  "define-property": {
				"version": "1.0.0",
				"resolved": "https://registry.npmjs.org/define-property/-/define-property-1.0.0.tgz",
				"integrity": "sha1-dp66rz9KY6rTr56NMEybvnm/sOY=",
				"dev": true,
				"requires": {
				  "is-descriptor": "^1.0.0"
				}
			  },
			  "extend-shallow": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/extend-shallow/-/extend-shallow-2.0.1.tgz",
				"integrity": "sha1-Ua99YUrZqfYQ6huvu5idaxxWiQ8=",
				"dev": true,
				"requires": {
				  "is-extendable": "^0.1.0"
				}
			  },
			  "is-accessor-descriptor": {
				"version": "1.0.0",
				"resolved": "https://registry.npmjs.org/is-accessor-descriptor/-/is-accessor-descriptor-1.0.0.tgz",
				"integrity": "sha512-m5hnHTkcVsPfqx3AKlyttIPb7J+XykHvJP2B9bZDjlhLIoEq4XoK64Vg7boZlVWYK6LUY94dYPEE7Lh0ZkZKcQ==",
				"dev": true,
				"requires": {
				  "kind-of": "^6.0.0"
				}
			  },
			  "is-data-descriptor": {
				"version": "1.0.0",
				"resolved": "https://registry.npmjs.org/is-data-descriptor/-/is-data-descriptor-1.0.0.tgz",
				"integrity": "sha512-jbRXy1FmtAoCjQkVmIVYwuuqDFUbaOeDjmed1tOGPrsMhtJA4rD9tkgA0F1qJ3gRFRXcHYVkdeaP50Q5rE/jLQ==",
				"dev": true,
				"requires": {
				  "kind-of": "^6.0.0"
				}
			  },
			  "is-descriptor": {
				"version": "1.0.2",
				"resolved": "https://registry.npmjs.org/is-descriptor/-/is-descriptor-1.0.2.tgz",
				"integrity": "sha512-2eis5WqQGV7peooDyLmNEPUrps9+SXX5c9pL3xEB+4e9HnGuDa7mB7kHxHw4CbqS9k1T2hOH3miL8n8WtiYVtg==",
				"dev": true,
				"requires": {
				  "is-accessor-descriptor": "^1.0.0",
				  "is-data-descriptor": "^1.0.0",
				  "kind-of": "^6.0.2"
				}
			  }
			}
		  },
		  "extsprintf": {
			"version": "1.3.0",
			"resolved": "https://registry.npmjs.org/extsprintf/-/extsprintf-1.3.0.tgz",
			"integrity": "sha1-lpGEQOMEGnpBT4xS48V06zw+HgU=",
			"dev": true
		  },
		  "fast-deep-equal": {
			"version": "2.0.1",
			"resolved": "https://registry.npmjs.org/fast-deep-equal/-/fast-deep-equal-2.0.1.tgz",
			"integrity": "sha1-ewUhjd+WZ79/Nwv3/bLLFf3Qqkk=",
			"dev": true
		  },
		  "fast-json-stable-stringify": {
			"version": "2.0.0",
			"resolved": "https://registry.npmjs.org/fast-json-stable-stringify/-/fast-json-stable-stringify-2.0.0.tgz",
			"integrity": "sha1-1RQsDK7msRifh9OnYREGT4bIu/I=",
			"dev": true
		  },
		  "fast-levenshtein": {
			"version": "2.0.6",
			"resolved": "https://registry.npmjs.org/fast-levenshtein/-/fast-levenshtein-2.0.6.tgz",
			"integrity": "sha1-PYpcZog6FqMMqGQ+hR8Zuqd5eRc=",
			"dev": true
		  },
		  "fast-safe-stringify": {
			"version": "2.0.7",
			"resolved": "https://registry.npmjs.org/fast-safe-stringify/-/fast-safe-stringify-2.0.7.tgz",
			"integrity": "sha512-Utm6CdzT+6xsDk2m8S6uL8VHxNwI6Jub+e9NYTcAms28T84pTa25GJQV9j0CY0N1rM8hK4x6grpF2BQf+2qwVA=="
		  },
		  "fb-watchman": {
			"version": "2.0.1",
			"resolved": "https://registry.npmjs.org/fb-watchman/-/fb-watchman-2.0.1.tgz",
			"integrity": "sha512-DkPJKQeY6kKwmuMretBhr7G6Vodr7bFwDYTXIkfG1gjvNpaxBTQV3PbXg6bR1c1UP4jPOX0jHUbbHANL9vRjVg==",
			"dev": true,
			"requires": {
			  "bser": "2.1.1"
			}
		  },
		  "figgy-pudding": {
			"version": "3.5.2",
			"resolved": "https://registry.npmjs.org/figgy-pudding/-/figgy-pudding-3.5.2.tgz",
			"integrity": "sha512-0btnI/H8f2pavGMN8w40mlSKOfTK2SVJmBfBeVIj3kNw0swwgzyRq0d5TJVOwodFmtvpPeWPN/MCcfuWF0Ezbw==",
			"dev": true
		  },
		  "figures": {
			"version": "3.2.0",
			"resolved": "https://registry.npmjs.org/figures/-/figures-3.2.0.tgz",
			"integrity": "sha512-yaduQFRKLXYOGgEn6AZau90j3ggSOyiqXU0F9JZfeXYhNa+Jk4X+s45A2zg5jns87GAFa34BBm2kXw4XpNcbdg==",
			"dev": true,
			"requires": {
			  "escape-string-regexp": "^1.0.5"
			}
		  },
		  "file-entry-cache": {
			"version": "5.0.1",
			"resolved": "https://registry.npmjs.org/file-entry-cache/-/file-entry-cache-5.0.1.tgz",
			"integrity": "sha512-bCg29ictuBaKUwwArK4ouCaqDgLZcysCFLmM/Yn/FDoqndh/9vNuQfXRDvTuXKLxfD/JtZQGKFT8MGcJBK644g==",
			"dev": true,
			"requires": {
			  "flat-cache": "^2.0.1"
			}
		  },
		  "file-uri-to-path": {
			"version": "1.0.0",
			"resolved": "https://registry.npmjs.org/file-uri-to-path/-/file-uri-to-path-1.0.0.tgz",
			"integrity": "sha512-0Zt+s3L7Vf1biwWZ29aARiVYLx7iMGnEUl9x33fbB/j3jR81u/O2LbqK+Bm1CDSNDKVtJ/YjwY7TUd5SkeLQLw==",
			"dev": true,
			"optional": true
		  },
		  "fill-range": {
			"version": "7.0.1",
			"resolved": "https://registry.npmjs.org/fill-range/-/fill-range-7.0.1.tgz",
			"integrity": "sha512-qOo9F+dMUmC2Lcb4BbVvnKJxTPjCm+RRpe4gDuGrzkL7mEVl/djYSu2OdQ2Pa302N4oqkSg9ir6jaLWJ2USVpQ==",
			"dev": true,
			"requires": {
			  "to-regex-range": "^5.0.1"
			}
		  },
		  "finalhandler": {
			"version": "1.1.2",
			"resolved": "https://registry.npmjs.org/finalhandler/-/finalhandler-1.1.2.tgz",
			"integrity": "sha512-aAWcW57uxVNrQZqFXjITpW3sIUQmHGG3qSb9mUah9MgMC4NeWhNOlNjXEYq3HjRAvL6arUviZGGJsBg6z0zsWA==",
			"requires": {
			  "debug": "2.6.9",
			  "encodeurl": "~1.0.2",
			  "escape-html": "~1.0.3",
			  "on-finished": "~2.3.0",
			  "parseurl": "~1.3.3",
			  "statuses": "~1.5.0",
			  "unpipe": "~1.0.0"
			},
			"dependencies": {
			  "debug": {
				"version": "2.6.9",
				"resolved": "https://registry.npmjs.org/debug/-/debug-2.6.9.tgz",
				"integrity": "sha512-bC7ElrdJaJnPbAP+1EotYvqZsb3ecl5wi6Bfi6BJTUcNowp6cvspg0jXznRTKDjm/E7AdgFBVeAPVMNcKGsHMA==",
				"requires": {
				  "ms": "2.0.0"
				}
			  }
			}
		  },
		  "find-cache-dir": {
			"version": "2.1.0",
			"resolved": "https://registry.npmjs.org/find-cache-dir/-/find-cache-dir-2.1.0.tgz",
			"integrity": "sha512-Tq6PixE0w/VMFfCgbONnkiQIVol/JJL7nRMi20fqzA4NRs9AfeqMGeRdPi3wIhYkxjeBaWh2rxwapn5Tu3IqOQ==",
			"dev": true,
			"requires": {
			  "commondir": "^1.0.1",
			  "make-dir": "^2.0.0",
			  "pkg-dir": "^3.0.0"
			}
		  },
		  "find-up": {
			"version": "3.0.0",
			"resolved": "https://registry.npmjs.org/find-up/-/find-up-3.0.0.tgz",
			"integrity": "sha512-1yD6RmLI1XBfxugvORwlck6f75tYL+iR0jqwsOrOxMZyGYqUuDhJ0l4AXdO1iX/FTs9cBAMEk1gWSEx1kSbylg==",
			"dev": true,
			"requires": {
			  "locate-path": "^3.0.0"
			}
		  },
		  "flat-cache": {
			"version": "2.0.1",
			"resolved": "https://registry.npmjs.org/flat-cache/-/flat-cache-2.0.1.tgz",
			"integrity": "sha512-LoQe6yDuUMDzQAEH8sgmh4Md6oZnc/7PjtwjNFSzveXqSHt6ka9fPBuso7IGf9Rz4uqnSnWiFH2B/zj24a5ReA==",
			"dev": true,
			"requires": {
			  "flatted": "^2.0.0",
			  "rimraf": "2.6.3",
			  "write": "1.0.3"
			},
			"dependencies": {
			  "rimraf": {
				"version": "2.6.3",
				"resolved": "https://registry.npmjs.org/rimraf/-/rimraf-2.6.3.tgz",
				"integrity": "sha512-mwqeW5XsA2qAejG46gYdENaxXjx9onRNCfn7L0duuP4hCuTIi/QO7PDK07KJfp1d+izWPrzEJDcSqBa0OZQriA==",
				"dev": true,
				"requires": {
				  "glob": "^7.1.3"
				}
			  }
			}
		  },
		  "flatted": {
			"version": "2.0.2",
			"resolved": "https://registry.npmjs.org/flatted/-/flatted-2.0.2.tgz",
			"integrity": "sha512-r5wGx7YeOwNWNlCA0wQ86zKyDLMQr+/RB8xy74M4hTphfmjlijTSSXGuH8rnvKZnfT9i+75zmd8jcKdMR4O6jA==",
			"dev": true
		  },
		  "flush-write-stream": {
			"version": "1.1.1",
			"resolved": "https://registry.npmjs.org/flush-write-stream/-/flush-write-stream-1.1.1.tgz",
			"integrity": "sha512-3Z4XhFZ3992uIq0XOqb9AreonueSYphE6oYbpt5+3u06JWklbsPkNv3ZKkP9Bz/r+1MWCaMoSQ28P85+1Yc77w==",
			"dev": true,
			"requires": {
			  "inherits": "^2.0.3",
			  "readable-stream": "^2.3.6"
			},
			"dependencies": {
			  "isarray": {
				"version": "1.0.0",
				"resolved": "https://registry.npmjs.org/isarray/-/isarray-1.0.0.tgz",
				"integrity": "sha1-u5NdSFgsuhaMBoNJV6VKPgcSTxE=",
				"dev": true
			  },
			  "readable-stream": {
				"version": "2.3.7",
				"resolved": "https://registry.npmjs.org/readable-stream/-/readable-stream-2.3.7.tgz",
				"integrity": "sha512-Ebho8K4jIbHAxnuxi7o42OrZgF/ZTNcsZj6nRKyUmkhLFq8CHItp/fy6hQZuZmP/n3yZ9VBUbp4zz/mX8hmYPw==",
				"dev": true,
				"requires": {
				  "core-util-is": "~1.0.0",
				  "inherits": "~2.0.3",
				  "isarray": "~1.0.0",
				  "process-nextick-args": "~2.0.0",
				  "safe-buffer": "~5.1.1",
				  "string_decoder": "~1.1.1",
				  "util-deprecate": "~1.0.1"
				}
			  },
			  "string_decoder": {
				"version": "1.1.1",
				"resolved": "https://registry.npmjs.org/string_decoder/-/string_decoder-1.1.1.tgz",
				"integrity": "sha512-n/ShnvDi6FHbbVfviro+WojiFzv+s8MPMHBczVePfUpDJLwoLT0ht1l4YwBCbi8pJAveEEdnkHyPyTP/mzRfwg==",
				"dev": true,
				"requires": {
				  "safe-buffer": "~5.1.0"
				}
			  }
			}
		  },
		  "follow-redirects": {
			"version": "1.5.10",
			"resolved": "https://registry.npmjs.org/follow-redirects/-/follow-redirects-1.5.10.tgz",
			"integrity": "sha512-0V5l4Cizzvqt5D44aTXbFZz+FtyXV1vrDN6qrelxtfYQKW0KO0W2T/hkE8xvGa/540LkZlkaUjO4ailYTFtHVQ==",
			"requires": {
			  "debug": "=3.1.0"
			}
		  },
		  "for-in": {
			"version": "1.0.2",
			"resolved": "https://registry.npmjs.org/for-in/-/for-in-1.0.2.tgz",
			"integrity": "sha1-gQaNKVqBQuwKxybG4iAMMPttXoA=",
			"dev": true
		  },
		  "forever-agent": {
			"version": "0.6.1",
			"resolved": "https://registry.npmjs.org/forever-agent/-/forever-agent-0.6.1.tgz",
			"integrity": "sha1-+8cfDEGt6zf5bFd60e1C2P2sypE=",
			"dev": true
		  },
		  "fork-ts-checker-webpack-plugin": {
			"version": "4.1.4",
			"resolved": "https://registry.npmjs.org/fork-ts-checker-webpack-plugin/-/fork-ts-checker-webpack-plugin-4.1.4.tgz",
			"integrity": "sha512-R0nTlZSyV0uCCzYe1kgR7Ve8mXyDvMm1pJwUFb6zzRVF5rTNb24G6gn2DFQy+W5aJYp2eq8aexpCOO+1SCyCSA==",
			"dev": true,
			"requires": {
			  "@babel/code-frame": "^7.5.5",
			  "chalk": "^2.4.1",
			  "micromatch": "^3.1.10",
			  "minimatch": "^3.0.4",
			  "semver": "^5.6.0",
			  "tapable": "^1.0.0",
			  "worker-rpc": "^0.1.0"
			}
		  },
		  "form-data": {
			"version": "2.3.3",
			"resolved": "https://registry.npmjs.org/form-data/-/form-data-2.3.3.tgz",
			"integrity": "sha512-1lLKB2Mu3aGP1Q/2eCOx0fNbRMe7XdwktwOruhfqqd0rIJWwN4Dh+E3hrPSlDCXnSR7UtZ1N38rVXm+6+MEhJQ==",
			"dev": true,
			"requires": {
			  "asynckit": "^0.4.0",
			  "combined-stream": "^1.0.6",
			  "mime-types": "^2.1.12"
			}
		  },
		  "formidable": {
			"version": "1.2.1",
			"resolved": "https://registry.npmjs.org/formidable/-/formidable-1.2.1.tgz",
			"integrity": "sha512-Fs9VRguL0gqGHkXS5GQiMCr1VhZBxz0JnJs4JmMp/2jL18Fmbzvv7vOFRU+U8TBkHEE/CX1qDXzJplVULgsLeg==",
			"dev": true
		  },
		  "forwarded": {
			"version": "0.1.2",
			"resolved": "https://registry.npmjs.org/forwarded/-/forwarded-0.1.2.tgz",
			"integrity": "sha1-mMI9qxF1ZXuMBXPozszZGw/xjIQ="
		  },
		  "fragment-cache": {
			"version": "0.2.1",
			"resolved": "https://registry.npmjs.org/fragment-cache/-/fragment-cache-0.2.1.tgz",
			"integrity": "sha1-QpD60n8T6Jvn8zeZxrxaCr//DRk=",
			"dev": true,
			"requires": {
			  "map-cache": "^0.2.2"
			}
		  },
		  "fresh": {
			"version": "0.5.2",
			"resolved": "https://registry.npmjs.org/fresh/-/fresh-0.5.2.tgz",
			"integrity": "sha1-PYyt2Q2XZWn6g1qx+OSyOhBWBac="
		  },
		  "from2": {
			"version": "2.3.0",
			"resolved": "https://registry.npmjs.org/from2/-/from2-2.3.0.tgz",
			"integrity": "sha1-i/tVAr3kpNNs/e6gB/zKIdfjgq8=",
			"dev": true,
			"requires": {
			  "inherits": "^2.0.1",
			  "readable-stream": "^2.0.0"
			},
			"dependencies": {
			  "isarray": {
				"version": "1.0.0",
				"resolved": "https://registry.npmjs.org/isarray/-/isarray-1.0.0.tgz",
				"integrity": "sha1-u5NdSFgsuhaMBoNJV6VKPgcSTxE=",
				"dev": true
			  },
			  "readable-stream": {
				"version": "2.3.7",
				"resolved": "https://registry.npmjs.org/readable-stream/-/readable-stream-2.3.7.tgz",
				"integrity": "sha512-Ebho8K4jIbHAxnuxi7o42OrZgF/ZTNcsZj6nRKyUmkhLFq8CHItp/fy6hQZuZmP/n3yZ9VBUbp4zz/mX8hmYPw==",
				"dev": true,
				"requires": {
				  "core-util-is": "~1.0.0",
				  "inherits": "~2.0.3",
				  "isarray": "~1.0.0",
				  "process-nextick-args": "~2.0.0",
				  "safe-buffer": "~5.1.1",
				  "string_decoder": "~1.1.1",
				  "util-deprecate": "~1.0.1"
				}
			  },
			  "string_decoder": {
				"version": "1.1.1",
				"resolved": "https://registry.npmjs.org/string_decoder/-/string_decoder-1.1.1.tgz",
				"integrity": "sha512-n/ShnvDi6FHbbVfviro+WojiFzv+s8MPMHBczVePfUpDJLwoLT0ht1l4YwBCbi8pJAveEEdnkHyPyTP/mzRfwg==",
				"dev": true,
				"requires": {
				  "safe-buffer": "~5.1.0"
				}
			  }
			}
		  },
		  "fs-extra": {
			"version": "8.1.0",
			"resolved": "https://registry.npmjs.org/fs-extra/-/fs-extra-8.1.0.tgz",
			"integrity": "sha512-yhlQgA6mnOJUKOsRUFsgJdQCvkKhcz8tlZG5HBQfReYZy46OwLcY+Zia0mtdHsOo9y/hP+CxMN0TU9QxoOtG4g==",
			"dev": true,
			"requires": {
			  "graceful-fs": "^4.2.0",
			  "jsonfile": "^4.0.0",
			  "universalify": "^0.1.0"
			}
		  },
		  "fs-write-stream-atomic": {
			"version": "1.0.10",
			"resolved": "https://registry.npmjs.org/fs-write-stream-atomic/-/fs-write-stream-atomic-1.0.10.tgz",
			"integrity": "sha1-tH31NJPvkR33VzHnCp3tAYnbQMk=",
			"dev": true,
			"requires": {
			  "graceful-fs": "^4.1.2",
			  "iferr": "^0.1.5",
			  "imurmurhash": "^0.1.4",
			  "readable-stream": "1 || 2"
			}
		  },
		  "fs.realpath": {
			"version": "1.0.0",
			"resolved": "https://registry.npmjs.org/fs.realpath/-/fs.realpath-1.0.0.tgz",
			"integrity": "sha1-FQStJSMVjKpA20onh8sBQRmU6k8="
		  },
		  "fsevents": {
			"version": "2.1.3",
			"resolved": "https://registry.npmjs.org/fsevents/-/fsevents-2.1.3.tgz",
			"integrity": "sha512-Auw9a4AxqWpa9GUfj370BMPzzyncfBABW8Mab7BGWBYDj4Isgq+cDKtx0i6u9jcX9pQDnswsaaOTgTmA5pEjuQ==",
			"dev": true,
			"optional": true
		  },
		  "function-bind": {
			"version": "1.1.1",
			"resolved": "https://registry.npmjs.org/function-bind/-/function-bind-1.1.1.tgz",
			"integrity": "sha512-yIovAzMX49sF8Yl58fSCWJ5svSLuaibPxXQJFLmBObTuCr0Mf1KiPopGM9NiFjiYBCbfaa2Fh6breQ6ANVTI0A==",
			"dev": true
		  },
		  "functional-red-black-tree": {
			"version": "1.0.1",
			"resolved": "https://registry.npmjs.org/functional-red-black-tree/-/functional-red-black-tree-1.0.1.tgz",
			"integrity": "sha1-GwqzvVU7Kg1jmdKcDj6gslIHgyc=",
			"dev": true
		  },
		  "gensync": {
			"version": "1.0.0-beta.1",
			"resolved": "https://registry.npmjs.org/gensync/-/gensync-1.0.0-beta.1.tgz",
			"integrity": "sha512-r8EC6NO1sngH/zdD9fiRDLdcgnbayXah+mLgManTaIZJqEC1MZstmnox8KpnI2/fxQwrp5OpCOYWLp4rBl4Jcg==",
			"dev": true
		  },
		  "get-caller-file": {
			"version": "2.0.5",
			"resolved": "https://registry.npmjs.org/get-caller-file/-/get-caller-file-2.0.5.tgz",
			"integrity": "sha512-DyFP3BM/3YHTQOCUL/w0OZHR0lpKeGrxotcHWcqNEdnltqFwXVfhEBQ94eIo34AfQpo0rGki4cyIiftY06h2Fg==",
			"dev": true
		  },
		  "get-stdin": {
			"version": "6.0.0",
			"resolved": "https://registry.npmjs.org/get-stdin/-/get-stdin-6.0.0.tgz",
			"integrity": "sha512-jp4tHawyV7+fkkSKyvjuLZswblUtz+SQKzSWnBbii16BuZksJlU1wuBYXY75r+duh/llF1ur6oNwi+2ZzjKZ7g==",
			"dev": true
		  },
		  "get-stream": {
			"version": "4.1.0",
			"resolved": "https://registry.npmjs.org/get-stream/-/get-stream-4.1.0.tgz",
			"integrity": "sha512-GMat4EJ5161kIy2HevLlr4luNjBgvmj413KaQA7jt4V8B4RDsfpHk7WQ9GVqfYyyx8OS/L66Kox+rJRNklLK7w==",
			"dev": true,
			"requires": {
			  "pump": "^3.0.0"
			}
		  },
		  "get-value": {
			"version": "2.0.6",
			"resolved": "https://registry.npmjs.org/get-value/-/get-value-2.0.6.tgz",
			"integrity": "sha1-3BXKHGcjh8p2vTesCjlbogQqLCg=",
			"dev": true
		  },
		  "getpass": {
			"version": "0.1.7",
			"resolved": "https://registry.npmjs.org/getpass/-/getpass-0.1.7.tgz",
			"integrity": "sha1-Xv+OPmhNVprkyysSgmBOi6YhSfo=",
			"dev": true,
			"requires": {
			  "assert-plus": "^1.0.0"
			}
		  },
		  "glob": {
			"version": "7.1.4",
			"resolved": "https://registry.npmjs.org/glob/-/glob-7.1.4.tgz",
			"integrity": "sha512-hkLPepehmnKk41pUGm3sYxoFs/umurYfYJCerbXEyFIWcAzvpipAgVkBqqT9RBKMGjnq6kMuyYwha6csxbiM1A==",
			"requires": {
			  "fs.realpath": "^1.0.0",
			  "inflight": "^1.0.4",
			  "inherits": "2",
			  "minimatch": "^3.0.4",
			  "once": "^1.3.0",
			  "path-is-absolute": "^1.0.0"
			}
		  },
		  "glob-parent": {
			"version": "5.1.1",
			"resolved": "https://registry.npmjs.org/glob-parent/-/glob-parent-5.1.1.tgz",
			"integrity": "sha512-FnI+VGOpnlGHWZxthPGR+QhR78fuiK0sNLkHQv+bL9fQi57lNNdquIbna/WrfROrolq8GK5Ek6BiMwqL/voRYQ==",
			"dev": true,
			"requires": {
			  "is-glob": "^4.0.1"
			}
		  },
		  "globals": {
			"version": "12.4.0",
			"resolved": "https://registry.npmjs.org/globals/-/globals-12.4.0.tgz",
			"integrity": "sha512-BWICuzzDvDoH54NHKCseDanAhE3CeDorgDL5MT6LMXXj2WCnd9UC2szdk4AWLfjdgNBCXLUanXYcpBBKOSWGwg==",
			"dev": true,
			"requires": {
			  "type-fest": "^0.8.1"
			},
			"dependencies": {
			  "type-fest": {
				"version": "0.8.1",
				"resolved": "https://registry.npmjs.org/type-fest/-/type-fest-0.8.1.tgz",
				"integrity": "sha512-4dbzIzqvjtgiM5rw1k5rEHtBANKmdudhGyBEajN01fEyhaAIhsoKNy6y7+IN93IfpFtwY9iqi7kD+xwKhQsNJA==",
				"dev": true
			  }
			}
		  },
		  "graceful-fs": {
			"version": "4.2.2",
			"resolved": "https://registry.npmjs.org/graceful-fs/-/graceful-fs-4.2.2.tgz",
			"integrity": "sha512-IItsdsea19BoLC7ELy13q1iJFNmd7ofZH5+X/pJr90/nRoPEX0DJo1dHDbgtYWOhJhcCgMDTOw84RZ72q6lB+Q==",
			"dev": true
		  },
		  "growly": {
			"version": "1.3.0",
			"resolved": "https://registry.npmjs.org/growly/-/growly-1.3.0.tgz",
			"integrity": "sha1-8QdIy+dq+WS3yWyTxrzCivEgwIE=",
			"dev": true,
			"optional": true
		  },
		  "har-schema": {
			"version": "2.0.0",
			"resolved": "https://registry.npmjs.org/har-schema/-/har-schema-2.0.0.tgz",
			"integrity": "sha1-qUwiJOvKwEeCoNkDVSHyRzW37JI=",
			"dev": true
		  },
		  "har-validator": {
			"version": "5.1.3",
			"resolved": "https://registry.npmjs.org/har-validator/-/har-validator-5.1.3.tgz",
			"integrity": "sha512-sNvOCzEQNr/qrvJgc3UG/kD4QtlHycrzwS+6mfTrrSq97BvaYcPZZI1ZSqGSPR73Cxn4LKTD4PttRwfU7jWq5g==",
			"dev": true,
			"requires": {
			  "ajv": "^6.5.5",
			  "har-schema": "^2.0.0"
			}
		  },
		  "has": {
			"version": "1.0.3",
			"resolved": "https://registry.npmjs.org/has/-/has-1.0.3.tgz",
			"integrity": "sha512-f2dvO0VU6Oej7RkWJGrehjbzMAjFp5/VKPp5tTpWIV4JHHZK1/BxbFRtf/siA2SWTe09caDmVtYYzWEIbBS4zw==",
			"dev": true,
			"requires": {
			  "function-bind": "^1.1.1"
			}
		  },
		  "has-flag": {
			"version": "3.0.0",
			"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-3.0.0.tgz",
			"integrity": "sha1-tdRU3CGZriJWmfNGfloH87lVuv0="
		  },
		  "has-symbols": {
			"version": "1.0.1",
			"resolved": "https://registry.npmjs.org/has-symbols/-/has-symbols-1.0.1.tgz",
			"integrity": "sha512-PLcsoqu++dmEIZB+6totNFKq/7Do+Z0u4oT0zKOJNl3lYK6vGwwu2hjHs+68OEZbTjiUE9bgOABXbP/GvrS0Kg==",
			"dev": true
		  },
		  "has-value": {
			"version": "1.0.0",
			"resolved": "https://registry.npmjs.org/has-value/-/has-value-1.0.0.tgz",
			"integrity": "sha1-GLKB2lhbHFxR3vJMkw7SmgvmsXc=",
			"dev": true,
			"requires": {
			  "get-value": "^2.0.6",
			  "has-values": "^1.0.0",
			  "isobject": "^3.0.0"
			}
		  },
		  "has-values": {
			"version": "1.0.0",
			"resolved": "https://registry.npmjs.org/has-values/-/has-values-1.0.0.tgz",
			"integrity": "sha1-lbC2P+whRmGab+V/51Yo1aOe/k8=",
			"dev": true,
			"requires": {
			  "is-number": "^3.0.0",
			  "kind-of": "^4.0.0"
			},
			"dependencies": {
			  "is-number": {
				"version": "3.0.0",
				"resolved": "https://registry.npmjs.org/is-number/-/is-number-3.0.0.tgz",
				"integrity": "sha1-JP1iAaR4LPUFYcgQJ2r8fRLXEZU=",
				"dev": true,
				"requires": {
				  "kind-of": "^3.0.2"
				},
				"dependencies": {
				  "kind-of": {
					"version": "3.2.2",
					"resolved": "https://registry.npmjs.org/kind-of/-/kind-of-3.2.2.tgz",
					"integrity": "sha1-MeohpzS6ubuw8yRm2JOupR5KPGQ=",
					"dev": true,
					"requires": {
					  "is-buffer": "^1.1.5"
					}
				  }
				}
			  },
			  "kind-of": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/kind-of/-/kind-of-4.0.0.tgz",
				"integrity": "sha1-IIE989cSkosgc3hpGkUGb65y3Vc=",
				"dev": true,
				"requires": {
				  "is-buffer": "^1.1.5"
				}
			  }
			}
		  },
		  "hash-base": {
			"version": "3.1.0",
			"resolved": "https://registry.npmjs.org/hash-base/-/hash-base-3.1.0.tgz",
			"integrity": "sha512-1nmYp/rhMDiE7AYkDw+lLwlAzz0AntGIe51F3RfFfEqyQ3feY2eI/NcwC6umIQVOASPMsWJLJScWKSSvzL9IVA==",
			"dev": true,
			"requires": {
			  "inherits": "^2.0.4",
			  "readable-stream": "^3.6.0",
			  "safe-buffer": "^5.2.0"
			},
			"dependencies": {
			  "inherits": {
				"version": "2.0.4",
				"resolved": "https://registry.npmjs.org/inherits/-/inherits-2.0.4.tgz",
				"integrity": "sha512-k/vGaX4/Yla3WzyMCvTQOXYeIHvqOKtnqBduzTHpzpQZzAskKMhZ2K+EnBiSM9zGSoIFeMpXKxa4dYeZIQqewQ==",
				"dev": true
			  },
			  "readable-stream": {
				"version": "3.6.0",
				"resolved": "https://registry.npmjs.org/readable-stream/-/readable-stream-3.6.0.tgz",
				"integrity": "sha512-BViHy7LKeTz4oNnkcLJ+lVSL6vpiFeX6/d3oSH8zCW7UxP2onchk+vTGB143xuFjHS3deTgkKoXXymXqymiIdA==",
				"dev": true,
				"requires": {
				  "inherits": "^2.0.3",
				  "string_decoder": "^1.1.1",
				  "util-deprecate": "^1.0.1"
				}
			  },
			  "safe-buffer": {
				"version": "5.2.1",
				"resolved": "https://registry.npmjs.org/safe-buffer/-/safe-buffer-5.2.1.tgz",
				"integrity": "sha512-rp3So07KcdmmKbGvgaNxQSJr7bGVSVk5S9Eq1F+ppbRo70+YeaDxkw5Dd8NPN+GD6bjnYm2VuPuCXmpuYvmCXQ==",
				"dev": true
			  },
			  "string_decoder": {
				"version": "1.3.0",
				"resolved": "https://registry.npmjs.org/string_decoder/-/string_decoder-1.3.0.tgz",
				"integrity": "sha512-hkRX8U1WjJFd8LsDJ2yQ/wWWxaopEsABU1XfkM8A+j0+85JAGppt16cr1Whg6KIbb4okU6Mql6BOj+uup/wKeA==",
				"dev": true,
				"requires": {
				  "safe-buffer": "~5.2.0"
				}
			  }
			}
		  },
		  "hash.js": {
			"version": "1.1.7",
			"resolved": "https://registry.npmjs.org/hash.js/-/hash.js-1.1.7.tgz",
			"integrity": "sha512-taOaskGt4z4SOANNseOviYDvjEJinIkRgmp7LbKP2YTTmVxWBl87s/uzK9r+44BclBSp2X7K1hqeNfz9JbBeXA==",
			"dev": true,
			"requires": {
			  "inherits": "^2.0.3",
			  "minimalistic-assert": "^1.0.1"
			}
		  },
		  "hmac-drbg": {
			"version": "1.0.1",
			"resolved": "https://registry.npmjs.org/hmac-drbg/-/hmac-drbg-1.0.1.tgz",
			"integrity": "sha1-0nRXAQJabHdabFRXk+1QL8DGSaE=",
			"dev": true,
			"requires": {
			  "hash.js": "^1.0.3",
			  "minimalistic-assert": "^1.0.0",
			  "minimalistic-crypto-utils": "^1.0.1"
			}
		  },
		  "hosted-git-info": {
			"version": "2.8.8",
			"resolved": "https://registry.npmjs.org/hosted-git-info/-/hosted-git-info-2.8.8.tgz",
			"integrity": "sha512-f/wzC2QaWBs7t9IYqB4T3sR1xviIViXJRJTWBlx2Gf3g0Xi5vI7Yy4koXQ1c9OYDGHN9sBy1DQ2AB8fqZBWhUg==",
			"dev": true
		  },
		  "html-encoding-sniffer": {
			"version": "2.0.1",
			"resolved": "https://registry.npmjs.org/html-encoding-sniffer/-/html-encoding-sniffer-2.0.1.tgz",
			"integrity": "sha512-D5JbOMBIR/TVZkubHT+OyT2705QvogUW4IBn6nHd756OwieSF9aDYFj4dv6HHEVGYbHaLETa3WggZYWWMyy3ZQ==",
			"dev": true,
			"requires": {
			  "whatwg-encoding": "^1.0.5"
			}
		  },
		  "html-escaper": {
			"version": "2.0.2",
			"resolved": "https://registry.npmjs.org/html-escaper/-/html-escaper-2.0.2.tgz",
			"integrity": "sha512-H2iMtd0I4Mt5eYiapRdIDjp+XzelXQ0tFE4JS7YFwFevXXMmOp9myNrUvCg0D6ws8iqkRPBfKHgbwig1SmlLfg==",
			"dev": true
		  },
		  "http-errors": {
			"version": "1.7.2",
			"resolved": "https://registry.npmjs.org/http-errors/-/http-errors-1.7.2.tgz",
			"integrity": "sha512-uUQBt3H/cSIVfch6i1EuPNy/YsRSOUBXTVfZ+yR7Zjez3qjBz6i9+i4zjNaoqcoFVI4lQJ5plg63TvGfRSDCRg==",
			"requires": {
			  "depd": "~1.1.2",
			  "inherits": "2.0.3",
			  "setprototypeof": "1.1.1",
			  "statuses": ">= 1.5.0 < 2",
			  "toidentifier": "1.0.0"
			}
		  },
		  "http-signature": {
			"version": "1.2.0",
			"resolved": "https://registry.npmjs.org/http-signature/-/http-signature-1.2.0.tgz",
			"integrity": "sha1-muzZJRFHcvPZW2WmCruPfBj7rOE=",
			"dev": true,
			"requires": {
			  "assert-plus": "^1.0.0",
			  "jsprim": "^1.2.2",
			  "sshpk": "^1.7.0"
			}
		  },
		  "https-browserify": {
			"version": "1.0.0",
			"resolved": "https://registry.npmjs.org/https-browserify/-/https-browserify-1.0.0.tgz",
			"integrity": "sha1-7AbBDgo0wPL68Zn3/X/Hj//QPHM=",
			"dev": true
		  },
		  "human-signals": {
			"version": "1.1.1",
			"resolved": "https://registry.npmjs.org/human-signals/-/human-signals-1.1.1.tgz",
			"integrity": "sha512-SEQu7vl8KjNL2eoGBLF3+wAjpsNfA9XMlXAYj/3EdaNfAlxKthD1xjEQfGOUhllCGGJVNY34bRr6lPINhNjyZw==",
			"dev": true
		  },
		  "iconv-lite": {
			"version": "0.4.24",
			"resolved": "https://registry.npmjs.org/iconv-lite/-/iconv-lite-0.4.24.tgz",
			"integrity": "sha512-v3MXnZAcvnywkTUEZomIActle7RXXeedOR31wwl7VlyoXO4Qi9arvSenNQWne1TcRwhCL1HwLI21bEqdpj8/rA==",
			"requires": {
			  "safer-buffer": ">= 2.1.2 < 3"
			}
		  },
		  "ieee754": {
			"version": "1.1.13",
			"resolved": "https://registry.npmjs.org/ieee754/-/ieee754-1.1.13.tgz",
			"integrity": "sha512-4vf7I2LYV/HaWerSo3XmlMkp5eZ83i+/CDluXi/IGTs/O1sejBNhTtnxzmRZfvOUqj7lZjqHkeTvpgSFDlWZTg==",
			"dev": true
		  },
		  "iferr": {
			"version": "0.1.5",
			"resolved": "https://registry.npmjs.org/iferr/-/iferr-0.1.5.tgz",
			"integrity": "sha1-xg7taebY/bazEEofy8ocGS3FtQE=",
			"dev": true
		  },
		  "ignore": {
			"version": "4.0.6",
			"resolved": "https://registry.npmjs.org/ignore/-/ignore-4.0.6.tgz",
			"integrity": "sha512-cyFDKrqc/YdcWFniJhzI42+AzS+gNwmUzOSFcRCQYwySuBBBy/KjuxWLZ/FHEH6Moq1NizMOBWyTcv8O4OZIMg==",
			"dev": true
		  },
		  "import-fresh": {
			"version": "3.2.1",
			"resolved": "https://registry.npmjs.org/import-fresh/-/import-fresh-3.2.1.tgz",
			"integrity": "sha512-6e1q1cnWP2RXD9/keSkxHScg508CdXqXWgWBaETNhyuBFz+kUZlKboh+ISK+bU++DmbHimVBrOz/zzPe0sZ3sQ==",
			"dev": true,
			"requires": {
			  "parent-module": "^1.0.0",
			  "resolve-from": "^4.0.0"
			}
		  },
		  "import-local": {
			"version": "3.0.2",
			"resolved": "https://registry.npmjs.org/import-local/-/import-local-3.0.2.tgz",
			"integrity": "sha512-vjL3+w0oulAVZ0hBHnxa/Nm5TAurf9YLQJDhqRZyqb+VKGOB6LU8t9H1Nr5CIo16vh9XfJTOoHwU0B71S557gA==",
			"dev": true,
			"requires": {
			  "pkg-dir": "^4.2.0",
			  "resolve-cwd": "^3.0.0"
			},
			"dependencies": {
			  "find-up": {
				"version": "4.1.0",
				"resolved": "https://registry.npmjs.org/find-up/-/find-up-4.1.0.tgz",
				"integrity": "sha512-PpOwAdQ/YlXQ2vj8a3h8IipDuYRi3wceVQQGYWxNINccq40Anw7BlsEXCMbt1Zt+OLA6Fq9suIpIWD0OsnISlw==",
				"dev": true,
				"requires": {
				  "locate-path": "^5.0.0",
				  "path-exists": "^4.0.0"
				}
			  },
			  "locate-path": {
				"version": "5.0.0",
				"resolved": "https://registry.npmjs.org/locate-path/-/locate-path-5.0.0.tgz",
				"integrity": "sha512-t7hw9pI+WvuwNJXwk5zVHpyhIqzg2qTlklJOf0mVxGSbe3Fp2VieZcduNYjaLDoy6p9uGpQEGWG87WpMKlNq8g==",
				"dev": true,
				"requires": {
				  "p-locate": "^4.1.0"
				}
			  },
			  "p-locate": {
				"version": "4.1.0",
				"resolved": "https://registry.npmjs.org/p-locate/-/p-locate-4.1.0.tgz",
				"integrity": "sha512-R79ZZ/0wAxKGu3oYMlz8jy/kbhsNrS7SKZ7PxEHBgJ5+F2mtFW2fK2cOtBh1cHYkQsbzFV7I+EoRKe6Yt0oK7A==",
				"dev": true,
				"requires": {
				  "p-limit": "^2.2.0"
				}
			  },
			  "path-exists": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/path-exists/-/path-exists-4.0.0.tgz",
				"integrity": "sha512-ak9Qy5Q7jYb2Wwcey5Fpvg2KoAc/ZIhLSLOSBmRmygPsGwkVVt0fZa0qrtMz+m6tJTAHfZQ8FnmB4MG4LWy7/w==",
				"dev": true
			  },
			  "pkg-dir": {
				"version": "4.2.0",
				"resolved": "https://registry.npmjs.org/pkg-dir/-/pkg-dir-4.2.0.tgz",
				"integrity": "sha512-HRDzbaKjC+AOWVXxAU/x54COGeIv9eb+6CkDSQoNTt4XyWoIJvuPsXizxu/Fr23EiekbtZwmh1IcIG/l/a10GQ==",
				"dev": true,
				"requires": {
				  "find-up": "^4.0.0"
				}
			  }
			}
		  },
		  "imurmurhash": {
			"version": "0.1.4",
			"resolved": "https://registry.npmjs.org/imurmurhash/-/imurmurhash-0.1.4.tgz",
			"integrity": "sha1-khi5srkoojixPcT7a21XbyMUU+o=",
			"dev": true
		  },
		  "infer-owner": {
			"version": "1.0.4",
			"resolved": "https://registry.npmjs.org/infer-owner/-/infer-owner-1.0.4.tgz",
			"integrity": "sha512-IClj+Xz94+d7irH5qRyfJonOdfTzuDaifE6ZPWfx0N0+/ATZCbuTPq2prFl526urkQd90WyUKIh1DfBQ2hMz9A==",
			"dev": true
		  },
		  "inflight": {
			"version": "1.0.6",
			"resolved": "https://registry.npmjs.org/inflight/-/inflight-1.0.6.tgz",
			"integrity": "sha1-Sb1jMdfQLQwJvJEKEHW6gWW1bfk=",
			"requires": {
			  "once": "^1.3.0",
			  "wrappy": "1"
			}
		  },
		  "inherits": {
			"version": "2.0.3",
			"resolved": "https://registry.npmjs.org/inherits/-/inherits-2.0.3.tgz",
			"integrity": "sha1-Yzwsg+PaQqUC9SRmAiSA9CCCYd4="
		  },
		  "inquirer": {
			"version": "7.1.0",
			"resolved": "https://registry.npmjs.org/inquirer/-/inquirer-7.1.0.tgz",
			"integrity": "sha512-5fJMWEmikSYu0nv/flMc475MhGbB7TSPd/2IpFV4I4rMklboCH2rQjYY5kKiYGHqUF9gvaambupcJFFG9dvReg==",
			"dev": true,
			"requires": {
			  "ansi-escapes": "^4.2.1",
			  "chalk": "^3.0.0",
			  "cli-cursor": "^3.1.0",
			  "cli-width": "^2.0.0",
			  "external-editor": "^3.0.3",
			  "figures": "^3.0.0",
			  "lodash": "^4.17.15",
			  "mute-stream": "0.0.8",
			  "run-async": "^2.4.0",
			  "rxjs": "^6.5.3",
			  "string-width": "^4.1.0",
			  "strip-ansi": "^6.0.0",
			  "through": "^2.3.6"
			},
			"dependencies": {
			  "ansi-regex": {
				"version": "5.0.0",
				"resolved": "https://registry.npmjs.org/ansi-regex/-/ansi-regex-5.0.0.tgz",
				"integrity": "sha512-bY6fj56OUQ0hU1KjFNDQuJFezqKdrAyFdIevADiqrWHwSlbmBNMHp5ak2f40Pm8JTFyM2mqxkG6ngkHO11f/lg==",
				"dev": true
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "3.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-3.0.0.tgz",
				"integrity": "sha512-4D3B6Wf41KOYRFdszmDqMCGq5VV/uMAB273JILmO+3jAlh8X4qDtdtgCR3fxtbLEMzSx22QdhnDcJvu2u1fVwg==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "strip-ansi": {
				"version": "6.0.0",
				"resolved": "https://registry.npmjs.org/strip-ansi/-/strip-ansi-6.0.0.tgz",
				"integrity": "sha512-AuvKTrTfQNYNIctbR1K/YGTR1756GycPsg7b9bdV9Duqur4gv6aKqHXah67Z8ImS7WEz5QVcOtlfW2rZEugt6w==",
				"dev": true,
				"requires": {
				  "ansi-regex": "^5.0.0"
				}
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "interpret": {
			"version": "1.3.0",
			"resolved": "https://registry.npmjs.org/interpret/-/interpret-1.3.0.tgz",
			"integrity": "sha512-RDVhhDkycLoSQtE9o0vpK/vOccVDsCbWVzRxArGYnlQLcihPl2loFbPyiH7CM0m2/ijOJU3+PZbnBPaB6NJ1MA==",
			"dev": true
		  },
		  "ip-regex": {
			"version": "2.1.0",
			"resolved": "https://registry.npmjs.org/ip-regex/-/ip-regex-2.1.0.tgz",
			"integrity": "sha1-+ni/XS5pE8kRzp+BnuUUa7bYROk=",
			"dev": true
		  },
		  "ipaddr.js": {
			"version": "1.9.1",
			"resolved": "https://registry.npmjs.org/ipaddr.js/-/ipaddr.js-1.9.1.tgz",
			"integrity": "sha512-0KI/607xoxSToH7GjN1FfSbLoU0+btTicjsQSWQlh/hZykN8KpmMf7uYwPW3R+akZ6R/w18ZlXSHBYXiYUPO3g=="
		  },
		  "is-accessor-descriptor": {
			"version": "0.1.6",
			"resolved": "https://registry.npmjs.org/is-accessor-descriptor/-/is-accessor-descriptor-0.1.6.tgz",
			"integrity": "sha1-qeEss66Nh2cn7u84Q/igiXtcmNY=",
			"dev": true,
			"requires": {
			  "kind-of": "^3.0.2"
			},
			"dependencies": {
			  "kind-of": {
				"version": "3.2.2",
				"resolved": "https://registry.npmjs.org/kind-of/-/kind-of-3.2.2.tgz",
				"integrity": "sha1-MeohpzS6ubuw8yRm2JOupR5KPGQ=",
				"dev": true,
				"requires": {
				  "is-buffer": "^1.1.5"
				}
			  }
			}
		  },
		  "is-arrayish": {
			"version": "0.2.1",
			"resolved": "https://registry.npmjs.org/is-arrayish/-/is-arrayish-0.2.1.tgz",
			"integrity": "sha1-d8mYQFJ6qOyxqLppe4BkWnqSap0=",
			"dev": true
		  },
		  "is-binary-path": {
			"version": "2.1.0",
			"resolved": "https://registry.npmjs.org/is-binary-path/-/is-binary-path-2.1.0.tgz",
			"integrity": "sha512-ZMERYes6pDydyuGidse7OsHxtbI7WVeUEozgR/g7rd0xUimYNlvZRE/K2MgZTjWy725IfelLeVcEM97mmtRGXw==",
			"dev": true,
			"requires": {
			  "binary-extensions": "^2.0.0"
			}
		  },
		  "is-buffer": {
			"version": "1.1.6",
			"resolved": "https://registry.npmjs.org/is-buffer/-/is-buffer-1.1.6.tgz",
			"integrity": "sha512-NcdALwpXkTm5Zvvbk7owOUSvVvBKDgKP5/ewfXEznmQFfs4ZRmanOeKBTjRVjka3QFoN6XJ+9F3USqfHqTaU5w==",
			"dev": true
		  },
		  "is-callable": {
			"version": "1.1.5",
			"resolved": "https://registry.npmjs.org/is-callable/-/is-callable-1.1.5.tgz",
			"integrity": "sha512-ESKv5sMCJB2jnHTWZ3O5itG+O128Hsus4K4Qh1h2/cgn2vbgnLSVqfV46AeJA9D5EeeLa9w81KUXMtn34zhX+Q==",
			"dev": true
		  },
		  "is-ci": {
			"version": "2.0.0",
			"resolved": "https://registry.npmjs.org/is-ci/-/is-ci-2.0.0.tgz",
			"integrity": "sha512-YfJT7rkpQB0updsdHLGWrvhBJfcfzNNawYDNIyQXJz0IViGf75O8EBPKSdvw2rF+LGCsX4FZ8tcr3b19LcZq4w==",
			"dev": true,
			"requires": {
			  "ci-info": "^2.0.0"
			}
		  },
		  "is-data-descriptor": {
			"version": "0.1.4",
			"resolved": "https://registry.npmjs.org/is-data-descriptor/-/is-data-descriptor-0.1.4.tgz",
			"integrity": "sha1-C17mSDiOLIYCgueT8YVv7D8wG1Y=",
			"dev": true,
			"requires": {
			  "kind-of": "^3.0.2"
			},
			"dependencies": {
			  "kind-of": {
				"version": "3.2.2",
				"resolved": "https://registry.npmjs.org/kind-of/-/kind-of-3.2.2.tgz",
				"integrity": "sha1-MeohpzS6ubuw8yRm2JOupR5KPGQ=",
				"dev": true,
				"requires": {
				  "is-buffer": "^1.1.5"
				}
			  }
			}
		  },
		  "is-date-object": {
			"version": "1.0.2",
			"resolved": "https://registry.npmjs.org/is-date-object/-/is-date-object-1.0.2.tgz",
			"integrity": "sha512-USlDT524woQ08aoZFzh3/Z6ch9Y/EWXEHQ/AaRN0SkKq4t2Jw2R2339tSXmwuVoY7LLlBCbOIlx2myP/L5zk0g==",
			"dev": true
		  },
		  "is-descriptor": {
			"version": "0.1.6",
			"resolved": "https://registry.npmjs.org/is-descriptor/-/is-descriptor-0.1.6.tgz",
			"integrity": "sha512-avDYr0SB3DwO9zsMov0gKCESFYqCnE4hq/4z3TdUlukEy5t9C0YRq7HLrsN52NAcqXKaepeCD0n+B0arnVG3Hg==",
			"dev": true,
			"requires": {
			  "is-accessor-descriptor": "^0.1.6",
			  "is-data-descriptor": "^0.1.4",
			  "kind-of": "^5.0.0"
			},
			"dependencies": {
			  "kind-of": {
				"version": "5.1.0",
				"resolved": "https://registry.npmjs.org/kind-of/-/kind-of-5.1.0.tgz",
				"integrity": "sha512-NGEErnH6F2vUuXDh+OlbcKW7/wOcfdRHaZ7VWtqCztfHri/++YKmP51OdWeGPuqCOba6kk2OTe5d02VmTB80Pw==",
				"dev": true
			  }
			}
		  },
		  "is-docker": {
			"version": "2.0.0",
			"resolved": "https://registry.npmjs.org/is-docker/-/is-docker-2.0.0.tgz",
			"integrity": "sha512-pJEdRugimx4fBMra5z2/5iRdZ63OhYV0vr0Dwm5+xtW4D1FvRkB8hamMIhnWfyJeDdyr/aa7BDyNbtG38VxgoQ==",
			"dev": true,
			"optional": true
		  },
		  "is-extendable": {
			"version": "0.1.1",
			"resolved": "https://registry.npmjs.org/is-extendable/-/is-extendable-0.1.1.tgz",
			"integrity": "sha1-YrEQ4omkcUGOPsNqYX1HLjAd/Ik=",
			"dev": true
		  },
		  "is-extglob": {
			"version": "2.1.1",
			"resolved": "https://registry.npmjs.org/is-extglob/-/is-extglob-2.1.1.tgz",
			"integrity": "sha1-qIwCU1eR8C7TfHahueqXc8gz+MI=",
			"dev": true
		  },
		  "is-fullwidth-code-point": {
			"version": "3.0.0",
			"resolved": "https://registry.npmjs.org/is-fullwidth-code-point/-/is-fullwidth-code-point-3.0.0.tgz",
			"integrity": "sha512-zymm5+u+sCsSWyD9qNaejV3DFvhCKclKdizYaJUuHA83RLjb7nSuGnddCHGv0hk+KY7BMAlsWeK4Ueg6EV6XQg==",
			"dev": true
		  },
		  "is-generator-fn": {
			"version": "2.1.0",
			"resolved": "https://registry.npmjs.org/is-generator-fn/-/is-generator-fn-2.1.0.tgz",
			"integrity": "sha512-cTIB4yPYL/Grw0EaSzASzg6bBy9gqCofvWN8okThAYIxKJZC+udlRAmGbM0XLeniEJSs8uEgHPGuHSe1XsOLSQ==",
			"dev": true
		  },
		  "is-glob": {
			"version": "4.0.1",
			"resolved": "https://registry.npmjs.org/is-glob/-/is-glob-4.0.1.tgz",
			"integrity": "sha512-5G0tKtBTFImOqDnLB2hG6Bp2qcKEFduo4tZu9MT/H6NQv/ghhy30o55ufafxJ/LdH79LLs2Kfrn85TLKyA7BUg==",
			"dev": true,
			"requires": {
			  "is-extglob": "^2.1.1"
			}
		  },
		  "is-interactive": {
			"version": "1.0.0",
			"resolved": "https://registry.npmjs.org/is-interactive/-/is-interactive-1.0.0.tgz",
			"integrity": "sha512-2HvIEKRoqS62guEC+qBjpvRubdX910WCMuJTZ+I9yvqKU2/12eSL549HMwtabb4oupdj2sMP50k+XJfB/8JE6w==",
			"dev": true
		  },
		  "is-number": {
			"version": "7.0.0",
			"resolved": "https://registry.npmjs.org/is-number/-/is-number-7.0.0.tgz",
			"integrity": "sha512-41Cifkg6e8TylSpdtTpeLVMqvSBEVzTttHvERD741+pnZ8ANv0004MRL43QKPDlK9cGvNp6NZWZUBlbGXYxxng==",
			"dev": true
		  },
		  "is-plain-object": {
			"version": "2.0.4",
			"resolved": "https://registry.npmjs.org/is-plain-object/-/is-plain-object-2.0.4.tgz",
			"integrity": "sha512-h5PpgXkWitc38BBMYawTYMWJHFZJVnBquFE57xFpjB8pJFiF6gZ+bU+WyI/yqXiFR5mdLsgYNaPe8uao6Uv9Og==",
			"dev": true,
			"requires": {
			  "isobject": "^3.0.1"
			}
		  },
		  "is-potential-custom-element-name": {
			"version": "1.0.0",
			"resolved": "https://registry.npmjs.org/is-potential-custom-element-name/-/is-potential-custom-element-name-1.0.0.tgz",
			"integrity": "sha1-DFLlS8yjkbssSUsh6GJtczbG45c=",
			"dev": true
		  },
		  "is-promise": {
			"version": "2.2.2",
			"resolved": "https://registry.npmjs.org/is-promise/-/is-promise-2.2.2.tgz",
			"integrity": "sha512-+lP4/6lKUBfQjZ2pdxThZvLUAafmZb8OAxFb8XXtiQmS35INgr85hdOGoEs124ez1FCnZJt6jau/T+alh58QFQ=="
		  },
		  "is-regex": {
			"version": "1.0.5",
			"resolved": "https://registry.npmjs.org/is-regex/-/is-regex-1.0.5.tgz",
			"integrity": "sha512-vlKW17SNq44owv5AQR3Cq0bQPEb8+kF3UKZ2fiZNOWtztYE5i0CzCZxFDwO58qAOWtxdBRVO/V5Qin1wjCqFYQ==",
			"dev": true,
			"requires": {
			  "has": "^1.0.3"
			}
		  },
		  "is-stream": {
			"version": "1.1.0",
			"resolved": "https://registry.npmjs.org/is-stream/-/is-stream-1.1.0.tgz",
			"integrity": "sha1-EtSj3U5o4Lec6428hBc66A2RykQ=",
			"dev": true
		  },
		  "is-string": {
			"version": "1.0.5",
			"resolved": "https://registry.npmjs.org/is-string/-/is-string-1.0.5.tgz",
			"integrity": "sha512-buY6VNRjhQMiF1qWDouloZlQbRhDPCebwxSjxMjxgemYT46YMd2NR0/H+fBhEfWX4A/w9TBJ+ol+okqJKFE6vQ==",
			"dev": true
		  },
		  "is-symbol": {
			"version": "1.0.3",
			"resolved": "https://registry.npmjs.org/is-symbol/-/is-symbol-1.0.3.tgz",
			"integrity": "sha512-OwijhaRSgqvhm/0ZdAcXNZt9lYdKFpcRDT5ULUuYXPoT794UNOdU+gpT6Rzo7b4V2HUl/op6GqY894AZwv9faQ==",
			"dev": true,
			"requires": {
			  "has-symbols": "^1.0.1"
			}
		  },
		  "is-typedarray": {
			"version": "1.0.0",
			"resolved": "https://registry.npmjs.org/is-typedarray/-/is-typedarray-1.0.0.tgz",
			"integrity": "sha1-5HnICFjfDBsR3dppQPlgEfzaSpo=",
			"dev": true
		  },
		  "is-windows": {
			"version": "1.0.2",
			"resolved": "https://registry.npmjs.org/is-windows/-/is-windows-1.0.2.tgz",
			"integrity": "sha512-eXK1UInq2bPmjyX6e3VHIzMLobc4J94i4AWn+Hpq3OU5KkrRC96OAcR3PRJ/pGu6m8TRnBHP9dkXQVsT/COVIA==",
			"dev": true
		  },
		  "is-wsl": {
			"version": "1.1.0",
			"resolved": "https://registry.npmjs.org/is-wsl/-/is-wsl-1.1.0.tgz",
			"integrity": "sha1-HxbkqiKwTRM2tmGIpmrzxgDDpm0=",
			"dev": true
		  },
		  "isarray": {
			"version": "0.0.1",
			"resolved": "https://registry.npmjs.org/isarray/-/isarray-0.0.1.tgz",
			"integrity": "sha1-ihis/Kmo9Bd+Cav8YDiTmwXR7t8="
		  },
		  "isexe": {
			"version": "2.0.0",
			"resolved": "https://registry.npmjs.org/isexe/-/isexe-2.0.0.tgz",
			"integrity": "sha1-6PvzdNxVb/iUehDcsFctYz8s+hA=",
			"dev": true
		  },
		  "isobject": {
			"version": "3.0.1",
			"resolved": "https://registry.npmjs.org/isobject/-/isobject-3.0.1.tgz",
			"integrity": "sha1-TkMekrEalzFjaqH5yNHMvP2reN8=",
			"dev": true
		  },
		  "isstream": {
			"version": "0.1.2",
			"resolved": "https://registry.npmjs.org/isstream/-/isstream-0.1.2.tgz",
			"integrity": "sha1-R+Y/evVa+m+S4VAOaQ64uFKcCZo=",
			"dev": true
		  },
		  "istanbul-lib-coverage": {
			"version": "3.0.0",
			"resolved": "https://registry.npmjs.org/istanbul-lib-coverage/-/istanbul-lib-coverage-3.0.0.tgz",
			"integrity": "sha512-UiUIqxMgRDET6eR+o5HbfRYP1l0hqkWOs7vNxC/mggutCMUIhWMm8gAHb8tHlyfD3/l6rlgNA5cKdDzEAf6hEg==",
			"dev": true
		  },
		  "istanbul-lib-instrument": {
			"version": "4.0.3",
			"resolved": "https://registry.npmjs.org/istanbul-lib-instrument/-/istanbul-lib-instrument-4.0.3.tgz",
			"integrity": "sha512-BXgQl9kf4WTCPCCpmFGoJkz/+uhvm7h7PFKUYxh7qarQd3ER33vHG//qaE8eN25l07YqZPpHXU9I09l/RD5aGQ==",
			"dev": true,
			"requires": {
			  "@babel/core": "^7.7.5",
			  "@istanbuljs/schema": "^0.1.2",
			  "istanbul-lib-coverage": "^3.0.0",
			  "semver": "^6.3.0"
			},
			"dependencies": {
			  "semver": {
				"version": "6.3.0",
				"resolved": "https://registry.npmjs.org/semver/-/semver-6.3.0.tgz",
				"integrity": "sha512-b39TBaTSfV6yBrapU89p5fKekE2m/NwnDocOVruQFS1/veMgdzuPcnOM34M6CwxW8jH/lxEa5rBoDeUwu5HHTw==",
				"dev": true
			  }
			}
		  },
		  "istanbul-lib-report": {
			"version": "3.0.0",
			"resolved": "https://registry.npmjs.org/istanbul-lib-report/-/istanbul-lib-report-3.0.0.tgz",
			"integrity": "sha512-wcdi+uAKzfiGT2abPpKZ0hSU1rGQjUQnLvtY5MpQ7QCTahD3VODhcu4wcfY1YtkGaDD5yuydOLINXsfbus9ROw==",
			"dev": true,
			"requires": {
			  "istanbul-lib-coverage": "^3.0.0",
			  "make-dir": "^3.0.0",
			  "supports-color": "^7.1.0"
			},
			"dependencies": {
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "make-dir": {
				"version": "3.1.0",
				"resolved": "https://registry.npmjs.org/make-dir/-/make-dir-3.1.0.tgz",
				"integrity": "sha512-g3FeP20LNwhALb/6Cz6Dd4F2ngze0jz7tbzrD2wAV+o9FeNHe4rL+yK2md0J/fiSf1sa1ADhXqi5+oVwOM/eGw==",
				"dev": true,
				"requires": {
				  "semver": "^6.0.0"
				}
			  },
			  "semver": {
				"version": "6.3.0",
				"resolved": "https://registry.npmjs.org/semver/-/semver-6.3.0.tgz",
				"integrity": "sha512-b39TBaTSfV6yBrapU89p5fKekE2m/NwnDocOVruQFS1/veMgdzuPcnOM34M6CwxW8jH/lxEa5rBoDeUwu5HHTw==",
				"dev": true
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "istanbul-lib-source-maps": {
			"version": "4.0.0",
			"resolved": "https://registry.npmjs.org/istanbul-lib-source-maps/-/istanbul-lib-source-maps-4.0.0.tgz",
			"integrity": "sha512-c16LpFRkR8vQXyHZ5nLpY35JZtzj1PQY1iZmesUbf1FZHbIupcWfjgOXBY9YHkLEQ6puz1u4Dgj6qmU/DisrZg==",
			"dev": true,
			"requires": {
			  "debug": "^4.1.1",
			  "istanbul-lib-coverage": "^3.0.0",
			  "source-map": "^0.6.1"
			},
			"dependencies": {
			  "debug": {
				"version": "4.1.1",
				"resolved": "https://registry.npmjs.org/debug/-/debug-4.1.1.tgz",
				"integrity": "sha512-pYAIzeRo8J6KPEaJ0VWOh5Pzkbw/RetuzehGM7QRRX5he4fPHx2rdKMB256ehJCkX+XRQm16eZLqLNS8RSZXZw==",
				"dev": true,
				"requires": {
				  "ms": "^2.1.1"
				}
			  },
			  "ms": {
				"version": "2.1.2",
				"resolved": "https://registry.npmjs.org/ms/-/ms-2.1.2.tgz",
				"integrity": "sha512-sGkPx+VjMtmA6MX27oA4FBFELFCZZ4S4XqeGOXCv68tT+jb3vk/RyaKWP0PTKyWtmLSM0b+adUTEvbs1PEaH2w==",
				"dev": true
			  },
			  "source-map": {
				"version": "0.6.1",
				"resolved": "https://registry.npmjs.org/source-map/-/source-map-0.6.1.tgz",
				"integrity": "sha512-UjgapumWlbMhkBgzT7Ykc5YXUT46F0iKu8SGXq0bcwP5dz/h0Plj6enJqjz1Zbq2l5WaqYnrVbwWOWMyF3F47g==",
				"dev": true
			  }
			}
		  },
		  "istanbul-reports": {
			"version": "3.0.2",
			"resolved": "https://registry.npmjs.org/istanbul-reports/-/istanbul-reports-3.0.2.tgz",
			"integrity": "sha512-9tZvz7AiR3PEDNGiV9vIouQ/EAcqMXFmkcA1CDFTwOB98OZVDL0PH9glHotf5Ugp6GCOTypfzGWI/OqjWNCRUw==",
			"dev": true,
			"requires": {
			  "html-escaper": "^2.0.0",
			  "istanbul-lib-report": "^3.0.0"
			}
		  },
		  "iterare": {
			"version": "1.2.0",
			"resolved": "https://registry.npmjs.org/iterare/-/iterare-1.2.0.tgz",
			"integrity": "sha512-RxMV9p/UzdK0Iplnd8mVgRvNdXlsTOiuDrqMRnDi3wIhbT+JP4xDquAX9ay13R3CH72NBzQ91KWe0+C168QAyQ=="
		  },
		  "jest": {
			"version": "26.0.1",
			"resolved": "https://registry.npmjs.org/jest/-/jest-26.0.1.tgz",
			"integrity": "sha512-29Q54kn5Bm7ZGKIuH2JRmnKl85YRigp0o0asTc6Sb6l2ch1DCXIeZTLLFy9ultJvhkTqbswF5DEx4+RlkmCxWg==",
			"dev": true,
			"requires": {
			  "@jest/core": "^26.0.1",
			  "import-local": "^3.0.2",
			  "jest-cli": "^26.0.1"
			},
			"dependencies": {
			  "@jest/types": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/@jest/types/-/types-26.0.1.tgz",
				"integrity": "sha512-IbtjvqI9+eS1qFnOIEL7ggWmT+iK/U+Vde9cGWtYb/b6XgKb3X44ZAe/z9YZzoAAZ/E92m0DqrilF934IGNnQA==",
				"dev": true,
				"requires": {
				  "@types/istanbul-lib-coverage": "^2.0.0",
				  "@types/istanbul-reports": "^1.1.1",
				  "@types/yargs": "^15.0.0",
				  "chalk": "^4.0.0"
				}
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-4.0.0.tgz",
				"integrity": "sha512-N9oWFcegS0sFr9oh1oz2d7Npos6vNoWW9HvtCg5N1KRFpUhaAhvTv5Y58g880fZaEYSNm3qDz8SU1UrGvp+n7A==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "graceful-fs": {
				"version": "4.2.4",
				"resolved": "https://registry.npmjs.org/graceful-fs/-/graceful-fs-4.2.4.tgz",
				"integrity": "sha512-WjKPNJF79dtJAVniUlGGWHYGz2jWxT6VhN/4m1NdkbZ2nOsEF+cI1Edgql5zCRhs/VsQYRvrXctxktVXZUkixw==",
				"dev": true
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "jest-cli": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/jest-cli/-/jest-cli-26.0.1.tgz",
				"integrity": "sha512-pFLfSOBcbG9iOZWaMK4Een+tTxi/Wcm34geqZEqrst9cZDkTQ1LZ2CnBrTlHWuYAiTMFr0EQeK52ScyFU8wK+w==",
				"dev": true,
				"requires": {
				  "@jest/core": "^26.0.1",
				  "@jest/test-result": "^26.0.1",
				  "@jest/types": "^26.0.1",
				  "chalk": "^4.0.0",
				  "exit": "^0.1.2",
				  "graceful-fs": "^4.2.4",
				  "import-local": "^3.0.2",
				  "is-ci": "^2.0.0",
				  "jest-config": "^26.0.1",
				  "jest-util": "^26.0.1",
				  "jest-validate": "^26.0.1",
				  "prompts": "^2.0.1",
				  "yargs": "^15.3.1"
				}
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "jest-changed-files": {
			"version": "26.0.1",
			"resolved": "https://registry.npmjs.org/jest-changed-files/-/jest-changed-files-26.0.1.tgz",
			"integrity": "sha512-q8LP9Sint17HaE2LjxQXL+oYWW/WeeXMPE2+Op9X3mY8IEGFVc14xRxFjUuXUbcPAlDLhtWdIEt59GdQbn76Hw==",
			"dev": true,
			"requires": {
			  "@jest/types": "^26.0.1",
			  "execa": "^4.0.0",
			  "throat": "^5.0.0"
			},
			"dependencies": {
			  "@jest/types": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/@jest/types/-/types-26.0.1.tgz",
				"integrity": "sha512-IbtjvqI9+eS1qFnOIEL7ggWmT+iK/U+Vde9cGWtYb/b6XgKb3X44ZAe/z9YZzoAAZ/E92m0DqrilF934IGNnQA==",
				"dev": true,
				"requires": {
				  "@types/istanbul-lib-coverage": "^2.0.0",
				  "@types/istanbul-reports": "^1.1.1",
				  "@types/yargs": "^15.0.0",
				  "chalk": "^4.0.0"
				}
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-4.0.0.tgz",
				"integrity": "sha512-N9oWFcegS0sFr9oh1oz2d7Npos6vNoWW9HvtCg5N1KRFpUhaAhvTv5Y58g880fZaEYSNm3qDz8SU1UrGvp+n7A==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "cross-spawn": {
				"version": "7.0.2",
				"resolved": "https://registry.npmjs.org/cross-spawn/-/cross-spawn-7.0.2.tgz",
				"integrity": "sha512-PD6G8QG3S4FK/XCGFbEQrDqO2AnMMsy0meR7lerlIOHAAbkuavGU/pOqprrlvfTNjvowivTeBsjebAL0NSoMxw==",
				"dev": true,
				"requires": {
				  "path-key": "^3.1.0",
				  "shebang-command": "^2.0.0",
				  "which": "^2.0.1"
				}
			  },
			  "execa": {
				"version": "4.0.1",
				"resolved": "https://registry.npmjs.org/execa/-/execa-4.0.1.tgz",
				"integrity": "sha512-SCjM/zlBdOK8Q5TIjOn6iEHZaPHFsMoTxXQ2nvUvtPnuohz3H2dIozSg+etNR98dGoYUp2ENSKLL/XaMmbxVgw==",
				"dev": true,
				"requires": {
				  "cross-spawn": "^7.0.0",
				  "get-stream": "^5.0.0",
				  "human-signals": "^1.1.1",
				  "is-stream": "^2.0.0",
				  "merge-stream": "^2.0.0",
				  "npm-run-path": "^4.0.0",
				  "onetime": "^5.1.0",
				  "signal-exit": "^3.0.2",
				  "strip-final-newline": "^2.0.0"
				}
			  },
			  "get-stream": {
				"version": "5.1.0",
				"resolved": "https://registry.npmjs.org/get-stream/-/get-stream-5.1.0.tgz",
				"integrity": "sha512-EXr1FOzrzTfGeL0gQdeFEvOMm2mzMOglyiOXSTpPC+iAjAKftbr3jpCMWynogwYnM+eSj9sHGc6wjIcDvYiygw==",
				"dev": true,
				"requires": {
				  "pump": "^3.0.0"
				}
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "is-stream": {
				"version": "2.0.0",
				"resolved": "https://registry.npmjs.org/is-stream/-/is-stream-2.0.0.tgz",
				"integrity": "sha512-XCoy+WlUr7d1+Z8GgSuXmpuUFC9fOhRXglJMx+dwLKTkL44Cjd4W1Z5P+BQZpr+cR93aGP4S/s7Ftw6Nd/kiEw==",
				"dev": true
			  },
			  "npm-run-path": {
				"version": "4.0.1",
				"resolved": "https://registry.npmjs.org/npm-run-path/-/npm-run-path-4.0.1.tgz",
				"integrity": "sha512-S48WzZW777zhNIrn7gxOlISNAqi9ZC/uQFnRdbeIHhZhCA6UqpkOT8T1G7BvfdgP4Er8gF4sUbaS0i7QvIfCWw==",
				"dev": true,
				"requires": {
				  "path-key": "^3.0.0"
				}
			  },
			  "path-key": {
				"version": "3.1.1",
				"resolved": "https://registry.npmjs.org/path-key/-/path-key-3.1.1.tgz",
				"integrity": "sha512-ojmeN0qd+y0jszEtoY48r0Peq5dwMEkIlCOu6Q5f41lfkswXuKtYrhgoTpLnyIcHm24Uhqx+5Tqm2InSwLhE6Q==",
				"dev": true
			  },
			  "shebang-command": {
				"version": "2.0.0",
				"resolved": "https://registry.npmjs.org/shebang-command/-/shebang-command-2.0.0.tgz",
				"integrity": "sha512-kHxr2zZpYtdmrN1qDjrrX/Z1rR1kG8Dx+gkpK1G4eXmvXswmcE1hTWBWYUzlraYw1/yZp6YuDY77YtvbN0dmDA==",
				"dev": true,
				"requires": {
				  "shebang-regex": "^3.0.0"
				}
			  },
			  "shebang-regex": {
				"version": "3.0.0",
				"resolved": "https://registry.npmjs.org/shebang-regex/-/shebang-regex-3.0.0.tgz",
				"integrity": "sha512-7++dFhtcx3353uBaq8DDR4NuxBetBzC7ZQOhmTQInHEd6bSrXdiEyzCvG07Z44UYdLShWUyXt5M/yhz8ekcb1A==",
				"dev": true
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  },
			  "which": {
				"version": "2.0.2",
				"resolved": "https://registry.npmjs.org/which/-/which-2.0.2.tgz",
				"integrity": "sha512-BLI3Tl1TW3Pvl70l3yq3Y64i+awpwXqsGBYWkkqMtnbXgrMD+yj7rhW0kuEDxzJaYXGjEW5ogapKNMEKNMjibA==",
				"dev": true,
				"requires": {
				  "isexe": "^2.0.0"
				}
			  }
			}
		  },
		  "jest-config": {
			"version": "26.0.1",
			"resolved": "https://registry.npmjs.org/jest-config/-/jest-config-26.0.1.tgz",
			"integrity": "sha512-9mWKx2L1LFgOXlDsC4YSeavnblN6A4CPfXFiobq+YYLaBMymA/SczN7xYTSmLaEYHZOcB98UdoN4m5uNt6tztg==",
			"dev": true,
			"requires": {
			  "@babel/core": "^7.1.0",
			  "@jest/test-sequencer": "^26.0.1",
			  "@jest/types": "^26.0.1",
			  "babel-jest": "^26.0.1",
			  "chalk": "^4.0.0",
			  "deepmerge": "^4.2.2",
			  "glob": "^7.1.1",
			  "graceful-fs": "^4.2.4",
			  "jest-environment-jsdom": "^26.0.1",
			  "jest-environment-node": "^26.0.1",
			  "jest-get-type": "^26.0.0",
			  "jest-jasmine2": "^26.0.1",
			  "jest-regex-util": "^26.0.0",
			  "jest-resolve": "^26.0.1",
			  "jest-util": "^26.0.1",
			  "jest-validate": "^26.0.1",
			  "micromatch": "^4.0.2",
			  "pretty-format": "^26.0.1"
			},
			"dependencies": {
			  "@jest/types": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/@jest/types/-/types-26.0.1.tgz",
				"integrity": "sha512-IbtjvqI9+eS1qFnOIEL7ggWmT+iK/U+Vde9cGWtYb/b6XgKb3X44ZAe/z9YZzoAAZ/E92m0DqrilF934IGNnQA==",
				"dev": true,
				"requires": {
				  "@types/istanbul-lib-coverage": "^2.0.0",
				  "@types/istanbul-reports": "^1.1.1",
				  "@types/yargs": "^15.0.0",
				  "chalk": "^4.0.0"
				}
			  },
			  "ansi-regex": {
				"version": "5.0.0",
				"resolved": "https://registry.npmjs.org/ansi-regex/-/ansi-regex-5.0.0.tgz",
				"integrity": "sha512-bY6fj56OUQ0hU1KjFNDQuJFezqKdrAyFdIevADiqrWHwSlbmBNMHp5ak2f40Pm8JTFyM2mqxkG6ngkHO11f/lg==",
				"dev": true
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-4.0.0.tgz",
				"integrity": "sha512-N9oWFcegS0sFr9oh1oz2d7Npos6vNoWW9HvtCg5N1KRFpUhaAhvTv5Y58g880fZaEYSNm3qDz8SU1UrGvp+n7A==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "graceful-fs": {
				"version": "4.2.4",
				"resolved": "https://registry.npmjs.org/graceful-fs/-/graceful-fs-4.2.4.tgz",
				"integrity": "sha512-WjKPNJF79dtJAVniUlGGWHYGz2jWxT6VhN/4m1NdkbZ2nOsEF+cI1Edgql5zCRhs/VsQYRvrXctxktVXZUkixw==",
				"dev": true
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "jest-get-type": {
				"version": "26.0.0",
				"resolved": "https://registry.npmjs.org/jest-get-type/-/jest-get-type-26.0.0.tgz",
				"integrity": "sha512-zRc1OAPnnws1EVfykXOj19zo2EMw5Hi6HLbFCSjpuJiXtOWAYIjNsHVSbpQ8bDX7L5BGYGI8m+HmKdjHYFF0kg==",
				"dev": true
			  },
			  "micromatch": {
				"version": "4.0.2",
				"resolved": "https://registry.npmjs.org/micromatch/-/micromatch-4.0.2.tgz",
				"integrity": "sha512-y7FpHSbMUMoyPbYUSzO6PaZ6FyRnQOpHuKwbo1G+Knck95XVU4QAiKdGEnj5wwoS7PlOgthX/09u5iFJ+aYf5Q==",
				"dev": true,
				"requires": {
				  "braces": "^3.0.1",
				  "picomatch": "^2.0.5"
				}
			  },
			  "pretty-format": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/pretty-format/-/pretty-format-26.0.1.tgz",
				"integrity": "sha512-SWxz6MbupT3ZSlL0Po4WF/KujhQaVehijR2blyRDCzk9e45EaYMVhMBn49fnRuHxtkSpXTes1GxNpVmH86Bxfw==",
				"dev": true,
				"requires": {
				  "@jest/types": "^26.0.1",
				  "ansi-regex": "^5.0.0",
				  "ansi-styles": "^4.0.0",
				  "react-is": "^16.12.0"
				}
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "jest-diff": {
			"version": "25.5.0",
			"resolved": "https://registry.npmjs.org/jest-diff/-/jest-diff-25.5.0.tgz",
			"integrity": "sha512-z1kygetuPiREYdNIumRpAHY6RXiGmp70YHptjdaxTWGmA085W3iCnXNx0DhflK3vwrKmrRWyY1wUpkPMVxMK7A==",
			"dev": true,
			"requires": {
			  "chalk": "^3.0.0",
			  "diff-sequences": "^25.2.6",
			  "jest-get-type": "^25.2.6",
			  "pretty-format": "^25.5.0"
			},
			"dependencies": {
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "3.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-3.0.0.tgz",
				"integrity": "sha512-4D3B6Wf41KOYRFdszmDqMCGq5VV/uMAB273JILmO+3jAlh8X4qDtdtgCR3fxtbLEMzSx22QdhnDcJvu2u1fVwg==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "jest-docblock": {
			"version": "26.0.0",
			"resolved": "https://registry.npmjs.org/jest-docblock/-/jest-docblock-26.0.0.tgz",
			"integrity": "sha512-RDZ4Iz3QbtRWycd8bUEPxQsTlYazfYn/h5R65Fc6gOfwozFhoImx+affzky/FFBuqISPTqjXomoIGJVKBWoo0w==",
			"dev": true,
			"requires": {
			  "detect-newline": "^3.0.0"
			}
		  },
		  "jest-each": {
			"version": "26.0.1",
			"resolved": "https://registry.npmjs.org/jest-each/-/jest-each-26.0.1.tgz",
			"integrity": "sha512-OTgJlwXCAR8NIWaXFL5DBbeS4QIYPuNASkzSwMCJO+ywo9BEa6TqkaSWsfR7VdbMLdgYJqSfQcIyjJCNwl5n4Q==",
			"dev": true,
			"requires": {
			  "@jest/types": "^26.0.1",
			  "chalk": "^4.0.0",
			  "jest-get-type": "^26.0.0",
			  "jest-util": "^26.0.1",
			  "pretty-format": "^26.0.1"
			},
			"dependencies": {
			  "@jest/types": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/@jest/types/-/types-26.0.1.tgz",
				"integrity": "sha512-IbtjvqI9+eS1qFnOIEL7ggWmT+iK/U+Vde9cGWtYb/b6XgKb3X44ZAe/z9YZzoAAZ/E92m0DqrilF934IGNnQA==",
				"dev": true,
				"requires": {
				  "@types/istanbul-lib-coverage": "^2.0.0",
				  "@types/istanbul-reports": "^1.1.1",
				  "@types/yargs": "^15.0.0",
				  "chalk": "^4.0.0"
				}
			  },
			  "ansi-regex": {
				"version": "5.0.0",
				"resolved": "https://registry.npmjs.org/ansi-regex/-/ansi-regex-5.0.0.tgz",
				"integrity": "sha512-bY6fj56OUQ0hU1KjFNDQuJFezqKdrAyFdIevADiqrWHwSlbmBNMHp5ak2f40Pm8JTFyM2mqxkG6ngkHO11f/lg==",
				"dev": true
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-4.0.0.tgz",
				"integrity": "sha512-N9oWFcegS0sFr9oh1oz2d7Npos6vNoWW9HvtCg5N1KRFpUhaAhvTv5Y58g880fZaEYSNm3qDz8SU1UrGvp+n7A==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "jest-get-type": {
				"version": "26.0.0",
				"resolved": "https://registry.npmjs.org/jest-get-type/-/jest-get-type-26.0.0.tgz",
				"integrity": "sha512-zRc1OAPnnws1EVfykXOj19zo2EMw5Hi6HLbFCSjpuJiXtOWAYIjNsHVSbpQ8bDX7L5BGYGI8m+HmKdjHYFF0kg==",
				"dev": true
			  },
			  "pretty-format": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/pretty-format/-/pretty-format-26.0.1.tgz",
				"integrity": "sha512-SWxz6MbupT3ZSlL0Po4WF/KujhQaVehijR2blyRDCzk9e45EaYMVhMBn49fnRuHxtkSpXTes1GxNpVmH86Bxfw==",
				"dev": true,
				"requires": {
				  "@jest/types": "^26.0.1",
				  "ansi-regex": "^5.0.0",
				  "ansi-styles": "^4.0.0",
				  "react-is": "^16.12.0"
				}
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "jest-environment-jsdom": {
			"version": "26.0.1",
			"resolved": "https://registry.npmjs.org/jest-environment-jsdom/-/jest-environment-jsdom-26.0.1.tgz",
			"integrity": "sha512-u88NJa3aptz2Xix2pFhihRBAatwZHWwSiRLBDBQE1cdJvDjPvv7ZGA0NQBxWwDDn7D0g1uHqxM8aGgfA9Bx49g==",
			"dev": true,
			"requires": {
			  "@jest/environment": "^26.0.1",
			  "@jest/fake-timers": "^26.0.1",
			  "@jest/types": "^26.0.1",
			  "jest-mock": "^26.0.1",
			  "jest-util": "^26.0.1",
			  "jsdom": "^16.2.2"
			},
			"dependencies": {
			  "@jest/types": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/@jest/types/-/types-26.0.1.tgz",
				"integrity": "sha512-IbtjvqI9+eS1qFnOIEL7ggWmT+iK/U+Vde9cGWtYb/b6XgKb3X44ZAe/z9YZzoAAZ/E92m0DqrilF934IGNnQA==",
				"dev": true,
				"requires": {
				  "@types/istanbul-lib-coverage": "^2.0.0",
				  "@types/istanbul-reports": "^1.1.1",
				  "@types/yargs": "^15.0.0",
				  "chalk": "^4.0.0"
				}
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-4.0.0.tgz",
				"integrity": "sha512-N9oWFcegS0sFr9oh1oz2d7Npos6vNoWW9HvtCg5N1KRFpUhaAhvTv5Y58g880fZaEYSNm3qDz8SU1UrGvp+n7A==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "jest-environment-node": {
			"version": "26.0.1",
			"resolved": "https://registry.npmjs.org/jest-environment-node/-/jest-environment-node-26.0.1.tgz",
			"integrity": "sha512-4FRBWcSn5yVo0KtNav7+5NH5Z/tEgDLp7VRQVS5tCouWORxj+nI+1tOLutM07Zb2Qi7ja+HEDoOUkjBSWZg/IQ==",
			"dev": true,
			"requires": {
			  "@jest/environment": "^26.0.1",
			  "@jest/fake-timers": "^26.0.1",
			  "@jest/types": "^26.0.1",
			  "jest-mock": "^26.0.1",
			  "jest-util": "^26.0.1"
			},
			"dependencies": {
			  "@jest/types": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/@jest/types/-/types-26.0.1.tgz",
				"integrity": "sha512-IbtjvqI9+eS1qFnOIEL7ggWmT+iK/U+Vde9cGWtYb/b6XgKb3X44ZAe/z9YZzoAAZ/E92m0DqrilF934IGNnQA==",
				"dev": true,
				"requires": {
				  "@types/istanbul-lib-coverage": "^2.0.0",
				  "@types/istanbul-reports": "^1.1.1",
				  "@types/yargs": "^15.0.0",
				  "chalk": "^4.0.0"
				}
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-4.0.0.tgz",
				"integrity": "sha512-N9oWFcegS0sFr9oh1oz2d7Npos6vNoWW9HvtCg5N1KRFpUhaAhvTv5Y58g880fZaEYSNm3qDz8SU1UrGvp+n7A==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "jest-get-type": {
			"version": "25.2.6",
			"resolved": "https://registry.npmjs.org/jest-get-type/-/jest-get-type-25.2.6.tgz",
			"integrity": "sha512-DxjtyzOHjObRM+sM1knti6or+eOgcGU4xVSb2HNP1TqO4ahsT+rqZg+nyqHWJSvWgKC5cG3QjGFBqxLghiF/Ig==",
			"dev": true
		  },
		  "jest-haste-map": {
			"version": "26.0.1",
			"resolved": "https://registry.npmjs.org/jest-haste-map/-/jest-haste-map-26.0.1.tgz",
			"integrity": "sha512-J9kBl/EdjmDsvyv7CiyKY5+DsTvVOScenprz/fGqfLg/pm1gdjbwwQ98nW0t+OIt+f+5nAVaElvn/6wP5KO7KA==",
			"dev": true,
			"requires": {
			  "@jest/types": "^26.0.1",
			  "@types/graceful-fs": "^4.1.2",
			  "anymatch": "^3.0.3",
			  "fb-watchman": "^2.0.0",
			  "fsevents": "^2.1.2",
			  "graceful-fs": "^4.2.4",
			  "jest-serializer": "^26.0.0",
			  "jest-util": "^26.0.1",
			  "jest-worker": "^26.0.0",
			  "micromatch": "^4.0.2",
			  "sane": "^4.0.3",
			  "walker": "^1.0.7",
			  "which": "^2.0.2"
			},
			"dependencies": {
			  "@jest/types": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/@jest/types/-/types-26.0.1.tgz",
				"integrity": "sha512-IbtjvqI9+eS1qFnOIEL7ggWmT+iK/U+Vde9cGWtYb/b6XgKb3X44ZAe/z9YZzoAAZ/E92m0DqrilF934IGNnQA==",
				"dev": true,
				"requires": {
				  "@types/istanbul-lib-coverage": "^2.0.0",
				  "@types/istanbul-reports": "^1.1.1",
				  "@types/yargs": "^15.0.0",
				  "chalk": "^4.0.0"
				}
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-4.0.0.tgz",
				"integrity": "sha512-N9oWFcegS0sFr9oh1oz2d7Npos6vNoWW9HvtCg5N1KRFpUhaAhvTv5Y58g880fZaEYSNm3qDz8SU1UrGvp+n7A==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "graceful-fs": {
				"version": "4.2.4",
				"resolved": "https://registry.npmjs.org/graceful-fs/-/graceful-fs-4.2.4.tgz",
				"integrity": "sha512-WjKPNJF79dtJAVniUlGGWHYGz2jWxT6VhN/4m1NdkbZ2nOsEF+cI1Edgql5zCRhs/VsQYRvrXctxktVXZUkixw==",
				"dev": true
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "micromatch": {
				"version": "4.0.2",
				"resolved": "https://registry.npmjs.org/micromatch/-/micromatch-4.0.2.tgz",
				"integrity": "sha512-y7FpHSbMUMoyPbYUSzO6PaZ6FyRnQOpHuKwbo1G+Knck95XVU4QAiKdGEnj5wwoS7PlOgthX/09u5iFJ+aYf5Q==",
				"dev": true,
				"requires": {
				  "braces": "^3.0.1",
				  "picomatch": "^2.0.5"
				}
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  },
			  "which": {
				"version": "2.0.2",
				"resolved": "https://registry.npmjs.org/which/-/which-2.0.2.tgz",
				"integrity": "sha512-BLI3Tl1TW3Pvl70l3yq3Y64i+awpwXqsGBYWkkqMtnbXgrMD+yj7rhW0kuEDxzJaYXGjEW5ogapKNMEKNMjibA==",
				"dev": true,
				"requires": {
				  "isexe": "^2.0.0"
				}
			  }
			}
		  },
		  "jest-jasmine2": {
			"version": "26.0.1",
			"resolved": "https://registry.npmjs.org/jest-jasmine2/-/jest-jasmine2-26.0.1.tgz",
			"integrity": "sha512-ILaRyiWxiXOJ+RWTKupzQWwnPaeXPIoLS5uW41h18varJzd9/7I0QJGqg69fhTT1ev9JpSSo9QtalriUN0oqOg==",
			"dev": true,
			"requires": {
			  "@babel/traverse": "^7.1.0",
			  "@jest/environment": "^26.0.1",
			  "@jest/source-map": "^26.0.0",
			  "@jest/test-result": "^26.0.1",
			  "@jest/types": "^26.0.1",
			  "chalk": "^4.0.0",
			  "co": "^4.6.0",
			  "expect": "^26.0.1",
			  "is-generator-fn": "^2.0.0",
			  "jest-each": "^26.0.1",
			  "jest-matcher-utils": "^26.0.1",
			  "jest-message-util": "^26.0.1",
			  "jest-runtime": "^26.0.1",
			  "jest-snapshot": "^26.0.1",
			  "jest-util": "^26.0.1",
			  "pretty-format": "^26.0.1",
			  "throat": "^5.0.0"
			},
			"dependencies": {
			  "@jest/types": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/@jest/types/-/types-26.0.1.tgz",
				"integrity": "sha512-IbtjvqI9+eS1qFnOIEL7ggWmT+iK/U+Vde9cGWtYb/b6XgKb3X44ZAe/z9YZzoAAZ/E92m0DqrilF934IGNnQA==",
				"dev": true,
				"requires": {
				  "@types/istanbul-lib-coverage": "^2.0.0",
				  "@types/istanbul-reports": "^1.1.1",
				  "@types/yargs": "^15.0.0",
				  "chalk": "^4.0.0"
				}
			  },
			  "ansi-regex": {
				"version": "5.0.0",
				"resolved": "https://registry.npmjs.org/ansi-regex/-/ansi-regex-5.0.0.tgz",
				"integrity": "sha512-bY6fj56OUQ0hU1KjFNDQuJFezqKdrAyFdIevADiqrWHwSlbmBNMHp5ak2f40Pm8JTFyM2mqxkG6ngkHO11f/lg==",
				"dev": true
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-4.0.0.tgz",
				"integrity": "sha512-N9oWFcegS0sFr9oh1oz2d7Npos6vNoWW9HvtCg5N1KRFpUhaAhvTv5Y58g880fZaEYSNm3qDz8SU1UrGvp+n7A==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "pretty-format": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/pretty-format/-/pretty-format-26.0.1.tgz",
				"integrity": "sha512-SWxz6MbupT3ZSlL0Po4WF/KujhQaVehijR2blyRDCzk9e45EaYMVhMBn49fnRuHxtkSpXTes1GxNpVmH86Bxfw==",
				"dev": true,
				"requires": {
				  "@jest/types": "^26.0.1",
				  "ansi-regex": "^5.0.0",
				  "ansi-styles": "^4.0.0",
				  "react-is": "^16.12.0"
				}
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "jest-leak-detector": {
			"version": "26.0.1",
			"resolved": "https://registry.npmjs.org/jest-leak-detector/-/jest-leak-detector-26.0.1.tgz",
			"integrity": "sha512-93FR8tJhaYIWrWsbmVN1pQ9ZNlbgRpfvrnw5LmgLRX0ckOJ8ut/I35CL7awi2ecq6Ca4lL59bEK9hr7nqoHWPA==",
			"dev": true,
			"requires": {
			  "jest-get-type": "^26.0.0",
			  "pretty-format": "^26.0.1"
			},
			"dependencies": {
			  "@jest/types": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/@jest/types/-/types-26.0.1.tgz",
				"integrity": "sha512-IbtjvqI9+eS1qFnOIEL7ggWmT+iK/U+Vde9cGWtYb/b6XgKb3X44ZAe/z9YZzoAAZ/E92m0DqrilF934IGNnQA==",
				"dev": true,
				"requires": {
				  "@types/istanbul-lib-coverage": "^2.0.0",
				  "@types/istanbul-reports": "^1.1.1",
				  "@types/yargs": "^15.0.0",
				  "chalk": "^4.0.0"
				}
			  },
			  "ansi-regex": {
				"version": "5.0.0",
				"resolved": "https://registry.npmjs.org/ansi-regex/-/ansi-regex-5.0.0.tgz",
				"integrity": "sha512-bY6fj56OUQ0hU1KjFNDQuJFezqKdrAyFdIevADiqrWHwSlbmBNMHp5ak2f40Pm8JTFyM2mqxkG6ngkHO11f/lg==",
				"dev": true
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-4.0.0.tgz",
				"integrity": "sha512-N9oWFcegS0sFr9oh1oz2d7Npos6vNoWW9HvtCg5N1KRFpUhaAhvTv5Y58g880fZaEYSNm3qDz8SU1UrGvp+n7A==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "jest-get-type": {
				"version": "26.0.0",
				"resolved": "https://registry.npmjs.org/jest-get-type/-/jest-get-type-26.0.0.tgz",
				"integrity": "sha512-zRc1OAPnnws1EVfykXOj19zo2EMw5Hi6HLbFCSjpuJiXtOWAYIjNsHVSbpQ8bDX7L5BGYGI8m+HmKdjHYFF0kg==",
				"dev": true
			  },
			  "pretty-format": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/pretty-format/-/pretty-format-26.0.1.tgz",
				"integrity": "sha512-SWxz6MbupT3ZSlL0Po4WF/KujhQaVehijR2blyRDCzk9e45EaYMVhMBn49fnRuHxtkSpXTes1GxNpVmH86Bxfw==",
				"dev": true,
				"requires": {
				  "@jest/types": "^26.0.1",
				  "ansi-regex": "^5.0.0",
				  "ansi-styles": "^4.0.0",
				  "react-is": "^16.12.0"
				}
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "jest-matcher-utils": {
			"version": "26.0.1",
			"resolved": "https://registry.npmjs.org/jest-matcher-utils/-/jest-matcher-utils-26.0.1.tgz",
			"integrity": "sha512-PUMlsLth0Azen8Q2WFTwnSkGh2JZ8FYuwijC8NR47vXKpsrKmA1wWvgcj1CquuVfcYiDEdj985u5Wmg7COEARw==",
			"dev": true,
			"requires": {
			  "chalk": "^4.0.0",
			  "jest-diff": "^26.0.1",
			  "jest-get-type": "^26.0.0",
			  "pretty-format": "^26.0.1"
			},
			"dependencies": {
			  "@jest/types": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/@jest/types/-/types-26.0.1.tgz",
				"integrity": "sha512-IbtjvqI9+eS1qFnOIEL7ggWmT+iK/U+Vde9cGWtYb/b6XgKb3X44ZAe/z9YZzoAAZ/E92m0DqrilF934IGNnQA==",
				"dev": true,
				"requires": {
				  "@types/istanbul-lib-coverage": "^2.0.0",
				  "@types/istanbul-reports": "^1.1.1",
				  "@types/yargs": "^15.0.0",
				  "chalk": "^4.0.0"
				}
			  },
			  "ansi-regex": {
				"version": "5.0.0",
				"resolved": "https://registry.npmjs.org/ansi-regex/-/ansi-regex-5.0.0.tgz",
				"integrity": "sha512-bY6fj56OUQ0hU1KjFNDQuJFezqKdrAyFdIevADiqrWHwSlbmBNMHp5ak2f40Pm8JTFyM2mqxkG6ngkHO11f/lg==",
				"dev": true
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-4.0.0.tgz",
				"integrity": "sha512-N9oWFcegS0sFr9oh1oz2d7Npos6vNoWW9HvtCg5N1KRFpUhaAhvTv5Y58g880fZaEYSNm3qDz8SU1UrGvp+n7A==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "diff-sequences": {
				"version": "26.0.0",
				"resolved": "https://registry.npmjs.org/diff-sequences/-/diff-sequences-26.0.0.tgz",
				"integrity": "sha512-JC/eHYEC3aSS0vZGjuoc4vHA0yAQTzhQQldXMeMF+JlxLGJlCO38Gma82NV9gk1jGFz8mDzUMeaKXvjRRdJ2dg==",
				"dev": true
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "jest-diff": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/jest-diff/-/jest-diff-26.0.1.tgz",
				"integrity": "sha512-odTcHyl5X+U+QsczJmOjWw5tPvww+y9Yim5xzqxVl/R1j4z71+fHW4g8qu1ugMmKdFdxw+AtQgs5mupPnzcIBQ==",
				"dev": true,
				"requires": {
				  "chalk": "^4.0.0",
				  "diff-sequences": "^26.0.0",
				  "jest-get-type": "^26.0.0",
				  "pretty-format": "^26.0.1"
				}
			  },
			  "jest-get-type": {
				"version": "26.0.0",
				"resolved": "https://registry.npmjs.org/jest-get-type/-/jest-get-type-26.0.0.tgz",
				"integrity": "sha512-zRc1OAPnnws1EVfykXOj19zo2EMw5Hi6HLbFCSjpuJiXtOWAYIjNsHVSbpQ8bDX7L5BGYGI8m+HmKdjHYFF0kg==",
				"dev": true
			  },
			  "pretty-format": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/pretty-format/-/pretty-format-26.0.1.tgz",
				"integrity": "sha512-SWxz6MbupT3ZSlL0Po4WF/KujhQaVehijR2blyRDCzk9e45EaYMVhMBn49fnRuHxtkSpXTes1GxNpVmH86Bxfw==",
				"dev": true,
				"requires": {
				  "@jest/types": "^26.0.1",
				  "ansi-regex": "^5.0.0",
				  "ansi-styles": "^4.0.0",
				  "react-is": "^16.12.0"
				}
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "jest-message-util": {
			"version": "26.0.1",
			"resolved": "https://registry.npmjs.org/jest-message-util/-/jest-message-util-26.0.1.tgz",
			"integrity": "sha512-CbK8uQREZ8umUfo8+zgIfEt+W7HAHjQCoRaNs4WxKGhAYBGwEyvxuK81FXa7VeB9pwDEXeeKOB2qcsNVCAvB7Q==",
			"dev": true,
			"requires": {
			  "@babel/code-frame": "^7.0.0",
			  "@jest/types": "^26.0.1",
			  "@types/stack-utils": "^1.0.1",
			  "chalk": "^4.0.0",
			  "graceful-fs": "^4.2.4",
			  "micromatch": "^4.0.2",
			  "slash": "^3.0.0",
			  "stack-utils": "^2.0.2"
			},
			"dependencies": {
			  "@jest/types": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/@jest/types/-/types-26.0.1.tgz",
				"integrity": "sha512-IbtjvqI9+eS1qFnOIEL7ggWmT+iK/U+Vde9cGWtYb/b6XgKb3X44ZAe/z9YZzoAAZ/E92m0DqrilF934IGNnQA==",
				"dev": true,
				"requires": {
				  "@types/istanbul-lib-coverage": "^2.0.0",
				  "@types/istanbul-reports": "^1.1.1",
				  "@types/yargs": "^15.0.0",
				  "chalk": "^4.0.0"
				}
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-4.0.0.tgz",
				"integrity": "sha512-N9oWFcegS0sFr9oh1oz2d7Npos6vNoWW9HvtCg5N1KRFpUhaAhvTv5Y58g880fZaEYSNm3qDz8SU1UrGvp+n7A==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "graceful-fs": {
				"version": "4.2.4",
				"resolved": "https://registry.npmjs.org/graceful-fs/-/graceful-fs-4.2.4.tgz",
				"integrity": "sha512-WjKPNJF79dtJAVniUlGGWHYGz2jWxT6VhN/4m1NdkbZ2nOsEF+cI1Edgql5zCRhs/VsQYRvrXctxktVXZUkixw==",
				"dev": true
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "micromatch": {
				"version": "4.0.2",
				"resolved": "https://registry.npmjs.org/micromatch/-/micromatch-4.0.2.tgz",
				"integrity": "sha512-y7FpHSbMUMoyPbYUSzO6PaZ6FyRnQOpHuKwbo1G+Knck95XVU4QAiKdGEnj5wwoS7PlOgthX/09u5iFJ+aYf5Q==",
				"dev": true,
				"requires": {
				  "braces": "^3.0.1",
				  "picomatch": "^2.0.5"
				}
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "jest-mock": {
			"version": "26.0.1",
			"resolved": "https://registry.npmjs.org/jest-mock/-/jest-mock-26.0.1.tgz",
			"integrity": "sha512-MpYTBqycuPYSY6xKJognV7Ja46/TeRbAZept987Zp+tuJvMN0YBWyyhG9mXyYQaU3SBI0TUlSaO5L3p49agw7Q==",
			"dev": true,
			"requires": {
			  "@jest/types": "^26.0.1"
			},
			"dependencies": {
			  "@jest/types": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/@jest/types/-/types-26.0.1.tgz",
				"integrity": "sha512-IbtjvqI9+eS1qFnOIEL7ggWmT+iK/U+Vde9cGWtYb/b6XgKb3X44ZAe/z9YZzoAAZ/E92m0DqrilF934IGNnQA==",
				"dev": true,
				"requires": {
				  "@types/istanbul-lib-coverage": "^2.0.0",
				  "@types/istanbul-reports": "^1.1.1",
				  "@types/yargs": "^15.0.0",
				  "chalk": "^4.0.0"
				}
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-4.0.0.tgz",
				"integrity": "sha512-N9oWFcegS0sFr9oh1oz2d7Npos6vNoWW9HvtCg5N1KRFpUhaAhvTv5Y58g880fZaEYSNm3qDz8SU1UrGvp+n7A==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "jest-pnp-resolver": {
			"version": "1.2.1",
			"resolved": "https://registry.npmjs.org/jest-pnp-resolver/-/jest-pnp-resolver-1.2.1.tgz",
			"integrity": "sha512-pgFw2tm54fzgYvc/OHrnysABEObZCUNFnhjoRjaVOCN8NYc032/gVjPaHD4Aq6ApkSieWtfKAFQtmDKAmhupnQ==",
			"dev": true
		  },
		  "jest-regex-util": {
			"version": "26.0.0",
			"resolved": "https://registry.npmjs.org/jest-regex-util/-/jest-regex-util-26.0.0.tgz",
			"integrity": "sha512-Gv3ZIs/nA48/Zvjrl34bf+oD76JHiGDUxNOVgUjh3j890sblXryjY4rss71fPtD/njchl6PSE2hIhvyWa1eT0A==",
			"dev": true
		  },
		  "jest-resolve": {
			"version": "26.0.1",
			"resolved": "https://registry.npmjs.org/jest-resolve/-/jest-resolve-26.0.1.tgz",
			"integrity": "sha512-6jWxk0IKZkPIVTvq6s72RH735P8f9eCJW3IM5CX/SJFeKq1p2cZx0U49wf/SdMlhaB/anann5J2nCJj6HrbezQ==",
			"dev": true,
			"requires": {
			  "@jest/types": "^26.0.1",
			  "chalk": "^4.0.0",
			  "graceful-fs": "^4.2.4",
			  "jest-pnp-resolver": "^1.2.1",
			  "jest-util": "^26.0.1",
			  "read-pkg-up": "^7.0.1",
			  "resolve": "^1.17.0",
			  "slash": "^3.0.0"
			},
			"dependencies": {
			  "@jest/types": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/@jest/types/-/types-26.0.1.tgz",
				"integrity": "sha512-IbtjvqI9+eS1qFnOIEL7ggWmT+iK/U+Vde9cGWtYb/b6XgKb3X44ZAe/z9YZzoAAZ/E92m0DqrilF934IGNnQA==",
				"dev": true,
				"requires": {
				  "@types/istanbul-lib-coverage": "^2.0.0",
				  "@types/istanbul-reports": "^1.1.1",
				  "@types/yargs": "^15.0.0",
				  "chalk": "^4.0.0"
				}
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-4.0.0.tgz",
				"integrity": "sha512-N9oWFcegS0sFr9oh1oz2d7Npos6vNoWW9HvtCg5N1KRFpUhaAhvTv5Y58g880fZaEYSNm3qDz8SU1UrGvp+n7A==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "find-up": {
				"version": "4.1.0",
				"resolved": "https://registry.npmjs.org/find-up/-/find-up-4.1.0.tgz",
				"integrity": "sha512-PpOwAdQ/YlXQ2vj8a3h8IipDuYRi3wceVQQGYWxNINccq40Anw7BlsEXCMbt1Zt+OLA6Fq9suIpIWD0OsnISlw==",
				"dev": true,
				"requires": {
				  "locate-path": "^5.0.0",
				  "path-exists": "^4.0.0"
				}
			  },
			  "graceful-fs": {
				"version": "4.2.4",
				"resolved": "https://registry.npmjs.org/graceful-fs/-/graceful-fs-4.2.4.tgz",
				"integrity": "sha512-WjKPNJF79dtJAVniUlGGWHYGz2jWxT6VhN/4m1NdkbZ2nOsEF+cI1Edgql5zCRhs/VsQYRvrXctxktVXZUkixw==",
				"dev": true
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "locate-path": {
				"version": "5.0.0",
				"resolved": "https://registry.npmjs.org/locate-path/-/locate-path-5.0.0.tgz",
				"integrity": "sha512-t7hw9pI+WvuwNJXwk5zVHpyhIqzg2qTlklJOf0mVxGSbe3Fp2VieZcduNYjaLDoy6p9uGpQEGWG87WpMKlNq8g==",
				"dev": true,
				"requires": {
				  "p-locate": "^4.1.0"
				}
			  },
			  "p-locate": {
				"version": "4.1.0",
				"resolved": "https://registry.npmjs.org/p-locate/-/p-locate-4.1.0.tgz",
				"integrity": "sha512-R79ZZ/0wAxKGu3oYMlz8jy/kbhsNrS7SKZ7PxEHBgJ5+F2mtFW2fK2cOtBh1cHYkQsbzFV7I+EoRKe6Yt0oK7A==",
				"dev": true,
				"requires": {
				  "p-limit": "^2.2.0"
				}
			  },
			  "parse-json": {
				"version": "5.0.0",
				"resolved": "https://registry.npmjs.org/parse-json/-/parse-json-5.0.0.tgz",
				"integrity": "sha512-OOY5b7PAEFV0E2Fir1KOkxchnZNCdowAJgQ5NuxjpBKTRP3pQhwkrkxqQjeoKJ+fO7bCpmIZaogI4eZGDMEGOw==",
				"dev": true,
				"requires": {
				  "@babel/code-frame": "^7.0.0",
				  "error-ex": "^1.3.1",
				  "json-parse-better-errors": "^1.0.1",
				  "lines-and-columns": "^1.1.6"
				}
			  },
			  "path-exists": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/path-exists/-/path-exists-4.0.0.tgz",
				"integrity": "sha512-ak9Qy5Q7jYb2Wwcey5Fpvg2KoAc/ZIhLSLOSBmRmygPsGwkVVt0fZa0qrtMz+m6tJTAHfZQ8FnmB4MG4LWy7/w==",
				"dev": true
			  },
			  "read-pkg": {
				"version": "5.2.0",
				"resolved": "https://registry.npmjs.org/read-pkg/-/read-pkg-5.2.0.tgz",
				"integrity": "sha512-Ug69mNOpfvKDAc2Q8DRpMjjzdtrnv9HcSMX+4VsZxD1aZ6ZzrIE7rlzXBtWTyhULSMKg076AW6WR5iZpD0JiOg==",
				"dev": true,
				"requires": {
				  "@types/normalize-package-data": "^2.4.0",
				  "normalize-package-data": "^2.5.0",
				  "parse-json": "^5.0.0",
				  "type-fest": "^0.6.0"
				},
				"dependencies": {
				  "type-fest": {
					"version": "0.6.0",
					"resolved": "https://registry.npmjs.org/type-fest/-/type-fest-0.6.0.tgz",
					"integrity": "sha512-q+MB8nYR1KDLrgr4G5yemftpMC7/QLqVndBmEEdqzmNj5dcFOO4Oo8qlwZE3ULT3+Zim1F8Kq4cBnikNhlCMlg==",
					"dev": true
				  }
				}
			  },
			  "read-pkg-up": {
				"version": "7.0.1",
				"resolved": "https://registry.npmjs.org/read-pkg-up/-/read-pkg-up-7.0.1.tgz",
				"integrity": "sha512-zK0TB7Xd6JpCLmlLmufqykGE+/TlOePD6qKClNW7hHDKFh/J7/7gCWGR7joEQEW1bKq3a3yUZSObOoWLFQ4ohg==",
				"dev": true,
				"requires": {
				  "find-up": "^4.1.0",
				  "read-pkg": "^5.2.0",
				  "type-fest": "^0.8.1"
				}
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  },
			  "type-fest": {
				"version": "0.8.1",
				"resolved": "https://registry.npmjs.org/type-fest/-/type-fest-0.8.1.tgz",
				"integrity": "sha512-4dbzIzqvjtgiM5rw1k5rEHtBANKmdudhGyBEajN01fEyhaAIhsoKNy6y7+IN93IfpFtwY9iqi7kD+xwKhQsNJA==",
				"dev": true
			  }
			}
		  },
		  "jest-resolve-dependencies": {
			"version": "26.0.1",
			"resolved": "https://registry.npmjs.org/jest-resolve-dependencies/-/jest-resolve-dependencies-26.0.1.tgz",
			"integrity": "sha512-9d5/RS/ft0vB/qy7jct/qAhzJsr6fRQJyGAFigK3XD4hf9kIbEH5gks4t4Z7kyMRhowU6HWm/o8ILqhaHdSqLw==",
			"dev": true,
			"requires": {
			  "@jest/types": "^26.0.1",
			  "jest-regex-util": "^26.0.0",
			  "jest-snapshot": "^26.0.1"
			},
			"dependencies": {
			  "@jest/types": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/@jest/types/-/types-26.0.1.tgz",
				"integrity": "sha512-IbtjvqI9+eS1qFnOIEL7ggWmT+iK/U+Vde9cGWtYb/b6XgKb3X44ZAe/z9YZzoAAZ/E92m0DqrilF934IGNnQA==",
				"dev": true,
				"requires": {
				  "@types/istanbul-lib-coverage": "^2.0.0",
				  "@types/istanbul-reports": "^1.1.1",
				  "@types/yargs": "^15.0.0",
				  "chalk": "^4.0.0"
				}
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-4.0.0.tgz",
				"integrity": "sha512-N9oWFcegS0sFr9oh1oz2d7Npos6vNoWW9HvtCg5N1KRFpUhaAhvTv5Y58g880fZaEYSNm3qDz8SU1UrGvp+n7A==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "jest-runner": {
			"version": "26.0.1",
			"resolved": "https://registry.npmjs.org/jest-runner/-/jest-runner-26.0.1.tgz",
			"integrity": "sha512-CApm0g81b49Znm4cZekYQK67zY7kkB4umOlI2Dx5CwKAzdgw75EN+ozBHRvxBzwo1ZLYZ07TFxkaPm+1t4d8jA==",
			"dev": true,
			"requires": {
			  "@jest/console": "^26.0.1",
			  "@jest/environment": "^26.0.1",
			  "@jest/test-result": "^26.0.1",
			  "@jest/types": "^26.0.1",
			  "chalk": "^4.0.0",
			  "exit": "^0.1.2",
			  "graceful-fs": "^4.2.4",
			  "jest-config": "^26.0.1",
			  "jest-docblock": "^26.0.0",
			  "jest-haste-map": "^26.0.1",
			  "jest-jasmine2": "^26.0.1",
			  "jest-leak-detector": "^26.0.1",
			  "jest-message-util": "^26.0.1",
			  "jest-resolve": "^26.0.1",
			  "jest-runtime": "^26.0.1",
			  "jest-util": "^26.0.1",
			  "jest-worker": "^26.0.0",
			  "source-map-support": "^0.5.6",
			  "throat": "^5.0.0"
			},
			"dependencies": {
			  "@jest/types": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/@jest/types/-/types-26.0.1.tgz",
				"integrity": "sha512-IbtjvqI9+eS1qFnOIEL7ggWmT+iK/U+Vde9cGWtYb/b6XgKb3X44ZAe/z9YZzoAAZ/E92m0DqrilF934IGNnQA==",
				"dev": true,
				"requires": {
				  "@types/istanbul-lib-coverage": "^2.0.0",
				  "@types/istanbul-reports": "^1.1.1",
				  "@types/yargs": "^15.0.0",
				  "chalk": "^4.0.0"
				}
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-4.0.0.tgz",
				"integrity": "sha512-N9oWFcegS0sFr9oh1oz2d7Npos6vNoWW9HvtCg5N1KRFpUhaAhvTv5Y58g880fZaEYSNm3qDz8SU1UrGvp+n7A==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "graceful-fs": {
				"version": "4.2.4",
				"resolved": "https://registry.npmjs.org/graceful-fs/-/graceful-fs-4.2.4.tgz",
				"integrity": "sha512-WjKPNJF79dtJAVniUlGGWHYGz2jWxT6VhN/4m1NdkbZ2nOsEF+cI1Edgql5zCRhs/VsQYRvrXctxktVXZUkixw==",
				"dev": true
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "jest-runtime": {
			"version": "26.0.1",
			"resolved": "https://registry.npmjs.org/jest-runtime/-/jest-runtime-26.0.1.tgz",
			"integrity": "sha512-Ci2QhYFmANg5qaXWf78T2Pfo6GtmIBn2rRaLnklRyEucmPccmCKvS9JPljcmtVamsdMmkyNkVFb9pBTD6si9Lw==",
			"dev": true,
			"requires": {
			  "@jest/console": "^26.0.1",
			  "@jest/environment": "^26.0.1",
			  "@jest/fake-timers": "^26.0.1",
			  "@jest/globals": "^26.0.1",
			  "@jest/source-map": "^26.0.0",
			  "@jest/test-result": "^26.0.1",
			  "@jest/transform": "^26.0.1",
			  "@jest/types": "^26.0.1",
			  "@types/yargs": "^15.0.0",
			  "chalk": "^4.0.0",
			  "collect-v8-coverage": "^1.0.0",
			  "exit": "^0.1.2",
			  "glob": "^7.1.3",
			  "graceful-fs": "^4.2.4",
			  "jest-config": "^26.0.1",
			  "jest-haste-map": "^26.0.1",
			  "jest-message-util": "^26.0.1",
			  "jest-mock": "^26.0.1",
			  "jest-regex-util": "^26.0.0",
			  "jest-resolve": "^26.0.1",
			  "jest-snapshot": "^26.0.1",
			  "jest-util": "^26.0.1",
			  "jest-validate": "^26.0.1",
			  "slash": "^3.0.0",
			  "strip-bom": "^4.0.0",
			  "yargs": "^15.3.1"
			},
			"dependencies": {
			  "@jest/types": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/@jest/types/-/types-26.0.1.tgz",
				"integrity": "sha512-IbtjvqI9+eS1qFnOIEL7ggWmT+iK/U+Vde9cGWtYb/b6XgKb3X44ZAe/z9YZzoAAZ/E92m0DqrilF934IGNnQA==",
				"dev": true,
				"requires": {
				  "@types/istanbul-lib-coverage": "^2.0.0",
				  "@types/istanbul-reports": "^1.1.1",
				  "@types/yargs": "^15.0.0",
				  "chalk": "^4.0.0"
				}
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-4.0.0.tgz",
				"integrity": "sha512-N9oWFcegS0sFr9oh1oz2d7Npos6vNoWW9HvtCg5N1KRFpUhaAhvTv5Y58g880fZaEYSNm3qDz8SU1UrGvp+n7A==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "graceful-fs": {
				"version": "4.2.4",
				"resolved": "https://registry.npmjs.org/graceful-fs/-/graceful-fs-4.2.4.tgz",
				"integrity": "sha512-WjKPNJF79dtJAVniUlGGWHYGz2jWxT6VhN/4m1NdkbZ2nOsEF+cI1Edgql5zCRhs/VsQYRvrXctxktVXZUkixw==",
				"dev": true
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "strip-bom": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/strip-bom/-/strip-bom-4.0.0.tgz",
				"integrity": "sha512-3xurFv5tEgii33Zi8Jtp55wEIILR9eh34FAW00PZf+JnSsTmV/ioewSgQl97JHvgjoRGwPShsWm+IdrxB35d0w==",
				"dev": true
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "jest-serializer": {
			"version": "26.0.0",
			"resolved": "https://registry.npmjs.org/jest-serializer/-/jest-serializer-26.0.0.tgz",
			"integrity": "sha512-sQGXLdEGWFAE4wIJ2ZaIDb+ikETlUirEOBsLXdoBbeLhTHkZUJwgk3+M8eyFizhM6le43PDCCKPA1hzkSDo4cQ==",
			"dev": true,
			"requires": {
			  "graceful-fs": "^4.2.4"
			},
			"dependencies": {
			  "graceful-fs": {
				"version": "4.2.4",
				"resolved": "https://registry.npmjs.org/graceful-fs/-/graceful-fs-4.2.4.tgz",
				"integrity": "sha512-WjKPNJF79dtJAVniUlGGWHYGz2jWxT6VhN/4m1NdkbZ2nOsEF+cI1Edgql5zCRhs/VsQYRvrXctxktVXZUkixw==",
				"dev": true
			  }
			}
		  },
		  "jest-snapshot": {
			"version": "26.0.1",
			"resolved": "https://registry.npmjs.org/jest-snapshot/-/jest-snapshot-26.0.1.tgz",
			"integrity": "sha512-jxd+cF7+LL+a80qh6TAnTLUZHyQoWwEHSUFJjkw35u3Gx+BZUNuXhYvDqHXr62UQPnWo2P6fvQlLjsU93UKyxA==",
			"dev": true,
			"requires": {
			  "@babel/types": "^7.0.0",
			  "@jest/types": "^26.0.1",
			  "@types/prettier": "^2.0.0",
			  "chalk": "^4.0.0",
			  "expect": "^26.0.1",
			  "graceful-fs": "^4.2.4",
			  "jest-diff": "^26.0.1",
			  "jest-get-type": "^26.0.0",
			  "jest-matcher-utils": "^26.0.1",
			  "jest-message-util": "^26.0.1",
			  "jest-resolve": "^26.0.1",
			  "make-dir": "^3.0.0",
			  "natural-compare": "^1.4.0",
			  "pretty-format": "^26.0.1",
			  "semver": "^7.3.2"
			},
			"dependencies": {
			  "@jest/types": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/@jest/types/-/types-26.0.1.tgz",
				"integrity": "sha512-IbtjvqI9+eS1qFnOIEL7ggWmT+iK/U+Vde9cGWtYb/b6XgKb3X44ZAe/z9YZzoAAZ/E92m0DqrilF934IGNnQA==",
				"dev": true,
				"requires": {
				  "@types/istanbul-lib-coverage": "^2.0.0",
				  "@types/istanbul-reports": "^1.1.1",
				  "@types/yargs": "^15.0.0",
				  "chalk": "^4.0.0"
				}
			  },
			  "ansi-regex": {
				"version": "5.0.0",
				"resolved": "https://registry.npmjs.org/ansi-regex/-/ansi-regex-5.0.0.tgz",
				"integrity": "sha512-bY6fj56OUQ0hU1KjFNDQuJFezqKdrAyFdIevADiqrWHwSlbmBNMHp5ak2f40Pm8JTFyM2mqxkG6ngkHO11f/lg==",
				"dev": true
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-4.0.0.tgz",
				"integrity": "sha512-N9oWFcegS0sFr9oh1oz2d7Npos6vNoWW9HvtCg5N1KRFpUhaAhvTv5Y58g880fZaEYSNm3qDz8SU1UrGvp+n7A==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "diff-sequences": {
				"version": "26.0.0",
				"resolved": "https://registry.npmjs.org/diff-sequences/-/diff-sequences-26.0.0.tgz",
				"integrity": "sha512-JC/eHYEC3aSS0vZGjuoc4vHA0yAQTzhQQldXMeMF+JlxLGJlCO38Gma82NV9gk1jGFz8mDzUMeaKXvjRRdJ2dg==",
				"dev": true
			  },
			  "graceful-fs": {
				"version": "4.2.4",
				"resolved": "https://registry.npmjs.org/graceful-fs/-/graceful-fs-4.2.4.tgz",
				"integrity": "sha512-WjKPNJF79dtJAVniUlGGWHYGz2jWxT6VhN/4m1NdkbZ2nOsEF+cI1Edgql5zCRhs/VsQYRvrXctxktVXZUkixw==",
				"dev": true
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "jest-diff": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/jest-diff/-/jest-diff-26.0.1.tgz",
				"integrity": "sha512-odTcHyl5X+U+QsczJmOjWw5tPvww+y9Yim5xzqxVl/R1j4z71+fHW4g8qu1ugMmKdFdxw+AtQgs5mupPnzcIBQ==",
				"dev": true,
				"requires": {
				  "chalk": "^4.0.0",
				  "diff-sequences": "^26.0.0",
				  "jest-get-type": "^26.0.0",
				  "pretty-format": "^26.0.1"
				}
			  },
			  "jest-get-type": {
				"version": "26.0.0",
				"resolved": "https://registry.npmjs.org/jest-get-type/-/jest-get-type-26.0.0.tgz",
				"integrity": "sha512-zRc1OAPnnws1EVfykXOj19zo2EMw5Hi6HLbFCSjpuJiXtOWAYIjNsHVSbpQ8bDX7L5BGYGI8m+HmKdjHYFF0kg==",
				"dev": true
			  },
			  "make-dir": {
				"version": "3.1.0",
				"resolved": "https://registry.npmjs.org/make-dir/-/make-dir-3.1.0.tgz",
				"integrity": "sha512-g3FeP20LNwhALb/6Cz6Dd4F2ngze0jz7tbzrD2wAV+o9FeNHe4rL+yK2md0J/fiSf1sa1ADhXqi5+oVwOM/eGw==",
				"dev": true,
				"requires": {
				  "semver": "^6.0.0"
				},
				"dependencies": {
				  "semver": {
					"version": "6.3.0",
					"resolved": "https://registry.npmjs.org/semver/-/semver-6.3.0.tgz",
					"integrity": "sha512-b39TBaTSfV6yBrapU89p5fKekE2m/NwnDocOVruQFS1/veMgdzuPcnOM34M6CwxW8jH/lxEa5rBoDeUwu5HHTw==",
					"dev": true
				  }
				}
			  },
			  "pretty-format": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/pretty-format/-/pretty-format-26.0.1.tgz",
				"integrity": "sha512-SWxz6MbupT3ZSlL0Po4WF/KujhQaVehijR2blyRDCzk9e45EaYMVhMBn49fnRuHxtkSpXTes1GxNpVmH86Bxfw==",
				"dev": true,
				"requires": {
				  "@jest/types": "^26.0.1",
				  "ansi-regex": "^5.0.0",
				  "ansi-styles": "^4.0.0",
				  "react-is": "^16.12.0"
				}
			  },
			  "semver": {
				"version": "7.3.2",
				"resolved": "https://registry.npmjs.org/semver/-/semver-7.3.2.tgz",
				"integrity": "sha512-OrOb32TeeambH6UrhtShmF7CRDqhL6/5XpPNp2DuRH6+9QLw/orhp72j87v8Qa1ScDkvrrBNpZcDejAirJmfXQ==",
				"dev": true
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "jest-util": {
			"version": "26.0.1",
			"resolved": "https://registry.npmjs.org/jest-util/-/jest-util-26.0.1.tgz",
			"integrity": "sha512-byQ3n7ad1BO/WyFkYvlWQHTsomB6GIewBh8tlGtusiylAlaxQ1UpS0XYH0ngOyhZuHVLN79Qvl6/pMiDMSSG1g==",
			"dev": true,
			"requires": {
			  "@jest/types": "^26.0.1",
			  "chalk": "^4.0.0",
			  "graceful-fs": "^4.2.4",
			  "is-ci": "^2.0.0",
			  "make-dir": "^3.0.0"
			},
			"dependencies": {
			  "@jest/types": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/@jest/types/-/types-26.0.1.tgz",
				"integrity": "sha512-IbtjvqI9+eS1qFnOIEL7ggWmT+iK/U+Vde9cGWtYb/b6XgKb3X44ZAe/z9YZzoAAZ/E92m0DqrilF934IGNnQA==",
				"dev": true,
				"requires": {
				  "@types/istanbul-lib-coverage": "^2.0.0",
				  "@types/istanbul-reports": "^1.1.1",
				  "@types/yargs": "^15.0.0",
				  "chalk": "^4.0.0"
				}
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-4.0.0.tgz",
				"integrity": "sha512-N9oWFcegS0sFr9oh1oz2d7Npos6vNoWW9HvtCg5N1KRFpUhaAhvTv5Y58g880fZaEYSNm3qDz8SU1UrGvp+n7A==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "graceful-fs": {
				"version": "4.2.4",
				"resolved": "https://registry.npmjs.org/graceful-fs/-/graceful-fs-4.2.4.tgz",
				"integrity": "sha512-WjKPNJF79dtJAVniUlGGWHYGz2jWxT6VhN/4m1NdkbZ2nOsEF+cI1Edgql5zCRhs/VsQYRvrXctxktVXZUkixw==",
				"dev": true
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "make-dir": {
				"version": "3.1.0",
				"resolved": "https://registry.npmjs.org/make-dir/-/make-dir-3.1.0.tgz",
				"integrity": "sha512-g3FeP20LNwhALb/6Cz6Dd4F2ngze0jz7tbzrD2wAV+o9FeNHe4rL+yK2md0J/fiSf1sa1ADhXqi5+oVwOM/eGw==",
				"dev": true,
				"requires": {
				  "semver": "^6.0.0"
				}
			  },
			  "semver": {
				"version": "6.3.0",
				"resolved": "https://registry.npmjs.org/semver/-/semver-6.3.0.tgz",
				"integrity": "sha512-b39TBaTSfV6yBrapU89p5fKekE2m/NwnDocOVruQFS1/veMgdzuPcnOM34M6CwxW8jH/lxEa5rBoDeUwu5HHTw==",
				"dev": true
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "jest-validate": {
			"version": "26.0.1",
			"resolved": "https://registry.npmjs.org/jest-validate/-/jest-validate-26.0.1.tgz",
			"integrity": "sha512-u0xRc+rbmov/VqXnX3DlkxD74rHI/CfS5xaV2VpeaVySjbb1JioNVOyly5b56q2l9ZKe7bVG5qWmjfctkQb0bA==",
			"dev": true,
			"requires": {
			  "@jest/types": "^26.0.1",
			  "camelcase": "^6.0.0",
			  "chalk": "^4.0.0",
			  "jest-get-type": "^26.0.0",
			  "leven": "^3.1.0",
			  "pretty-format": "^26.0.1"
			},
			"dependencies": {
			  "@jest/types": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/@jest/types/-/types-26.0.1.tgz",
				"integrity": "sha512-IbtjvqI9+eS1qFnOIEL7ggWmT+iK/U+Vde9cGWtYb/b6XgKb3X44ZAe/z9YZzoAAZ/E92m0DqrilF934IGNnQA==",
				"dev": true,
				"requires": {
				  "@types/istanbul-lib-coverage": "^2.0.0",
				  "@types/istanbul-reports": "^1.1.1",
				  "@types/yargs": "^15.0.0",
				  "chalk": "^4.0.0"
				}
			  },
			  "ansi-regex": {
				"version": "5.0.0",
				"resolved": "https://registry.npmjs.org/ansi-regex/-/ansi-regex-5.0.0.tgz",
				"integrity": "sha512-bY6fj56OUQ0hU1KjFNDQuJFezqKdrAyFdIevADiqrWHwSlbmBNMHp5ak2f40Pm8JTFyM2mqxkG6ngkHO11f/lg==",
				"dev": true
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "camelcase": {
				"version": "6.0.0",
				"resolved": "https://registry.npmjs.org/camelcase/-/camelcase-6.0.0.tgz",
				"integrity": "sha512-8KMDF1Vz2gzOq54ONPJS65IvTUaB1cHJ2DMM7MbPmLZljDH1qpzzLsWdiN9pHh6qvkRVDTi/07+eNGch/oLU4w==",
				"dev": true
			  },
			  "chalk": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-4.0.0.tgz",
				"integrity": "sha512-N9oWFcegS0sFr9oh1oz2d7Npos6vNoWW9HvtCg5N1KRFpUhaAhvTv5Y58g880fZaEYSNm3qDz8SU1UrGvp+n7A==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "jest-get-type": {
				"version": "26.0.0",
				"resolved": "https://registry.npmjs.org/jest-get-type/-/jest-get-type-26.0.0.tgz",
				"integrity": "sha512-zRc1OAPnnws1EVfykXOj19zo2EMw5Hi6HLbFCSjpuJiXtOWAYIjNsHVSbpQ8bDX7L5BGYGI8m+HmKdjHYFF0kg==",
				"dev": true
			  },
			  "pretty-format": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/pretty-format/-/pretty-format-26.0.1.tgz",
				"integrity": "sha512-SWxz6MbupT3ZSlL0Po4WF/KujhQaVehijR2blyRDCzk9e45EaYMVhMBn49fnRuHxtkSpXTes1GxNpVmH86Bxfw==",
				"dev": true,
				"requires": {
				  "@jest/types": "^26.0.1",
				  "ansi-regex": "^5.0.0",
				  "ansi-styles": "^4.0.0",
				  "react-is": "^16.12.0"
				}
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "jest-watcher": {
			"version": "26.0.1",
			"resolved": "https://registry.npmjs.org/jest-watcher/-/jest-watcher-26.0.1.tgz",
			"integrity": "sha512-pdZPydsS8475f89kGswaNsN3rhP6lnC3/QDCppP7bg1L9JQz7oU9Mb/5xPETk1RHDCWeqmVC47M4K5RR7ejxFw==",
			"dev": true,
			"requires": {
			  "@jest/test-result": "^26.0.1",
			  "@jest/types": "^26.0.1",
			  "ansi-escapes": "^4.2.1",
			  "chalk": "^4.0.0",
			  "jest-util": "^26.0.1",
			  "string-length": "^4.0.1"
			},
			"dependencies": {
			  "@jest/types": {
				"version": "26.0.1",
				"resolved": "https://registry.npmjs.org/@jest/types/-/types-26.0.1.tgz",
				"integrity": "sha512-IbtjvqI9+eS1qFnOIEL7ggWmT+iK/U+Vde9cGWtYb/b6XgKb3X44ZAe/z9YZzoAAZ/E92m0DqrilF934IGNnQA==",
				"dev": true,
				"requires": {
				  "@types/istanbul-lib-coverage": "^2.0.0",
				  "@types/istanbul-reports": "^1.1.1",
				  "@types/yargs": "^15.0.0",
				  "chalk": "^4.0.0"
				}
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-4.0.0.tgz",
				"integrity": "sha512-N9oWFcegS0sFr9oh1oz2d7Npos6vNoWW9HvtCg5N1KRFpUhaAhvTv5Y58g880fZaEYSNm3qDz8SU1UrGvp+n7A==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "jest-worker": {
			"version": "26.0.0",
			"resolved": "https://registry.npmjs.org/jest-worker/-/jest-worker-26.0.0.tgz",
			"integrity": "sha512-pPaYa2+JnwmiZjK9x7p9BoZht+47ecFCDFA/CJxspHzeDvQcfVBLWzCiWyo+EGrSiQMWZtCFo9iSvMZnAAo8vw==",
			"dev": true,
			"requires": {
			  "merge-stream": "^2.0.0",
			  "supports-color": "^7.0.0"
			},
			"dependencies": {
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "js-yaml": {
			"version": "3.13.1",
			"resolved": "https://registry.npmjs.org/js-yaml/-/js-yaml-3.13.1.tgz",
			"integrity": "sha512-YfbcO7jXDdyj0DGxYVSlSeQNHbD7XPWvrVWeVUujrQEoZzWJIRrCPoyk6kL6IAjAG2IolMK4T0hNUe0HOUs5Jw==",
			"dev": true,
			"requires": {
			  "argparse": "^1.0.7",
			  "esprima": "^4.0.0"
			}
		  },
		  "jsbn": {
			"version": "0.1.1",
			"resolved": "https://registry.npmjs.org/jsbn/-/jsbn-0.1.1.tgz",
			"integrity": "sha1-peZUwuWi3rXyAdls77yoDA7y9RM=",
			"dev": true
		  },
		  "jsdom": {
			"version": "16.2.2",
			"resolved": "https://registry.npmjs.org/jsdom/-/jsdom-16.2.2.tgz",
			"integrity": "sha512-pDFQbcYtKBHxRaP55zGXCJWgFHkDAYbKcsXEK/3Icu9nKYZkutUXfLBwbD+09XDutkYSHcgfQLZ0qvpAAm9mvg==",
			"dev": true,
			"requires": {
			  "abab": "^2.0.3",
			  "acorn": "^7.1.1",
			  "acorn-globals": "^6.0.0",
			  "cssom": "^0.4.4",
			  "cssstyle": "^2.2.0",
			  "data-urls": "^2.0.0",
			  "decimal.js": "^10.2.0",
			  "domexception": "^2.0.1",
			  "escodegen": "^1.14.1",
			  "html-encoding-sniffer": "^2.0.1",
			  "is-potential-custom-element-name": "^1.0.0",
			  "nwsapi": "^2.2.0",
			  "parse5": "5.1.1",
			  "request": "^2.88.2",
			  "request-promise-native": "^1.0.8",
			  "saxes": "^5.0.0",
			  "symbol-tree": "^3.2.4",
			  "tough-cookie": "^3.0.1",
			  "w3c-hr-time": "^1.0.2",
			  "w3c-xmlserializer": "^2.0.0",
			  "webidl-conversions": "^6.0.0",
			  "whatwg-encoding": "^1.0.5",
			  "whatwg-mimetype": "^2.3.0",
			  "whatwg-url": "^8.0.0",
			  "ws": "^7.2.3",
			  "xml-name-validator": "^3.0.0"
			},
			"dependencies": {
			  "acorn": {
				"version": "7.2.0",
				"resolved": "https://registry.npmjs.org/acorn/-/acorn-7.2.0.tgz",
				"integrity": "sha512-apwXVmYVpQ34m/i71vrApRrRKCWQnZZF1+npOD0WV5xZFfwWOmKGQ2RWlfdy9vWITsenisM8M0Qeq8agcFHNiQ==",
				"dev": true
			  }
			}
		  },
		  "jsesc": {
			"version": "2.5.2",
			"resolved": "https://registry.npmjs.org/jsesc/-/jsesc-2.5.2.tgz",
			"integrity": "sha512-OYu7XEzjkCQ3C5Ps3QIZsQfNpqoJyZZA99wd9aWd05NCtC5pWOkShK2mkL6HXQR6/Cy2lbNdPlZBpuQHXE63gA==",
			"dev": true
		  },
		  "json-parse-better-errors": {
			"version": "1.0.2",
			"resolved": "https://registry.npmjs.org/json-parse-better-errors/-/json-parse-better-errors-1.0.2.tgz",
			"integrity": "sha512-mrqyZKfX5EhL7hvqcV6WG1yYjnjeuYDzDhhcAAUrq8Po85NBQBJP+ZDUT75qZQ98IkUoBqdkExkukOU7Ts2wrw==",
			"dev": true
		  },
		  "json-schema": {
			"version": "0.2.3",
			"resolved": "https://registry.npmjs.org/json-schema/-/json-schema-0.2.3.tgz",
			"integrity": "sha1-tIDIkuWaLwWVTOcnvT8qTogvnhM=",
			"dev": true
		  },
		  "json-schema-traverse": {
			"version": "0.4.1",
			"resolved": "https://registry.npmjs.org/json-schema-traverse/-/json-schema-traverse-0.4.1.tgz",
			"integrity": "sha512-xbbCH5dCYU5T8LcEhhuh7HJ88HXuW3qsI3Y0zOZFKfZEHcpWiHU/Jxzk629Brsab/mMiHQti9wMP+845RPe3Vg==",
			"dev": true
		  },
		  "json-stable-stringify-without-jsonify": {
			"version": "1.0.1",
			"resolved": "https://registry.npmjs.org/json-stable-stringify-without-jsonify/-/json-stable-stringify-without-jsonify-1.0.1.tgz",
			"integrity": "sha1-nbe1lJatPzz+8wp1FC0tkwrXJlE=",
			"dev": true
		  },
		  "json-stringify-safe": {
			"version": "5.0.1",
			"resolved": "https://registry.npmjs.org/json-stringify-safe/-/json-stringify-safe-5.0.1.tgz",
			"integrity": "sha1-Epai1Y/UXxmg9s4B1lcB4sc1tus=",
			"dev": true
		  },
		  "json5": {
			"version": "1.0.1",
			"resolved": "https://registry.npmjs.org/json5/-/json5-1.0.1.tgz",
			"integrity": "sha512-aKS4WQjPenRxiQsC93MNfjx+nbF4PAdYzmd/1JIj8HYzqfbu86beTuNgXDzPknWk0n0uARlyewZo4s++ES36Ow==",
			"dev": true,
			"requires": {
			  "minimist": "^1.2.0"
			}
		  },
		  "jsonfile": {
			"version": "4.0.0",
			"resolved": "https://registry.npmjs.org/jsonfile/-/jsonfile-4.0.0.tgz",
			"integrity": "sha1-h3Gq4HmbZAdrdmQPygWPnBDjPss=",
			"dev": true,
			"requires": {
			  "graceful-fs": "^4.1.6"
			}
		  },
		  "jsprim": {
			"version": "1.4.1",
			"resolved": "https://registry.npmjs.org/jsprim/-/jsprim-1.4.1.tgz",
			"integrity": "sha1-MT5mvB5cwG5Di8G3SZwuXFastqI=",
			"dev": true,
			"requires": {
			  "assert-plus": "1.0.0",
			  "extsprintf": "1.3.0",
			  "json-schema": "0.2.3",
			  "verror": "1.10.0"
			}
		  },
		  "kind-of": {
			"version": "6.0.3",
			"resolved": "https://registry.npmjs.org/kind-of/-/kind-of-6.0.3.tgz",
			"integrity": "sha512-dcS1ul+9tmeD95T+x28/ehLgd9mENa3LsvDTtzm3vyBEO7RPptvAD+t44WVXaUjTBRcrpFeFlC8WCruUR456hw==",
			"dev": true
		  },
		  "kleur": {
			"version": "3.0.3",
			"resolved": "https://registry.npmjs.org/kleur/-/kleur-3.0.3.tgz",
			"integrity": "sha512-eTIzlVOSUR+JxdDFepEYcBMtZ9Qqdef+rnzWdRZuMbOywu5tO2w2N7rqjoANZ5k9vywhL6Br1VRjUIgTQx4E8w==",
			"dev": true
		  },
		  "leven": {
			"version": "3.1.0",
			"resolved": "https://registry.npmjs.org/leven/-/leven-3.1.0.tgz",
			"integrity": "sha512-qsda+H8jTaUaN/x5vzW2rzc+8Rw4TAQ/4KjB46IwK5VH+IlVeeeje/EoZRpiXvIqjFgK84QffqPztGI3VBLG1A==",
			"dev": true
		  },
		  "levn": {
			"version": "0.4.1",
			"resolved": "https://registry.npmjs.org/levn/-/levn-0.4.1.tgz",
			"integrity": "sha512-+bT2uH4E5LGE7h/n3evcS/sQlJXCpIp6ym8OWJ5eV6+67Dsql/LaaT7qJBAt2rzfoa/5QBGBhxDix1dMt2kQKQ==",
			"dev": true,
			"requires": {
			  "prelude-ls": "^1.2.1",
			  "type-check": "~0.4.0"
			}
		  },
		  "lines-and-columns": {
			"version": "1.1.6",
			"resolved": "https://registry.npmjs.org/lines-and-columns/-/lines-and-columns-1.1.6.tgz",
			"integrity": "sha1-HADHQ7QzzQpOgHWPe2SldEDZ/wA=",
			"dev": true
		  },
		  "load-json-file": {
			"version": "2.0.0",
			"resolved": "https://registry.npmjs.org/load-json-file/-/load-json-file-2.0.0.tgz",
			"integrity": "sha1-eUfkIUmvgNaWy/eXvKq8/h/inKg=",
			"dev": true,
			"requires": {
			  "graceful-fs": "^4.1.2",
			  "parse-json": "^2.2.0",
			  "pify": "^2.0.0",
			  "strip-bom": "^3.0.0"
			},
			"dependencies": {
			  "pify": {
				"version": "2.3.0",
				"resolved": "https://registry.npmjs.org/pify/-/pify-2.3.0.tgz",
				"integrity": "sha1-7RQaasBDqEnqWISY59yosVMw6Qw=",
				"dev": true
			  }
			}
		  },
		  "loader-runner": {
			"version": "2.4.0",
			"resolved": "https://registry.npmjs.org/loader-runner/-/loader-runner-2.4.0.tgz",
			"integrity": "sha512-Jsmr89RcXGIwivFY21FcRrisYZfvLMTWx5kOLc+JTxtpBOG6xML0vzbc6SEQG2FO9/4Fc3wW4LVcB5DmGflaRw==",
			"dev": true
		  },
		  "loader-utils": {
			"version": "1.4.0",
			"resolved": "https://registry.npmjs.org/loader-utils/-/loader-utils-1.4.0.tgz",
			"integrity": "sha512-qH0WSMBtn/oHuwjy/NucEgbx5dbxxnxup9s4PVXJUDHZBQY+s0NWA9rJf53RBnQZxfch7euUui7hpoAPvALZdA==",
			"dev": true,
			"requires": {
			  "big.js": "^5.2.2",
			  "emojis-list": "^3.0.0",
			  "json5": "^1.0.1"
			}
		  },
		  "locate-path": {
			"version": "3.0.0",
			"resolved": "https://registry.npmjs.org/locate-path/-/locate-path-3.0.0.tgz",
			"integrity": "sha512-7AO748wWnIhNqAuaty2ZWHkQHRSNfPVIsPIfwEOWO22AmaoVrWavlOcMR5nzTLNYvp36X220/maaRsrec1G65A==",
			"dev": true,
			"requires": {
			  "p-locate": "^3.0.0",
			  "path-exists": "^3.0.0"
			}
		  },
		  "lodash": {
			"version": "4.17.19",
			"resolved": "https://registry.npmjs.org/lodash/-/lodash-4.17.19.tgz",
			"integrity": "sha512-JNvd8XER9GQX0v2qJgsaN/mzFCNA5BRe/j8JN9d+tWyGLSodKQHKFicdwNYzWwI3wjRnaKPsGj1XkBjx/F96DQ==",
			"dev": true
		  },
		  "lodash.memoize": {
			"version": "4.1.2",
			"resolved": "https://registry.npmjs.org/lodash.memoize/-/lodash.memoize-4.1.2.tgz",
			"integrity": "sha1-vMbEmkKihA7Zl/Mj6tpezRguC/4=",
			"dev": true
		  },
		  "lodash.sortby": {
			"version": "4.7.0",
			"resolved": "https://registry.npmjs.org/lodash.sortby/-/lodash.sortby-4.7.0.tgz",
			"integrity": "sha1-7dFMgk4sycHgsKG0K7UhBRakJDg=",
			"dev": true
		  },
		  "lodash.toarray": {
			"version": "4.4.0",
			"resolved": "https://registry.npmjs.org/lodash.toarray/-/lodash.toarray-4.4.0.tgz",
			"integrity": "sha1-JMS/zWsvuji/0FlNsRedjptlZWE=",
			"dev": true
		  },
		  "log-symbols": {
			"version": "3.0.0",
			"resolved": "https://registry.npmjs.org/log-symbols/-/log-symbols-3.0.0.tgz",
			"integrity": "sha512-dSkNGuI7iG3mfvDzUuYZyvk5dD9ocYCYzNU6CYDE6+Xqd+gwme6Z00NS3dUh8mq/73HaEtT7m6W+yUPtU6BZnQ==",
			"dev": true,
			"requires": {
			  "chalk": "^2.4.2"
			}
		  },
		  "lru-cache": {
			"version": "5.1.1",
			"resolved": "https://registry.npmjs.org/lru-cache/-/lru-cache-5.1.1.tgz",
			"integrity": "sha512-KpNARQA3Iwv+jTA0utUVVbrh+Jlrr1Fv0e56GGzAFOXN7dk/FviaDW8LHmK52DlcH4WP2n6gI8vN1aesBFgo9w==",
			"dev": true,
			"requires": {
			  "yallist": "^3.0.2"
			}
		  },
		  "lru-queue": {
			"version": "0.1.0",
			"resolved": "https://registry.npmjs.org/lru-queue/-/lru-queue-0.1.0.tgz",
			"integrity": "sha1-Jzi9nw089PhEkMVzbEhpmsYyzaM=",
			"requires": {
			  "es5-ext": "~0.10.2"
			}
		  },
		  "macos-release": {
			"version": "2.3.0",
			"resolved": "https://registry.npmjs.org/macos-release/-/macos-release-2.3.0.tgz",
			"integrity": "sha512-OHhSbtcviqMPt7yfw5ef5aghS2jzFVKEFyCJndQt2YpSQ9qRVSEv2axSJI1paVThEu+FFGs584h/1YhxjVqajA==",
			"dev": true
		  },
		  "magic-string": {
			"version": "0.25.4",
			"resolved": "https://registry.npmjs.org/magic-string/-/magic-string-0.25.4.tgz",
			"integrity": "sha512-oycWO9nEVAP2RVPbIoDoA4Y7LFIJ3xRYov93gAyJhZkET1tNuB0u7uWkZS2LpBWTJUWnmau/To8ECWRC+jKNfw==",
			"dev": true,
			"requires": {
			  "sourcemap-codec": "^1.4.4"
			}
		  },
		  "make-dir": {
			"version": "2.1.0",
			"resolved": "https://registry.npmjs.org/make-dir/-/make-dir-2.1.0.tgz",
			"integrity": "sha512-LS9X+dc8KLxXCb8dni79fLIIUA5VyZoyjSMCwTluaXA0o27cCK0bhXkpgw+sTXVpPy/lSO57ilRixqk0vDmtRA==",
			"dev": true,
			"requires": {
			  "pify": "^4.0.1",
			  "semver": "^5.6.0"
			}
		  },
		  "make-error": {
			"version": "1.3.6",
			"resolved": "https://registry.npmjs.org/make-error/-/make-error-1.3.6.tgz",
			"integrity": "sha512-s8UhlNe7vPKomQhC1qFelMokr/Sc3AgNbso3n74mVPA5LTZwkB9NlXf4XPamLxJE8h0gh73rM94xvwRT2CVInw==",
			"dev": true
		  },
		  "makeerror": {
			"version": "1.0.11",
			"resolved": "https://registry.npmjs.org/makeerror/-/makeerror-1.0.11.tgz",
			"integrity": "sha1-4BpckQnyr3lmDk6LlYd5AYT1qWw=",
			"dev": true,
			"requires": {
			  "tmpl": "1.0.x"
			}
		  },
		  "map-cache": {
			"version": "0.2.2",
			"resolved": "https://registry.npmjs.org/map-cache/-/map-cache-0.2.2.tgz",
			"integrity": "sha1-wyq9C9ZSXZsFFkW7TyasXcmKDb8=",
			"dev": true
		  },
		  "map-visit": {
			"version": "1.0.0",
			"resolved": "https://registry.npmjs.org/map-visit/-/map-visit-1.0.0.tgz",
			"integrity": "sha1-7Nyo8TFE5mDxtb1B8S80edmN+48=",
			"dev": true,
			"requires": {
			  "object-visit": "^1.0.0"
			}
		  },
		  "md5.js": {
			"version": "1.3.5",
			"resolved": "https://registry.npmjs.org/md5.js/-/md5.js-1.3.5.tgz",
			"integrity": "sha512-xitP+WxNPcTTOgnTJcrhM0xvdPepipPSf3I8EIpGKeFLjt3PlJLIDG3u8EX53ZIubkb+5U2+3rELYpEhHhzdkg==",
			"dev": true,
			"requires": {
			  "hash-base": "^3.0.0",
			  "inherits": "^2.0.1",
			  "safe-buffer": "^5.1.2"
			}
		  },
		  "media-typer": {
			"version": "0.3.0",
			"resolved": "https://registry.npmjs.org/media-typer/-/media-typer-0.3.0.tgz",
			"integrity": "sha1-hxDXrwqmJvj/+hzgAWhUUmMlV0g="
		  },
		  "memoizee": {
			"version": "0.4.14",
			"resolved": "https://registry.npmjs.org/memoizee/-/memoizee-0.4.14.tgz",
			"integrity": "sha512-/SWFvWegAIYAO4NQMpcX+gcra0yEZu4OntmUdrBaWrJncxOqAziGFlHxc7yjKVK2uu3lpPW27P27wkR82wA8mg==",
			"requires": {
			  "d": "1",
			  "es5-ext": "^0.10.45",
			  "es6-weak-map": "^2.0.2",
			  "event-emitter": "^0.3.5",
			  "is-promise": "^2.1",
			  "lru-queue": "0.1",
			  "next-tick": "1",
			  "timers-ext": "^0.1.5"
			}
		  },
		  "memory-fs": {
			"version": "0.5.0",
			"resolved": "https://registry.npmjs.org/memory-fs/-/memory-fs-0.5.0.tgz",
			"integrity": "sha512-jA0rdU5KoQMC0e6ppoNRtpp6vjFq6+NY7r8hywnC7V+1Xj/MtHwGIbB1QaK/dunyjWteJzmkpd7ooeWg10T7GA==",
			"dev": true,
			"requires": {
			  "errno": "^0.1.3",
			  "readable-stream": "^2.0.1"
			},
			"dependencies": {
			  "isarray": {
				"version": "1.0.0",
				"resolved": "https://registry.npmjs.org/isarray/-/isarray-1.0.0.tgz",
				"integrity": "sha1-u5NdSFgsuhaMBoNJV6VKPgcSTxE=",
				"dev": true
			  },
			  "readable-stream": {
				"version": "2.3.7",
				"resolved": "https://registry.npmjs.org/readable-stream/-/readable-stream-2.3.7.tgz",
				"integrity": "sha512-Ebho8K4jIbHAxnuxi7o42OrZgF/ZTNcsZj6nRKyUmkhLFq8CHItp/fy6hQZuZmP/n3yZ9VBUbp4zz/mX8hmYPw==",
				"dev": true,
				"requires": {
				  "core-util-is": "~1.0.0",
				  "inherits": "~2.0.3",
				  "isarray": "~1.0.0",
				  "process-nextick-args": "~2.0.0",
				  "safe-buffer": "~5.1.1",
				  "string_decoder": "~1.1.1",
				  "util-deprecate": "~1.0.1"
				}
			  },
			  "string_decoder": {
				"version": "1.1.1",
				"resolved": "https://registry.npmjs.org/string_decoder/-/string_decoder-1.1.1.tgz",
				"integrity": "sha512-n/ShnvDi6FHbbVfviro+WojiFzv+s8MPMHBczVePfUpDJLwoLT0ht1l4YwBCbi8pJAveEEdnkHyPyTP/mzRfwg==",
				"dev": true,
				"requires": {
				  "safe-buffer": "~5.1.0"
				}
			  }
			}
		  },
		  "merge-descriptors": {
			"version": "1.0.1",
			"resolved": "https://registry.npmjs.org/merge-descriptors/-/merge-descriptors-1.0.1.tgz",
			"integrity": "sha1-sAqqVW3YtEVoFQ7J0blT8/kMu2E="
		  },
		  "merge-stream": {
			"version": "2.0.0",
			"resolved": "https://registry.npmjs.org/merge-stream/-/merge-stream-2.0.0.tgz",
			"integrity": "sha512-abv/qOcuPfk3URPfDzmZU1LKmuw8kT+0nIHvKrKgFrwifol/doWcdA4ZqsWQ8ENrFKkd67Mfpo/LovbIUsbt3w==",
			"dev": true
		  },
		  "methods": {
			"version": "1.1.2",
			"resolved": "https://registry.npmjs.org/methods/-/methods-1.1.2.tgz",
			"integrity": "sha1-VSmk1nZUE07cxSZmVoNbD4Ua/O4="
		  },
		  "microevent.ts": {
			"version": "0.1.1",
			"resolved": "https://registry.npmjs.org/microevent.ts/-/microevent.ts-0.1.1.tgz",
			"integrity": "sha512-jo1OfR4TaEwd5HOrt5+tAZ9mqT4jmpNAusXtyfNzqVm9uiSYFZlKM1wYL4oU7azZW/PxQW53wM0S6OR1JHNa2g==",
			"dev": true
		  },
		  "micromatch": {
			"version": "3.1.10",
			"resolved": "https://registry.npmjs.org/micromatch/-/micromatch-3.1.10.tgz",
			"integrity": "sha512-MWikgl9n9M3w+bpsY3He8L+w9eF9338xRl8IAO5viDizwSzziFEyUzo2xrrloB64ADbTf8uA8vRqqttDTOmccg==",
			"dev": true,
			"requires": {
			  "arr-diff": "^4.0.0",
			  "array-unique": "^0.3.2",
			  "braces": "^2.3.1",
			  "define-property": "^2.0.2",
			  "extend-shallow": "^3.0.2",
			  "extglob": "^2.0.4",
			  "fragment-cache": "^0.2.1",
			  "kind-of": "^6.0.2",
			  "nanomatch": "^1.2.9",
			  "object.pick": "^1.3.0",
			  "regex-not": "^1.0.0",
			  "snapdragon": "^0.8.1",
			  "to-regex": "^3.0.2"
			},
			"dependencies": {
			  "braces": {
				"version": "2.3.2",
				"resolved": "https://registry.npmjs.org/braces/-/braces-2.3.2.tgz",
				"integrity": "sha512-aNdbnj9P8PjdXU4ybaWLK2IF3jc/EoDYbC7AazW6to3TRsfXxscC9UXOB5iDiEQrkyIbWp2SLQda4+QAa7nc3w==",
				"dev": true,
				"requires": {
				  "arr-flatten": "^1.1.0",
				  "array-unique": "^0.3.2",
				  "extend-shallow": "^2.0.1",
				  "fill-range": "^4.0.0",
				  "isobject": "^3.0.1",
				  "repeat-element": "^1.1.2",
				  "snapdragon": "^0.8.1",
				  "snapdragon-node": "^2.0.1",
				  "split-string": "^3.0.2",
				  "to-regex": "^3.0.1"
				},
				"dependencies": {
				  "extend-shallow": {
					"version": "2.0.1",
					"resolved": "https://registry.npmjs.org/extend-shallow/-/extend-shallow-2.0.1.tgz",
					"integrity": "sha1-Ua99YUrZqfYQ6huvu5idaxxWiQ8=",
					"dev": true,
					"requires": {
					  "is-extendable": "^0.1.0"
					}
				  }
				}
			  },
			  "fill-range": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/fill-range/-/fill-range-4.0.0.tgz",
				"integrity": "sha1-1USBHUKPmOsGpj3EAtJAPDKMOPc=",
				"dev": true,
				"requires": {
				  "extend-shallow": "^2.0.1",
				  "is-number": "^3.0.0",
				  "repeat-string": "^1.6.1",
				  "to-regex-range": "^2.1.0"
				},
				"dependencies": {
				  "extend-shallow": {
					"version": "2.0.1",
					"resolved": "https://registry.npmjs.org/extend-shallow/-/extend-shallow-2.0.1.tgz",
					"integrity": "sha1-Ua99YUrZqfYQ6huvu5idaxxWiQ8=",
					"dev": true,
					"requires": {
					  "is-extendable": "^0.1.0"
					}
				  }
				}
			  },
			  "is-number": {
				"version": "3.0.0",
				"resolved": "https://registry.npmjs.org/is-number/-/is-number-3.0.0.tgz",
				"integrity": "sha1-JP1iAaR4LPUFYcgQJ2r8fRLXEZU=",
				"dev": true,
				"requires": {
				  "kind-of": "^3.0.2"
				},
				"dependencies": {
				  "kind-of": {
					"version": "3.2.2",
					"resolved": "https://registry.npmjs.org/kind-of/-/kind-of-3.2.2.tgz",
					"integrity": "sha1-MeohpzS6ubuw8yRm2JOupR5KPGQ=",
					"dev": true,
					"requires": {
					  "is-buffer": "^1.1.5"
					}
				  }
				}
			  },
			  "to-regex-range": {
				"version": "2.1.1",
				"resolved": "https://registry.npmjs.org/to-regex-range/-/to-regex-range-2.1.1.tgz",
				"integrity": "sha1-fIDBe53+vlmeJzZ+DU3VWQFB2zg=",
				"dev": true,
				"requires": {
				  "is-number": "^3.0.0",
				  "repeat-string": "^1.6.1"
				}
			  }
			}
		  },
		  "miller-rabin": {
			"version": "4.0.1",
			"resolved": "https://registry.npmjs.org/miller-rabin/-/miller-rabin-4.0.1.tgz",
			"integrity": "sha512-115fLhvZVqWwHPbClyntxEVfVDfl9DLLTuJvq3g2O/Oxi8AiNouAHvDSzHS0viUJc+V5vm3eq91Xwqn9dp4jRA==",
			"dev": true,
			"requires": {
			  "bn.js": "^4.0.0",
			  "brorand": "^1.0.1"
			},
			"dependencies": {
			  "bn.js": {
				"version": "4.11.9",
				"resolved": "https://registry.npmjs.org/bn.js/-/bn.js-4.11.9.tgz",
				"integrity": "sha512-E6QoYqCKZfgatHTdHzs1RRKP7ip4vvm+EyRUeE2RF0NblwVvb0p6jSVeNTOFxPn26QXN2o6SMfNxKp6kU8zQaw==",
				"dev": true
			  }
			}
		  },
		  "mime": {
			"version": "1.6.0",
			"resolved": "https://registry.npmjs.org/mime/-/mime-1.6.0.tgz",
			"integrity": "sha512-x0Vn8spI+wuJ1O6S7gnbaQg8Pxh4NNHb7KSINmEWKiPE4RKOplvijn+NkmYmmRgP68mc70j2EbeTFRsrswaQeg=="
		  },
		  "mime-db": {
			"version": "1.40.0",
			"resolved": "https://registry.npmjs.org/mime-db/-/mime-db-1.40.0.tgz",
			"integrity": "sha512-jYdeOMPy9vnxEqFRRo6ZvTZ8d9oPb+k18PKoYNYUe2stVEBPPwsln/qWzdbmaIvnhZ9v2P+CuecK+fpUfsV2mA=="
		  },
		  "mime-types": {
			"version": "2.1.24",
			"resolved": "https://registry.npmjs.org/mime-types/-/mime-types-2.1.24.tgz",
			"integrity": "sha512-WaFHS3MCl5fapm3oLxU4eYDw77IQM2ACcxQ9RIxfaC3ooc6PFuBMGZZsYpvoXS5D5QTWPieo1jjLdAm3TBP3cQ==",
			"requires": {
			  "mime-db": "1.40.0"
			}
		  },
		  "mimic-fn": {
			"version": "2.1.0",
			"resolved": "https://registry.npmjs.org/mimic-fn/-/mimic-fn-2.1.0.tgz",
			"integrity": "sha512-OqbOk5oEQeAZ8WXWydlu9HJjz9WVdEIvamMCcXmuqUYjTknH/sqsWvhQ3vgwKFRR1HpjvNBKQ37nbJgYzGqGcg==",
			"dev": true
		  },
		  "minimalistic-assert": {
			"version": "1.0.1",
			"resolved": "https://registry.npmjs.org/minimalistic-assert/-/minimalistic-assert-1.0.1.tgz",
			"integrity": "sha512-UtJcAD4yEaGtjPezWuO9wC4nwUnVH/8/Im3yEHQP4b67cXlD/Qr9hdITCU1xDbSEXg2XKNaP8jsReV7vQd00/A==",
			"dev": true
		  },
		  "minimalistic-crypto-utils": {
			"version": "1.0.1",
			"resolved": "https://registry.npmjs.org/minimalistic-crypto-utils/-/minimalistic-crypto-utils-1.0.1.tgz",
			"integrity": "sha1-9sAMHAsIIkblxNmd+4x8CDsrWCo=",
			"dev": true
		  },
		  "minimatch": {
			"version": "3.0.4",
			"resolved": "https://registry.npmjs.org/minimatch/-/minimatch-3.0.4.tgz",
			"integrity": "sha512-yJHVQEhyqPLUTgt9B83PXu6W3rx4MvvHvSUvToogpwoGDOUQ+yDrR0HRot+yOCdCO7u4hX3pWft6kWBBcqh0UA==",
			"requires": {
			  "brace-expansion": "^1.1.7"
			}
		  },
		  "minimist": {
			"version": "1.2.5",
			"resolved": "https://registry.npmjs.org/minimist/-/minimist-1.2.5.tgz",
			"integrity": "sha512-FM9nNUYrRBAELZQT3xeZQ7fmMOBg6nWNmJKTcgsJeaLstP/UODVpGsr5OhXhhXg6f+qtJ8uiZ+PUxkDWcgIXLw=="
		  },
		  "mississippi": {
			"version": "3.0.0",
			"resolved": "https://registry.npmjs.org/mississippi/-/mississippi-3.0.0.tgz",
			"integrity": "sha512-x471SsVjUtBRtcvd4BzKE9kFC+/2TeWgKCgw0bZcw1b9l2X3QX5vCWgF+KaZaYm87Ss//rHnWryupDrgLvmSkA==",
			"dev": true,
			"requires": {
			  "concat-stream": "^1.5.0",
			  "duplexify": "^3.4.2",
			  "end-of-stream": "^1.1.0",
			  "flush-write-stream": "^1.0.0",
			  "from2": "^2.1.0",
			  "parallel-transform": "^1.1.0",
			  "pump": "^3.0.0",
			  "pumpify": "^1.3.3",
			  "stream-each": "^1.1.0",
			  "through2": "^2.0.0"
			}
		  },
		  "mixin-deep": {
			"version": "1.3.2",
			"resolved": "https://registry.npmjs.org/mixin-deep/-/mixin-deep-1.3.2.tgz",
			"integrity": "sha512-WRoDn//mXBiJ1H40rqa3vH0toePwSsGb45iInWlTySa+Uu4k3tYUSxa2v1KqAiLtvlrSzaExqS1gtk96A9zvEA==",
			"dev": true,
			"requires": {
			  "for-in": "^1.0.2",
			  "is-extendable": "^1.0.1"
			},
			"dependencies": {
			  "is-extendable": {
				"version": "1.0.1",
				"resolved": "https://registry.npmjs.org/is-extendable/-/is-extendable-1.0.1.tgz",
				"integrity": "sha512-arnXMxT1hhoKo9k1LZdmlNyJdDDfy2v0fXjFlmok4+i8ul/6WlbVge9bhM74OpNPQPMGUToDtz+KXa1PneJxOA==",
				"dev": true,
				"requires": {
				  "is-plain-object": "^2.0.4"
				}
			  }
			}
		  },
		  "mkdirp": {
			"version": "0.5.5",
			"resolved": "https://registry.npmjs.org/mkdirp/-/mkdirp-0.5.5.tgz",
			"integrity": "sha512-NKmAlESf6jMGym1++R0Ra7wvhV+wFW63FaSOFPwRahvea0gMUcGUhVeAg/0BC0wiv9ih5NYPB1Wn1UEI1/L+xQ==",
			"requires": {
			  "minimist": "^1.2.5"
			}
		  },
		  "move-concurrently": {
			"version": "1.0.1",
			"resolved": "https://registry.npmjs.org/move-concurrently/-/move-concurrently-1.0.1.tgz",
			"integrity": "sha1-viwAX9oy4LKa8fBdfEszIUxwH5I=",
			"dev": true,
			"requires": {
			  "aproba": "^1.1.1",
			  "copy-concurrently": "^1.0.0",
			  "fs-write-stream-atomic": "^1.0.8",
			  "mkdirp": "^0.5.1",
			  "rimraf": "^2.5.4",
			  "run-queue": "^1.0.3"
			},
			"dependencies": {
			  "rimraf": {
				"version": "2.7.1",
				"resolved": "https://registry.npmjs.org/rimraf/-/rimraf-2.7.1.tgz",
				"integrity": "sha512-uWjbaKIK3T1OSVptzX7Nl6PvQ3qAGtKEtVRjRuazjfL3Bx5eI409VZSqgND+4UNnmzLVdPj9FqFJNPqBZFve4w==",
				"dev": true,
				"requires": {
				  "glob": "^7.1.3"
				}
			  }
			}
		  },
		  "ms": {
			"version": "2.0.0",
			"resolved": "https://registry.npmjs.org/ms/-/ms-2.0.0.tgz",
			"integrity": "sha1-VgiurfwAvmwpAd9fmGF4jeDVl8g="
		  },
		  "multer": {
			"version": "1.4.2",
			"resolved": "https://registry.npmjs.org/multer/-/multer-1.4.2.tgz",
			"integrity": "sha512-xY8pX7V+ybyUpbYMxtjM9KAiD9ixtg5/JkeKUTD6xilfDv0vzzOFcCp4Ljb1UU3tSOM3VTZtKo63OmzOrGi3Cg==",
			"requires": {
			  "append-field": "^1.0.0",
			  "busboy": "^0.2.11",
			  "concat-stream": "^1.5.2",
			  "mkdirp": "^0.5.1",
			  "object-assign": "^4.1.1",
			  "on-finished": "^2.3.0",
			  "type-is": "^1.6.4",
			  "xtend": "^4.0.0"
			}
		  },
		  "mute-stream": {
			"version": "0.0.8",
			"resolved": "https://registry.npmjs.org/mute-stream/-/mute-stream-0.0.8.tgz",
			"integrity": "sha512-nnbWWOkoWyUsTjKrhgD0dcz22mdkSnpYqbEjIm2nhwhuxlSkpywJmBo8h0ZqJdkp73mb90SssHkN4rsRaBAfAA==",
			"dev": true
		  },
		  "nan": {
			"version": "2.14.1",
			"resolved": "https://registry.npmjs.org/nan/-/nan-2.14.1.tgz",
			"integrity": "sha512-isWHgVjnFjh2x2yuJ/tj3JbwoHu3UC2dX5G/88Cm24yB6YopVgxvBObDY7n5xW6ExmFhJpSEQqFPvq9zaXc8Jw==",
			"dev": true,
			"optional": true
		  },
		  "nanomatch": {
			"version": "1.2.13",
			"resolved": "https://registry.npmjs.org/nanomatch/-/nanomatch-1.2.13.tgz",
			"integrity": "sha512-fpoe2T0RbHwBTBUOftAfBPaDEi06ufaUai0mE6Yn1kacc3SnTErfb/h+X94VXzI64rKFHYImXSvdwGGCmwOqCA==",
			"dev": true,
			"requires": {
			  "arr-diff": "^4.0.0",
			  "array-unique": "^0.3.2",
			  "define-property": "^2.0.2",
			  "extend-shallow": "^3.0.2",
			  "fragment-cache": "^0.2.1",
			  "is-windows": "^1.0.2",
			  "kind-of": "^6.0.2",
			  "object.pick": "^1.3.0",
			  "regex-not": "^1.0.0",
			  "snapdragon": "^0.8.1",
			  "to-regex": "^3.0.1"
			}
		  },
		  "natural-compare": {
			"version": "1.4.0",
			"resolved": "https://registry.npmjs.org/natural-compare/-/natural-compare-1.4.0.tgz",
			"integrity": "sha1-Sr6/7tdUHywnrPspvbvRXI1bpPc=",
			"dev": true
		  },
		  "negotiator": {
			"version": "0.6.2",
			"resolved": "https://registry.npmjs.org/negotiator/-/negotiator-0.6.2.tgz",
			"integrity": "sha512-hZXc7K2e+PgeI1eDBe/10Ard4ekbfrrqG8Ep+8Jmf4JID2bNg7NvCPOZN+kfF574pFQI7mum2AUqDidoKqcTOw=="
		  },
		  "neo-async": {
			"version": "2.6.1",
			"resolved": "https://registry.npmjs.org/neo-async/-/neo-async-2.6.1.tgz",
			"integrity": "sha512-iyam8fBuCUpWeKPGpaNMetEocMt364qkCsfL9JuhjXX6dRnguRVOfk2GZaDpPjcOKiiXCPINZC1GczQ7iTq3Zw==",
			"dev": true
		  },
		  "next-tick": {
			"version": "1.0.0",
			"resolved": "https://registry.npmjs.org/next-tick/-/next-tick-1.0.0.tgz",
			"integrity": "sha1-yobR/ogoFpsBICCOPchCS524NCw="
		  },
		  "nice-try": {
			"version": "1.0.5",
			"resolved": "https://registry.npmjs.org/nice-try/-/nice-try-1.0.5.tgz",
			"integrity": "sha512-1nh45deeb5olNY7eX82BkPO7SSxR5SSYJiPTrTdFUVYwAl8CKMA5N9PjTYkHiRjisVcxcQ1HXdLhx2qxxJzLNQ==",
			"dev": true
		  },
		  "node-emoji": {
			"version": "1.10.0",
			"resolved": "https://registry.npmjs.org/node-emoji/-/node-emoji-1.10.0.tgz",
			"integrity": "sha512-Yt3384If5H6BYGVHiHwTL+99OzJKHhgp82S8/dktEK73T26BazdgZ4JZh92xSVtGNJvz9UbXdNAc5hcrXV42vw==",
			"dev": true,
			"requires": {
			  "lodash.toarray": "^4.4.0"
			}
		  },
		  "node-fetch": {
			"version": "2.6.1",
			"resolved": "https://registry.npmjs.org/node-fetch/-/node-fetch-2.6.1.tgz",
			"integrity": "sha512-V4aYg89jEoVRxRb2fJdAg8FHvI7cEyYdVAh94HH0UIK8oJxUfkjlDQN9RbMx+bEjP7+ggMiFRprSti032Oipxw=="
		  },
		  "node-int64": {
			"version": "0.4.0",
			"resolved": "https://registry.npmjs.org/node-int64/-/node-int64-0.4.0.tgz",
			"integrity": "sha1-h6kGXNs1XTGC2PlM4RGIuCXGijs=",
			"dev": true
		  },
		  "node-libs-browser": {
			"version": "2.2.1",
			"resolved": "https://registry.npmjs.org/node-libs-browser/-/node-libs-browser-2.2.1.tgz",
			"integrity": "sha512-h/zcD8H9kaDZ9ALUWwlBUDo6TKF8a7qBSCSEGfjTVIYeqsioSKaAX+BN7NgiMGp6iSIXZ3PxgCu8KS3b71YK5Q==",
			"dev": true,
			"requires": {
			  "assert": "^1.1.1",
			  "browserify-zlib": "^0.2.0",
			  "buffer": "^4.3.0",
			  "console-browserify": "^1.1.0",
			  "constants-browserify": "^1.0.0",
			  "crypto-browserify": "^3.11.0",
			  "domain-browser": "^1.1.1",
			  "events": "^3.0.0",
			  "https-browserify": "^1.0.0",
			  "os-browserify": "^0.3.0",
			  "path-browserify": "0.0.1",
			  "process": "^0.11.10",
			  "punycode": "^1.2.4",
			  "querystring-es3": "^0.2.0",
			  "readable-stream": "^2.3.3",
			  "stream-browserify": "^2.0.1",
			  "stream-http": "^2.7.2",
			  "string_decoder": "^1.0.0",
			  "timers-browserify": "^2.0.4",
			  "tty-browserify": "0.0.0",
			  "url": "^0.11.0",
			  "util": "^0.11.0",
			  "vm-browserify": "^1.0.1"
			},
			"dependencies": {
			  "isarray": {
				"version": "1.0.0",
				"resolved": "https://registry.npmjs.org/isarray/-/isarray-1.0.0.tgz",
				"integrity": "sha1-u5NdSFgsuhaMBoNJV6VKPgcSTxE=",
				"dev": true
			  },
			  "punycode": {
				"version": "1.4.1",
				"resolved": "https://registry.npmjs.org/punycode/-/punycode-1.4.1.tgz",
				"integrity": "sha1-wNWmOycYgArY4esPpSachN1BhF4=",
				"dev": true
			  },
			  "readable-stream": {
				"version": "2.3.7",
				"resolved": "https://registry.npmjs.org/readable-stream/-/readable-stream-2.3.7.tgz",
				"integrity": "sha512-Ebho8K4jIbHAxnuxi7o42OrZgF/ZTNcsZj6nRKyUmkhLFq8CHItp/fy6hQZuZmP/n3yZ9VBUbp4zz/mX8hmYPw==",
				"dev": true,
				"requires": {
				  "core-util-is": "~1.0.0",
				  "inherits": "~2.0.3",
				  "isarray": "~1.0.0",
				  "process-nextick-args": "~2.0.0",
				  "safe-buffer": "~5.1.1",
				  "string_decoder": "~1.1.1",
				  "util-deprecate": "~1.0.1"
				},
				"dependencies": {
				  "string_decoder": {
					"version": "1.1.1",
					"resolved": "https://registry.npmjs.org/string_decoder/-/string_decoder-1.1.1.tgz",
					"integrity": "sha512-n/ShnvDi6FHbbVfviro+WojiFzv+s8MPMHBczVePfUpDJLwoLT0ht1l4YwBCbi8pJAveEEdnkHyPyTP/mzRfwg==",
					"dev": true,
					"requires": {
					  "safe-buffer": "~5.1.0"
					}
				  }
				}
			  },
			  "string_decoder": {
				"version": "1.3.0",
				"resolved": "https://registry.npmjs.org/string_decoder/-/string_decoder-1.3.0.tgz",
				"integrity": "sha512-hkRX8U1WjJFd8LsDJ2yQ/wWWxaopEsABU1XfkM8A+j0+85JAGppt16cr1Whg6KIbb4okU6Mql6BOj+uup/wKeA==",
				"dev": true,
				"requires": {
				  "safe-buffer": "~5.2.0"
				},
				"dependencies": {
				  "safe-buffer": {
					"version": "5.2.1",
					"resolved": "https://registry.npmjs.org/safe-buffer/-/safe-buffer-5.2.1.tgz",
					"integrity": "sha512-rp3So07KcdmmKbGvgaNxQSJr7bGVSVk5S9Eq1F+ppbRo70+YeaDxkw5Dd8NPN+GD6bjnYm2VuPuCXmpuYvmCXQ==",
					"dev": true
				  }
				}
			  }
			}
		  },
		  "node-modules-regexp": {
			"version": "1.0.0",
			"resolved": "https://registry.npmjs.org/node-modules-regexp/-/node-modules-regexp-1.0.0.tgz",
			"integrity": "sha1-jZ2+KJZKSsVxLpExZCEHxx6Q7EA=",
			"dev": true
		  },
		  "node-notifier": {
			"version": "7.0.0",
			"resolved": "https://registry.npmjs.org/node-notifier/-/node-notifier-7.0.0.tgz",
			"integrity": "sha512-y8ThJESxsHcak81PGpzWwQKxzk+5YtP3IxR8AYdpXQ1IB6FmcVzFdZXrkPin49F/DKUCfeeiziB8ptY9npzGuA==",
			"dev": true,
			"optional": true,
			"requires": {
			  "growly": "^1.3.0",
			  "is-wsl": "^2.1.1",
			  "semver": "^7.2.1",
			  "shellwords": "^0.1.1",
			  "uuid": "^7.0.3",
			  "which": "^2.0.2"
			},
			"dependencies": {
			  "is-wsl": {
				"version": "2.2.0",
				"resolved": "https://registry.npmjs.org/is-wsl/-/is-wsl-2.2.0.tgz",
				"integrity": "sha512-fKzAra0rGJUUBwGBgNkHZuToZcn+TtXHpeCgmkMJMMYx1sQDYaCSyjJBSCa2nH1DGm7s3n1oBnohoVTBaN7Lww==",
				"dev": true,
				"optional": true,
				"requires": {
				  "is-docker": "^2.0.0"
				}
			  },
			  "semver": {
				"version": "7.3.2",
				"resolved": "https://registry.npmjs.org/semver/-/semver-7.3.2.tgz",
				"integrity": "sha512-OrOb32TeeambH6UrhtShmF7CRDqhL6/5XpPNp2DuRH6+9QLw/orhp72j87v8Qa1ScDkvrrBNpZcDejAirJmfXQ==",
				"dev": true,
				"optional": true
			  },
			  "which": {
				"version": "2.0.2",
				"resolved": "https://registry.npmjs.org/which/-/which-2.0.2.tgz",
				"integrity": "sha512-BLI3Tl1TW3Pvl70l3yq3Y64i+awpwXqsGBYWkkqMtnbXgrMD+yj7rhW0kuEDxzJaYXGjEW5ogapKNMEKNMjibA==",
				"dev": true,
				"optional": true,
				"requires": {
				  "isexe": "^2.0.0"
				}
			  }
			}
		  },
		  "normalize-package-data": {
			"version": "2.5.0",
			"resolved": "https://registry.npmjs.org/normalize-package-data/-/normalize-package-data-2.5.0.tgz",
			"integrity": "sha512-/5CMN3T0R4XTj4DcGaexo+roZSdSFW/0AOOTROrjxzCG1wrWXEsGbRKevjlIL+ZDE4sZlJr5ED4YW0yqmkK+eA==",
			"dev": true,
			"requires": {
			  "hosted-git-info": "^2.1.4",
			  "resolve": "^1.10.0",
			  "semver": "2 || 3 || 4 || 5",
			  "validate-npm-package-license": "^3.0.1"
			}
		  },
		  "normalize-path": {
			"version": "3.0.0",
			"resolved": "https://registry.npmjs.org/normalize-path/-/normalize-path-3.0.0.tgz",
			"integrity": "sha512-6eZs5Ls3WtCisHWp9S2GUy8dqkpGi4BVSz3GaqiE6ezub0512ESztXUwUB6C6IKbQkY2Pnb/mD4WYojCRwcwLA==",
			"dev": true
		  },
		  "npm-run-path": {
			"version": "2.0.2",
			"resolved": "https://registry.npmjs.org/npm-run-path/-/npm-run-path-2.0.2.tgz",
			"integrity": "sha1-NakjLfo11wZ7TLLd8jV7GHFTbF8=",
			"dev": true,
			"requires": {
			  "path-key": "^2.0.0"
			}
		  },
		  "nwsapi": {
			"version": "2.2.0",
			"resolved": "https://registry.npmjs.org/nwsapi/-/nwsapi-2.2.0.tgz",
			"integrity": "sha512-h2AatdwYH+JHiZpv7pt/gSX1XoRGb7L/qSIeuqA6GwYoF9w1vP1cw42TO0aI2pNyshRK5893hNSl+1//vHK7hQ==",
			"dev": true
		  },
		  "oauth-sign": {
			"version": "0.9.0",
			"resolved": "https://registry.npmjs.org/oauth-sign/-/oauth-sign-0.9.0.tgz",
			"integrity": "sha512-fexhUFFPTGV8ybAtSIGbV6gOkSv8UtRbDBnAyLQw4QPKkgNlsH2ByPGtMUqdWkos6YCRmAqViwgZrJc/mRDzZQ==",
			"dev": true
		  },
		  "object-assign": {
			"version": "4.1.1",
			"resolved": "https://registry.npmjs.org/object-assign/-/object-assign-4.1.1.tgz",
			"integrity": "sha1-IQmtx5ZYh8/AXLvUQsrIv7s2CGM="
		  },
		  "object-copy": {
			"version": "0.1.0",
			"resolved": "https://registry.npmjs.org/object-copy/-/object-copy-0.1.0.tgz",
			"integrity": "sha1-fn2Fi3gb18mRpBupde04EnVOmYw=",
			"dev": true,
			"requires": {
			  "copy-descriptor": "^0.1.0",
			  "define-property": "^0.2.5",
			  "kind-of": "^3.0.3"
			},
			"dependencies": {
			  "define-property": {
				"version": "0.2.5",
				"resolved": "https://registry.npmjs.org/define-property/-/define-property-0.2.5.tgz",
				"integrity": "sha1-w1se+RjsPJkPmlvFe+BKrOxcgRY=",
				"dev": true,
				"requires": {
				  "is-descriptor": "^0.1.0"
				}
			  },
			  "kind-of": {
				"version": "3.2.2",
				"resolved": "https://registry.npmjs.org/kind-of/-/kind-of-3.2.2.tgz",
				"integrity": "sha1-MeohpzS6ubuw8yRm2JOupR5KPGQ=",
				"dev": true,
				"requires": {
				  "is-buffer": "^1.1.5"
				}
			  }
			}
		  },
		  "object-hash": {
			"version": "2.0.3",
			"resolved": "https://registry.npmjs.org/object-hash/-/object-hash-2.0.3.tgz",
			"integrity": "sha512-JPKn0GMu+Fa3zt3Bmr66JhokJU5BaNBIh4ZeTlaCBzrBsOeXzwcKKAK1tbLiPKgvwmPXsDvvLHoWh5Bm7ofIYg=="
		  },
		  "object-inspect": {
			"version": "1.7.0",
			"resolved": "https://registry.npmjs.org/object-inspect/-/object-inspect-1.7.0.tgz",
			"integrity": "sha512-a7pEHdh1xKIAgTySUGgLMx/xwDZskN1Ud6egYYN3EdRW4ZMPNEDUTF+hwy2LUC+Bl+SyLXANnwz/jyh/qutKUw==",
			"dev": true
		  },
		  "object-keys": {
			"version": "1.1.1",
			"resolved": "https://registry.npmjs.org/object-keys/-/object-keys-1.1.1.tgz",
			"integrity": "sha512-NuAESUOUMrlIXOfHKzD6bpPu3tYt3xvjNdRIQ+FeT0lNb4K8WR70CaDxhuNguS2XG+GjkyMwOzsN5ZktImfhLA==",
			"dev": true
		  },
		  "object-visit": {
			"version": "1.0.1",
			"resolved": "https://registry.npmjs.org/object-visit/-/object-visit-1.0.1.tgz",
			"integrity": "sha1-95xEk68MU3e1n+OdOV5BBC3QRbs=",
			"dev": true,
			"requires": {
			  "isobject": "^3.0.0"
			}
		  },
		  "object.assign": {
			"version": "4.1.0",
			"resolved": "https://registry.npmjs.org/object.assign/-/object.assign-4.1.0.tgz",
			"integrity": "sha512-exHJeq6kBKj58mqGyTQ9DFvrZC/eR6OwxzoM9YRoGBqrXYonaFyGiFMuc9VZrXf7DarreEwMpurG3dd+CNyW5w==",
			"dev": true,
			"requires": {
			  "define-properties": "^1.1.2",
			  "function-bind": "^1.1.1",
			  "has-symbols": "^1.0.0",
			  "object-keys": "^1.0.11"
			}
		  },
		  "object.pick": {
			"version": "1.3.0",
			"resolved": "https://registry.npmjs.org/object.pick/-/object.pick-1.3.0.tgz",
			"integrity": "sha1-h6EKxMFpS9Lhy/U1kaZhQftd10c=",
			"dev": true,
			"requires": {
			  "isobject": "^3.0.1"
			}
		  },
		  "object.values": {
			"version": "1.1.1",
			"resolved": "https://registry.npmjs.org/object.values/-/object.values-1.1.1.tgz",
			"integrity": "sha512-WTa54g2K8iu0kmS/us18jEmdv1a4Wi//BZ/DTVYEcH0XhLM5NYdpDHja3gt57VrZLcNAO2WGA+KpWsDBaHt6eA==",
			"dev": true,
			"requires": {
			  "define-properties": "^1.1.3",
			  "es-abstract": "^1.17.0-next.1",
			  "function-bind": "^1.1.1",
			  "has": "^1.0.3"
			}
		  },
		  "on-finished": {
			"version": "2.3.0",
			"resolved": "https://registry.npmjs.org/on-finished/-/on-finished-2.3.0.tgz",
			"integrity": "sha1-IPEzZIGwg811M3mSoWlxqi2QaUc=",
			"requires": {
			  "ee-first": "1.1.1"
			}
		  },
		  "once": {
			"version": "1.4.0",
			"resolved": "https://registry.npmjs.org/once/-/once-1.4.0.tgz",
			"integrity": "sha1-WDsap3WWHUsROsF9nFC6753Xa9E=",
			"requires": {
			  "wrappy": "1"
			}
		  },
		  "onetime": {
			"version": "5.1.0",
			"resolved": "https://registry.npmjs.org/onetime/-/onetime-5.1.0.tgz",
			"integrity": "sha512-5NcSkPHhwTVFIQN+TUqXoS5+dlElHXdpAWu9I0HP20YOtIi+aZ0Ct82jdlILDxjLEAWwvm+qj1m6aEtsDVmm6Q==",
			"dev": true,
			"requires": {
			  "mimic-fn": "^2.1.0"
			}
		  },
		  "optional": {
			"version": "0.1.4",
			"resolved": "https://registry.npmjs.org/optional/-/optional-0.1.4.tgz",
			"integrity": "sha512-gtvrrCfkE08wKcgXaVwQVgwEQ8vel2dc5DDBn9RLQZ3YtmtkBss6A2HY6BnJH4N/4Ku97Ri/SF8sNWE2225WJw==",
			"dev": true
		  },
		  "optionator": {
			"version": "0.9.1",
			"resolved": "https://registry.npmjs.org/optionator/-/optionator-0.9.1.tgz",
			"integrity": "sha512-74RlY5FCnhq4jRxVUPKDaRwrVNXMqsGsiW6AJw4XK8hmtm10wC0ypZBLw5IIp85NZMr91+qd1RvvENwg7jjRFw==",
			"dev": true,
			"requires": {
			  "deep-is": "^0.1.3",
			  "fast-levenshtein": "^2.0.6",
			  "levn": "^0.4.1",
			  "prelude-ls": "^1.2.1",
			  "type-check": "^0.4.0",
			  "word-wrap": "^1.2.3"
			}
		  },
		  "ora": {
			"version": "4.0.4",
			"resolved": "https://registry.npmjs.org/ora/-/ora-4.0.4.tgz",
			"integrity": "sha512-77iGeVU1cIdRhgFzCK8aw1fbtT1B/iZAvWjS+l/o1x0RShMgxHUZaD2yDpWsNCPwXg9z1ZA78Kbdvr8kBmG/Ww==",
			"dev": true,
			"requires": {
			  "chalk": "^3.0.0",
			  "cli-cursor": "^3.1.0",
			  "cli-spinners": "^2.2.0",
			  "is-interactive": "^1.0.0",
			  "log-symbols": "^3.0.0",
			  "mute-stream": "0.0.8",
			  "strip-ansi": "^6.0.0",
			  "wcwidth": "^1.0.1"
			},
			"dependencies": {
			  "ansi-regex": {
				"version": "5.0.0",
				"resolved": "https://registry.npmjs.org/ansi-regex/-/ansi-regex-5.0.0.tgz",
				"integrity": "sha512-bY6fj56OUQ0hU1KjFNDQuJFezqKdrAyFdIevADiqrWHwSlbmBNMHp5ak2f40Pm8JTFyM2mqxkG6ngkHO11f/lg==",
				"dev": true
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "chalk": {
				"version": "3.0.0",
				"resolved": "https://registry.npmjs.org/chalk/-/chalk-3.0.0.tgz",
				"integrity": "sha512-4D3B6Wf41KOYRFdszmDqMCGq5VV/uMAB273JILmO+3jAlh8X4qDtdtgCR3fxtbLEMzSx22QdhnDcJvu2u1fVwg==",
				"dev": true,
				"requires": {
				  "ansi-styles": "^4.1.0",
				  "supports-color": "^7.1.0"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "strip-ansi": {
				"version": "6.0.0",
				"resolved": "https://registry.npmjs.org/strip-ansi/-/strip-ansi-6.0.0.tgz",
				"integrity": "sha512-AuvKTrTfQNYNIctbR1K/YGTR1756GycPsg7b9bdV9Duqur4gv6aKqHXah67Z8ImS7WEz5QVcOtlfW2rZEugt6w==",
				"dev": true,
				"requires": {
				  "ansi-regex": "^5.0.0"
				}
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "os-browserify": {
			"version": "0.3.0",
			"resolved": "https://registry.npmjs.org/os-browserify/-/os-browserify-0.3.0.tgz",
			"integrity": "sha1-hUNzx/XCMVkU/Jv8a9gjj92h7Cc=",
			"dev": true
		  },
		  "os-name": {
			"version": "3.1.0",
			"resolved": "https://registry.npmjs.org/os-name/-/os-name-3.1.0.tgz",
			"integrity": "sha512-h8L+8aNjNcMpo/mAIBPn5PXCM16iyPGjHNWo6U1YO8sJTMHtEtyczI6QJnLoplswm6goopQkqc7OAnjhWcugVg==",
			"dev": true,
			"requires": {
			  "macos-release": "^2.2.0",
			  "windows-release": "^3.1.0"
			}
		  },
		  "os-tmpdir": {
			"version": "1.0.2",
			"resolved": "https://registry.npmjs.org/os-tmpdir/-/os-tmpdir-1.0.2.tgz",
			"integrity": "sha1-u+Z0BseaqFxc/sdm/lc0VV36EnQ=",
			"dev": true
		  },
		  "p-each-series": {
			"version": "2.1.0",
			"resolved": "https://registry.npmjs.org/p-each-series/-/p-each-series-2.1.0.tgz",
			"integrity": "sha512-ZuRs1miPT4HrjFa+9fRfOFXxGJfORgelKV9f9nNOWw2gl6gVsRaVDOQP0+MI0G0wGKns1Yacsu0GjOFbTK0JFQ==",
			"dev": true
		  },
		  "p-finally": {
			"version": "1.0.0",
			"resolved": "https://registry.npmjs.org/p-finally/-/p-finally-1.0.0.tgz",
			"integrity": "sha1-P7z7FbiZpEEjs0ttzBi3JDNqLK4=",
			"dev": true
		  },
		  "p-limit": {
			"version": "2.3.0",
			"resolved": "https://registry.npmjs.org/p-limit/-/p-limit-2.3.0.tgz",
			"integrity": "sha512-//88mFWSJx8lxCzwdAABTJL2MyWB12+eIY7MDL2SqLmAkeKU9qxRvWuSyTjm3FUmpBEMuFfckAIqEaVGUDxb6w==",
			"dev": true,
			"requires": {
			  "p-try": "^2.0.0"
			}
		  },
		  "p-locate": {
			"version": "3.0.0",
			"resolved": "https://registry.npmjs.org/p-locate/-/p-locate-3.0.0.tgz",
			"integrity": "sha512-x+12w/To+4GFfgJhBEpiDcLozRJGegY+Ei7/z0tSLkMmxGZNybVMSfWj9aJn8Z5Fc7dBUNJOOVgPv2H7IwulSQ==",
			"dev": true,
			"requires": {
			  "p-limit": "^2.0.0"
			}
		  },
		  "p-try": {
			"version": "2.2.0",
			"resolved": "https://registry.npmjs.org/p-try/-/p-try-2.2.0.tgz",
			"integrity": "sha512-R4nPAVTAU0B9D35/Gk3uJf/7XYbQcyohSKdvAxIRSNghFl4e71hVoGnBNQz9cWaXxO2I10KTC+3jMdvvoKw6dQ==",
			"dev": true
		  },
		  "pako": {
			"version": "1.0.11",
			"resolved": "https://registry.npmjs.org/pako/-/pako-1.0.11.tgz",
			"integrity": "sha512-4hLB8Py4zZce5s4yd9XzopqwVv/yGNhV1Bl8NTmCq1763HeK2+EwVTv+leGeL13Dnh2wfbqowVPXCIO0z4taYw==",
			"dev": true
		  },
		  "parallel-transform": {
			"version": "1.2.0",
			"resolved": "https://registry.npmjs.org/parallel-transform/-/parallel-transform-1.2.0.tgz",
			"integrity": "sha512-P2vSmIu38uIlvdcU7fDkyrxj33gTUy/ABO5ZUbGowxNCopBq/OoD42bP4UmMrJoPyk4Uqf0mu3mtWBhHCZD8yg==",
			"dev": true,
			"requires": {
			  "cyclist": "^1.0.1",
			  "inherits": "^2.0.3",
			  "readable-stream": "^2.1.5"
			},
			"dependencies": {
			  "isarray": {
				"version": "1.0.0",
				"resolved": "https://registry.npmjs.org/isarray/-/isarray-1.0.0.tgz",
				"integrity": "sha1-u5NdSFgsuhaMBoNJV6VKPgcSTxE=",
				"dev": true
			  },
			  "readable-stream": {
				"version": "2.3.7",
				"resolved": "https://registry.npmjs.org/readable-stream/-/readable-stream-2.3.7.tgz",
				"integrity": "sha512-Ebho8K4jIbHAxnuxi7o42OrZgF/ZTNcsZj6nRKyUmkhLFq8CHItp/fy6hQZuZmP/n3yZ9VBUbp4zz/mX8hmYPw==",
				"dev": true,
				"requires": {
				  "core-util-is": "~1.0.0",
				  "inherits": "~2.0.3",
				  "isarray": "~1.0.0",
				  "process-nextick-args": "~2.0.0",
				  "safe-buffer": "~5.1.1",
				  "string_decoder": "~1.1.1",
				  "util-deprecate": "~1.0.1"
				}
			  },
			  "string_decoder": {
				"version": "1.1.1",
				"resolved": "https://registry.npmjs.org/string_decoder/-/string_decoder-1.1.1.tgz",
				"integrity": "sha512-n/ShnvDi6FHbbVfviro+WojiFzv+s8MPMHBczVePfUpDJLwoLT0ht1l4YwBCbi8pJAveEEdnkHyPyTP/mzRfwg==",
				"dev": true,
				"requires": {
				  "safe-buffer": "~5.1.0"
				}
			  }
			}
		  },
		  "parent-module": {
			"version": "1.0.1",
			"resolved": "https://registry.npmjs.org/parent-module/-/parent-module-1.0.1.tgz",
			"integrity": "sha512-GQ2EWRpQV8/o+Aw8YqtfZZPfNRWZYkbidE9k5rpl/hC3vtHHBfGm2Ifi6qWV+coDGkrUKZAxE3Lot5kcsRlh+g==",
			"dev": true,
			"requires": {
			  "callsites": "^3.0.0"
			}
		  },
		  "parse-asn1": {
			"version": "5.1.5",
			"resolved": "https://registry.npmjs.org/parse-asn1/-/parse-asn1-5.1.5.tgz",
			"integrity": "sha512-jkMYn1dcJqF6d5CpU689bq7w/b5ALS9ROVSpQDPrZsqqesUJii9qutvoT5ltGedNXMO2e16YUWIghG9KxaViTQ==",
			"dev": true,
			"requires": {
			  "asn1.js": "^4.0.0",
			  "browserify-aes": "^1.0.0",
			  "create-hash": "^1.1.0",
			  "evp_bytestokey": "^1.0.0",
			  "pbkdf2": "^3.0.3",
			  "safe-buffer": "^5.1.1"
			}
		  },
		  "parse-json": {
			"version": "2.2.0",
			"resolved": "https://registry.npmjs.org/parse-json/-/parse-json-2.2.0.tgz",
			"integrity": "sha1-9ID0BDTvgHQfhGkJn43qGPVaTck=",
			"dev": true,
			"requires": {
			  "error-ex": "^1.2.0"
			}
		  },
		  "parse5": {
			"version": "5.1.1",
			"resolved": "https://registry.npmjs.org/parse5/-/parse5-5.1.1.tgz",
			"integrity": "sha512-ugq4DFI0Ptb+WWjAdOK16+u/nHfiIrcE+sh8kZMaM0WllQKLI9rOUq6c2b7cwPkXdzfQESqvoqK6ug7U/Yyzug==",
			"dev": true
		  },
		  "parseurl": {
			"version": "1.3.3",
			"resolved": "https://registry.npmjs.org/parseurl/-/parseurl-1.3.3.tgz",
			"integrity": "sha512-CiyeOxFT/JZyN5m0z9PfXw4SCBJ6Sygz1Dpl0wqjlhDEGGBP1GnsUVEL0p63hoG1fcj3fHynXi9NYO4nWOL+qQ=="
		  },
		  "pascalcase": {
			"version": "0.1.1",
			"resolved": "https://registry.npmjs.org/pascalcase/-/pascalcase-0.1.1.tgz",
			"integrity": "sha1-s2PlXoAGym/iF4TS2yK9FdeRfxQ=",
			"dev": true
		  },
		  "path-browserify": {
			"version": "0.0.1",
			"resolved": "https://registry.npmjs.org/path-browserify/-/path-browserify-0.0.1.tgz",
			"integrity": "sha512-BapA40NHICOS+USX9SN4tyhq+A2RrN/Ws5F0Z5aMHDp98Fl86lX8Oti8B7uN93L4Ifv4fHOEA+pQw87gmMO/lQ==",
			"dev": true
		  },
		  "path-dirname": {
			"version": "1.0.2",
			"resolved": "https://registry.npmjs.org/path-dirname/-/path-dirname-1.0.2.tgz",
			"integrity": "sha1-zDPSTVJeCZpTiMAzbG4yuRYGCeA=",
			"dev": true,
			"optional": true
		  },
		  "path-exists": {
			"version": "3.0.0",
			"resolved": "https://registry.npmjs.org/path-exists/-/path-exists-3.0.0.tgz",
			"integrity": "sha1-zg6+ql94yxiSXqfYENe1mwEP1RU=",
			"dev": true
		  },
		  "path-is-absolute": {
			"version": "1.0.1",
			"resolved": "https://registry.npmjs.org/path-is-absolute/-/path-is-absolute-1.0.1.tgz",
			"integrity": "sha1-F0uSaHNVNP+8es5r9TpanhtcX18="
		  },
		  "path-key": {
			"version": "2.0.1",
			"resolved": "https://registry.npmjs.org/path-key/-/path-key-2.0.1.tgz",
			"integrity": "sha1-QRyttXTFoUDTpLGRDUDYDMn0C0A=",
			"dev": true
		  },
		  "path-parse": {
			"version": "1.0.6",
			"resolved": "https://registry.npmjs.org/path-parse/-/path-parse-1.0.6.tgz",
			"integrity": "sha512-GSmOT2EbHrINBf9SR7CDELwlJ8AENk3Qn7OikK4nFYAu3Ote2+JYNVvkpAEQm3/TLNEJFD/xZJjzyxg3KBWOzw==",
			"dev": true
		  },
		  "path-to-regexp": {
			"version": "3.2.0",
			"resolved": "https://registry.npmjs.org/path-to-regexp/-/path-to-regexp-3.2.0.tgz",
			"integrity": "sha512-jczvQbCUS7XmS7o+y1aEO9OBVFeZBQ1MDSEqmO7xSoPgOPoowY/SxLpZ6Vh97/8qHZOteiCKb7gkG9gA2ZUxJA=="
		  },
		  "path-type": {
			"version": "2.0.0",
			"resolved": "https://registry.npmjs.org/path-type/-/path-type-2.0.0.tgz",
			"integrity": "sha1-8BLMuEFbcJb8LaoQVMPXI4lZTHM=",
			"dev": true,
			"requires": {
			  "pify": "^2.0.0"
			},
			"dependencies": {
			  "pify": {
				"version": "2.3.0",
				"resolved": "https://registry.npmjs.org/pify/-/pify-2.3.0.tgz",
				"integrity": "sha1-7RQaasBDqEnqWISY59yosVMw6Qw=",
				"dev": true
			  }
			}
		  },
		  "pbkdf2": {
			"version": "3.0.17",
			"resolved": "https://registry.npmjs.org/pbkdf2/-/pbkdf2-3.0.17.tgz",
			"integrity": "sha512-U/il5MsrZp7mGg3mSQfn742na2T+1/vHDCG5/iTI3X9MKUuYUZVLQhyRsg06mCgDBTd57TxzgZt7P+fYfjRLtA==",
			"dev": true,
			"requires": {
			  "create-hash": "^1.1.2",
			  "create-hmac": "^1.1.4",
			  "ripemd160": "^2.0.1",
			  "safe-buffer": "^5.0.1",
			  "sha.js": "^2.4.8"
			}
		  },
		  "performance-now": {
			"version": "2.1.0",
			"resolved": "https://registry.npmjs.org/performance-now/-/performance-now-2.1.0.tgz",
			"integrity": "sha1-Ywn04OX6kT7BxpMHrjZLSzd8nns=",
			"dev": true
		  },
		  "picomatch": {
			"version": "2.2.2",
			"resolved": "https://registry.npmjs.org/picomatch/-/picomatch-2.2.2.tgz",
			"integrity": "sha512-q0M/9eZHzmr0AulXyPwNfZjtwZ/RBZlbN3K3CErVrk50T2ASYI7Bye0EvekFY3IP1Nt2DHu0re+V2ZHIpMkuWg==",
			"dev": true
		  },
		  "pify": {
			"version": "4.0.1",
			"resolved": "https://registry.npmjs.org/pify/-/pify-4.0.1.tgz",
			"integrity": "sha512-uB80kBFb/tfd68bVleG9T5GGsGPjJrLAUpR5PZIrhBnIaRTQRjqdJSsIKkOP6OAIFbj7GOrcudc5pNjZ+geV2g==",
			"dev": true
		  },
		  "pirates": {
			"version": "4.0.1",
			"resolved": "https://registry.npmjs.org/pirates/-/pirates-4.0.1.tgz",
			"integrity": "sha512-WuNqLTbMI3tmfef2TKxlQmAiLHKtFhlsCZnPIpuv2Ow0RDVO8lfy1Opf4NUzlMXLjPl+Men7AuVdX6TA+s+uGA==",
			"dev": true,
			"requires": {
			  "node-modules-regexp": "^1.0.0"
			}
		  },
		  "pkg-dir": {
			"version": "3.0.0",
			"resolved": "https://registry.npmjs.org/pkg-dir/-/pkg-dir-3.0.0.tgz",
			"integrity": "sha512-/E57AYkoeQ25qkxMj5PBOVgF8Kiu/h7cYS30Z5+R7WaiCCBfLq58ZI/dSeaEKb9WVJV5n/03QwrN3IeWIFllvw==",
			"dev": true,
			"requires": {
			  "find-up": "^3.0.0"
			}
		  },
		  "posix-character-classes": {
			"version": "0.1.1",
			"resolved": "https://registry.npmjs.org/posix-character-classes/-/posix-character-classes-0.1.1.tgz",
			"integrity": "sha1-AerA/jta9xoqbAL+q7jB/vfgDqs=",
			"dev": true
		  },
		  "prelude-ls": {
			"version": "1.2.1",
			"resolved": "https://registry.npmjs.org/prelude-ls/-/prelude-ls-1.2.1.tgz",
			"integrity": "sha512-vkcDPrRZo1QZLbn5RLGPpg/WmIQ65qoWWhcGKf/b5eplkkarX0m9z8ppCat4mlOqUsWpyNuYgO3VRyrYHSzX5g==",
			"dev": true
		  },
		  "prettier": {
			"version": "2.0.5",
			"resolved": "https://registry.npmjs.org/prettier/-/prettier-2.0.5.tgz",
			"integrity": "sha512-7PtVymN48hGcO4fGjybyBSIWDsLU4H4XlvOHfq91pz9kkGlonzwTfYkaIEwiRg/dAJF9YlbsduBAgtYLi+8cFg==",
			"dev": true
		  },
		  "pretty-format": {
			"version": "25.5.0",
			"resolved": "https://registry.npmjs.org/pretty-format/-/pretty-format-25.5.0.tgz",
			"integrity": "sha512-kbo/kq2LQ/A/is0PQwsEHM7Ca6//bGPPvU6UnsdDRSKTWxT/ru/xb88v4BJf6a69H+uTytOEsTusT9ksd/1iWQ==",
			"dev": true,
			"requires": {
			  "@jest/types": "^25.5.0",
			  "ansi-regex": "^5.0.0",
			  "ansi-styles": "^4.0.0",
			  "react-is": "^16.12.0"
			},
			"dependencies": {
			  "ansi-regex": {
				"version": "5.0.0",
				"resolved": "https://registry.npmjs.org/ansi-regex/-/ansi-regex-5.0.0.tgz",
				"integrity": "sha512-bY6fj56OUQ0hU1KjFNDQuJFezqKdrAyFdIevADiqrWHwSlbmBNMHp5ak2f40Pm8JTFyM2mqxkG6ngkHO11f/lg==",
				"dev": true
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  }
			}
		  },
		  "process": {
			"version": "0.11.10",
			"resolved": "https://registry.npmjs.org/process/-/process-0.11.10.tgz",
			"integrity": "sha1-czIwDoQBYb2j5podHZGn1LwW8YI=",
			"dev": true
		  },
		  "process-nextick-args": {
			"version": "2.0.1",
			"resolved": "https://registry.npmjs.org/process-nextick-args/-/process-nextick-args-2.0.1.tgz",
			"integrity": "sha512-3ouUOpQhtgrbOa17J7+uxOTpITYWaGP7/AhoR3+A+/1e9skrzelGi/dXzEYyvbxubEF6Wn2ypscTKiKJFFn1ag=="
		  },
		  "progress": {
			"version": "2.0.3",
			"resolved": "https://registry.npmjs.org/progress/-/progress-2.0.3.tgz",
			"integrity": "sha512-7PiHtLll5LdnKIMw100I+8xJXR5gW2QwWYkT6iJva0bXitZKa/XMrSbdmg3r2Xnaidz9Qumd0VPaMrZlF9V9sA==",
			"dev": true
		  },
		  "promise-inflight": {
			"version": "1.0.1",
			"resolved": "https://registry.npmjs.org/promise-inflight/-/promise-inflight-1.0.1.tgz",
			"integrity": "sha1-mEcocL8igTL8vdhoEputEsPAKeM=",
			"dev": true
		  },
		  "prompts": {
			"version": "2.3.2",
			"resolved": "https://registry.npmjs.org/prompts/-/prompts-2.3.2.tgz",
			"integrity": "sha512-Q06uKs2CkNYVID0VqwfAl9mipo99zkBv/n2JtWY89Yxa3ZabWSrs0e2KTudKVa3peLUvYXMefDqIleLPVUBZMA==",
			"dev": true,
			"requires": {
			  "kleur": "^3.0.3",
			  "sisteransi": "^1.0.4"
			}
		  },
		  "proxy-addr": {
			"version": "2.0.6",
			"resolved": "https://registry.npmjs.org/proxy-addr/-/proxy-addr-2.0.6.tgz",
			"integrity": "sha512-dh/frvCBVmSsDYzw6n926jv974gddhkFPfiN8hPOi30Wax25QZyZEGveluCgliBnqmuM+UJmBErbAUFIoDbjOw==",
			"requires": {
			  "forwarded": "~0.1.2",
			  "ipaddr.js": "1.9.1"
			}
		  },
		  "prr": {
			"version": "1.0.1",
			"resolved": "https://registry.npmjs.org/prr/-/prr-1.0.1.tgz",
			"integrity": "sha1-0/wRS6BplaRexok/SEzrHXj19HY=",
			"dev": true
		  },
		  "psl": {
			"version": "1.8.0",
			"resolved": "https://registry.npmjs.org/psl/-/psl-1.8.0.tgz",
			"integrity": "sha512-RIdOzyoavK+hA18OGGWDqUTsCLhtA7IcZ/6NCs4fFJaHBDab+pDDmDIByWFRQJq2Cd7r1OoQxBGKOaztq+hjIQ==",
			"dev": true
		  },
		  "public-encrypt": {
			"version": "4.0.3",
			"resolved": "https://registry.npmjs.org/public-encrypt/-/public-encrypt-4.0.3.tgz",
			"integrity": "sha512-zVpa8oKZSz5bTMTFClc1fQOnyyEzpl5ozpi1B5YcvBrdohMjH2rfsBtyXcuNuwjsDIXmBYlF2N5FlJYhR29t8Q==",
			"dev": true,
			"requires": {
			  "bn.js": "^4.1.0",
			  "browserify-rsa": "^4.0.0",
			  "create-hash": "^1.1.0",
			  "parse-asn1": "^5.0.0",
			  "randombytes": "^2.0.1",
			  "safe-buffer": "^5.1.2"
			},
			"dependencies": {
			  "bn.js": {
				"version": "4.11.9",
				"resolved": "https://registry.npmjs.org/bn.js/-/bn.js-4.11.9.tgz",
				"integrity": "sha512-E6QoYqCKZfgatHTdHzs1RRKP7ip4vvm+EyRUeE2RF0NblwVvb0p6jSVeNTOFxPn26QXN2o6SMfNxKp6kU8zQaw==",
				"dev": true
			  }
			}
		  },
		  "pump": {
			"version": "3.0.0",
			"resolved": "https://registry.npmjs.org/pump/-/pump-3.0.0.tgz",
			"integrity": "sha512-LwZy+p3SFs1Pytd/jYct4wpv49HiYCqd9Rlc5ZVdk0V+8Yzv6jR5Blk3TRmPL1ft69TxP0IMZGJ+WPFU2BFhww==",
			"dev": true,
			"requires": {
			  "end-of-stream": "^1.1.0",
			  "once": "^1.3.1"
			}
		  },
		  "pumpify": {
			"version": "1.5.1",
			"resolved": "https://registry.npmjs.org/pumpify/-/pumpify-1.5.1.tgz",
			"integrity": "sha512-oClZI37HvuUJJxSKKrC17bZ9Cu0ZYhEAGPsPUy9KlMUmv9dKX2o77RUmq7f3XjIxbwyGwYzbzQ1L2Ks8sIradQ==",
			"dev": true,
			"requires": {
			  "duplexify": "^3.6.0",
			  "inherits": "^2.0.3",
			  "pump": "^2.0.0"
			},
			"dependencies": {
			  "pump": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/pump/-/pump-2.0.1.tgz",
				"integrity": "sha512-ruPMNRkN3MHP1cWJc9OWr+T/xDP0jhXYCLfJcBuX54hhfIBnaQmAUMfDcG4DM5UMWByBbJY69QSphm3jtDKIkA==",
				"dev": true,
				"requires": {
				  "end-of-stream": "^1.1.0",
				  "once": "^1.3.1"
				}
			  }
			}
		  },
		  "punycode": {
			"version": "2.1.1",
			"resolved": "https://registry.npmjs.org/punycode/-/punycode-2.1.1.tgz",
			"integrity": "sha512-XRsRjdf+j5ml+y/6GKHPZbrF/8p2Yga0JPtdqTIY2Xe5ohJPD9saDJJLPvp9+NSBprVvevdXZybnj2cv8OEd0A==",
			"dev": true
		  },
		  "qs": {
			"version": "6.7.0",
			"resolved": "https://registry.npmjs.org/qs/-/qs-6.7.0.tgz",
			"integrity": "sha512-VCdBRNFTX1fyE7Nb6FYoURo/SPe62QCaAyzJvUjwRaIsc+NePBEniHlvxFmmX56+HZphIGtV0XeCirBtpDrTyQ=="
		  },
		  "querystring": {
			"version": "0.2.0",
			"resolved": "https://registry.npmjs.org/querystring/-/querystring-0.2.0.tgz",
			"integrity": "sha1-sgmEkgO7Jd+CDadW50cAWHhSFiA=",
			"dev": true
		  },
		  "querystring-es3": {
			"version": "0.2.1",
			"resolved": "https://registry.npmjs.org/querystring-es3/-/querystring-es3-0.2.1.tgz",
			"integrity": "sha1-nsYfeQSYdXB9aUFFlv2Qek1xHnM=",
			"dev": true
		  },
		  "randombytes": {
			"version": "2.1.0",
			"resolved": "https://registry.npmjs.org/randombytes/-/randombytes-2.1.0.tgz",
			"integrity": "sha512-vYl3iOX+4CKUWuxGi9Ukhie6fsqXqS9FE2Zaic4tNFD2N2QQaXOMFbuKK4QmDHC0JO6B1Zp41J0LpT0oR68amQ==",
			"dev": true,
			"requires": {
			  "safe-buffer": "^5.1.0"
			}
		  },
		  "randomfill": {
			"version": "1.0.4",
			"resolved": "https://registry.npmjs.org/randomfill/-/randomfill-1.0.4.tgz",
			"integrity": "sha512-87lcbR8+MhcWcUiQ+9e+Rwx8MyR2P7qnt15ynUlbm3TU/fjbgz4GsvfSUDTemtCCtVCqb4ZcEFlyPNTh9bBTLw==",
			"dev": true,
			"requires": {
			  "randombytes": "^2.0.5",
			  "safe-buffer": "^5.1.0"
			}
		  },
		  "range-parser": {
			"version": "1.2.1",
			"resolved": "https://registry.npmjs.org/range-parser/-/range-parser-1.2.1.tgz",
			"integrity": "sha512-Hrgsx+orqoygnmhFbKaHE6c296J+HTAQXoxEF6gNupROmmGJRoyzfG3ccAveqCBrwr/2yxQ5BVd/GTl5agOwSg=="
		  },
		  "raw-body": {
			"version": "2.4.0",
			"resolved": "https://registry.npmjs.org/raw-body/-/raw-body-2.4.0.tgz",
			"integrity": "sha512-4Oz8DUIwdvoa5qMJelxipzi/iJIi40O5cGV1wNYp5hvZP8ZN0T+jiNkL0QepXs+EsQ9XJ8ipEDoiH70ySUJP3Q==",
			"requires": {
			  "bytes": "3.1.0",
			  "http-errors": "1.7.2",
			  "iconv-lite": "0.4.24",
			  "unpipe": "1.0.0"
			}
		  },
		  "react-is": {
			"version": "16.13.1",
			"resolved": "https://registry.npmjs.org/react-is/-/react-is-16.13.1.tgz",
			"integrity": "sha512-24e6ynE2H+OKt4kqsOvNd8kBpV65zoxbA4BVsEOB3ARVWQki/DHzaUoC5KuON/BiccDaCCTZBuOcfZs70kR8bQ==",
			"dev": true
		  },
		  "read-pkg": {
			"version": "2.0.0",
			"resolved": "https://registry.npmjs.org/read-pkg/-/read-pkg-2.0.0.tgz",
			"integrity": "sha1-jvHAYjxqbbDcZxPEv6xGMysjaPg=",
			"dev": true,
			"requires": {
			  "load-json-file": "^2.0.0",
			  "normalize-package-data": "^2.3.2",
			  "path-type": "^2.0.0"
			}
		  },
		  "read-pkg-up": {
			"version": "2.0.0",
			"resolved": "https://registry.npmjs.org/read-pkg-up/-/read-pkg-up-2.0.0.tgz",
			"integrity": "sha1-a3KoBImE4MQeeVEP1en6mbO1Sb4=",
			"dev": true,
			"requires": {
			  "find-up": "^2.0.0",
			  "read-pkg": "^2.0.0"
			},
			"dependencies": {
			  "find-up": {
				"version": "2.1.0",
				"resolved": "https://registry.npmjs.org/find-up/-/find-up-2.1.0.tgz",
				"integrity": "sha1-RdG35QbHF93UgndaK3eSCjwMV6c=",
				"dev": true,
				"requires": {
				  "locate-path": "^2.0.0"
				}
			  },
			  "locate-path": {
				"version": "2.0.0",
				"resolved": "https://registry.npmjs.org/locate-path/-/locate-path-2.0.0.tgz",
				"integrity": "sha1-K1aLJl7slExtnA3pw9u7ygNUzY4=",
				"dev": true,
				"requires": {
				  "p-locate": "^2.0.0",
				  "path-exists": "^3.0.0"
				}
			  },
			  "p-limit": {
				"version": "1.3.0",
				"resolved": "https://registry.npmjs.org/p-limit/-/p-limit-1.3.0.tgz",
				"integrity": "sha512-vvcXsLAJ9Dr5rQOPk7toZQZJApBl2K4J6dANSsEuh6QI41JYcsS/qhTGa9ErIUUgK3WNQoJYvylxvjqmiqEA9Q==",
				"dev": true,
				"requires": {
				  "p-try": "^1.0.0"
				}
			  },
			  "p-locate": {
				"version": "2.0.0",
				"resolved": "https://registry.npmjs.org/p-locate/-/p-locate-2.0.0.tgz",
				"integrity": "sha1-IKAQOyIqcMj9OcwuWAaA893l7EM=",
				"dev": true,
				"requires": {
				  "p-limit": "^1.1.0"
				}
			  },
			  "p-try": {
				"version": "1.0.0",
				"resolved": "https://registry.npmjs.org/p-try/-/p-try-1.0.0.tgz",
				"integrity": "sha1-y8ec26+P1CKOE/Yh8rGiN8GyB7M=",
				"dev": true
			  }
			}
		  },
		  "readable-stream": {
			"version": "1.1.14",
			"resolved": "https://registry.npmjs.org/readable-stream/-/readable-stream-1.1.14.tgz",
			"integrity": "sha1-fPTFTvZI44EwhMY23SB54WbAgdk=",
			"requires": {
			  "core-util-is": "~1.0.0",
			  "inherits": "~2.0.1",
			  "isarray": "0.0.1",
			  "string_decoder": "~0.10.x"
			}
		  },
		  "readdirp": {
			"version": "3.4.0",
			"resolved": "https://registry.npmjs.org/readdirp/-/readdirp-3.4.0.tgz",
			"integrity": "sha512-0xe001vZBnJEK+uKcj8qOhyAKPzIT+gStxWr3LCB0DwcXR5NZJ3IaC+yGnHCYzB/S7ov3m3EEbZI2zeNvX+hGQ==",
			"dev": true,
			"requires": {
			  "picomatch": "^2.2.1"
			}
		  },
		  "rechoir": {
			"version": "0.6.2",
			"resolved": "https://registry.npmjs.org/rechoir/-/rechoir-0.6.2.tgz",
			"integrity": "sha1-hSBLVNuoLVdC4oyWdW70OvUOM4Q=",
			"dev": true,
			"requires": {
			  "resolve": "^1.1.6"
			}
		  },
		  "reflect-metadata": {
			"version": "0.1.13",
			"resolved": "https://registry.npmjs.org/reflect-metadata/-/reflect-metadata-0.1.13.tgz",
			"integrity": "sha512-Ts1Y/anZELhSsjMcU605fU9RE4Oi3p5ORujwbIKXfWa+0Zxs510Qrmrce5/Jowq3cHSZSJqBjypxmHarc+vEWg=="
		  },
		  "regex-not": {
			"version": "1.0.2",
			"resolved": "https://registry.npmjs.org/regex-not/-/regex-not-1.0.2.tgz",
			"integrity": "sha512-J6SDjUgDxQj5NusnOtdFxDwN/+HWykR8GELwctJ7mdqhcyy1xEc4SRFHUXvxTp661YaVKAjfRLZ9cCqS6tn32A==",
			"dev": true,
			"requires": {
			  "extend-shallow": "^3.0.2",
			  "safe-regex": "^1.1.0"
			}
		  },
		  "regexpp": {
			"version": "3.1.0",
			"resolved": "https://registry.npmjs.org/regexpp/-/regexpp-3.1.0.tgz",
			"integrity": "sha512-ZOIzd8yVsQQA7j8GCSlPGXwg5PfmA1mrq0JP4nGhh54LaKN3xdai/vHUDu74pKwV8OxseMS65u2NImosQcSD0Q==",
			"dev": true
		  },
		  "remove-trailing-separator": {
			"version": "1.1.0",
			"resolved": "https://registry.npmjs.org/remove-trailing-separator/-/remove-trailing-separator-1.1.0.tgz",
			"integrity": "sha1-wkvOKig62tW8P1jg1IJJuSN52O8=",
			"dev": true
		  },
		  "repeat-element": {
			"version": "1.1.3",
			"resolved": "https://registry.npmjs.org/repeat-element/-/repeat-element-1.1.3.tgz",
			"integrity": "sha512-ahGq0ZnV5m5XtZLMb+vP76kcAM5nkLqk0lpqAuojSKGgQtn4eRi4ZZGm2olo2zKFH+sMsWaqOCW1dqAnOru72g==",
			"dev": true
		  },
		  "repeat-string": {
			"version": "1.6.1",
			"resolved": "https://registry.npmjs.org/repeat-string/-/repeat-string-1.6.1.tgz",
			"integrity": "sha1-jcrkcOHIirwtYA//Sndihtp15jc=",
			"dev": true
		  },
		  "request": {
			"version": "2.88.2",
			"resolved": "https://registry.npmjs.org/request/-/request-2.88.2.tgz",
			"integrity": "sha512-MsvtOrfG9ZcrOwAW+Qi+F6HbD0CWXEh9ou77uOb7FM2WPhwT7smM833PzanhJLsgXjN89Ir6V2PczXNnMpwKhw==",
			"dev": true,
			"requires": {
			  "aws-sign2": "~0.7.0",
			  "aws4": "^1.8.0",
			  "caseless": "~0.12.0",
			  "combined-stream": "~1.0.6",
			  "extend": "~3.0.2",
			  "forever-agent": "~0.6.1",
			  "form-data": "~2.3.2",
			  "har-validator": "~5.1.3",
			  "http-signature": "~1.2.0",
			  "is-typedarray": "~1.0.0",
			  "isstream": "~0.1.2",
			  "json-stringify-safe": "~5.0.1",
			  "mime-types": "~2.1.19",
			  "oauth-sign": "~0.9.0",
			  "performance-now": "^2.1.0",
			  "qs": "~6.5.2",
			  "safe-buffer": "^5.1.2",
			  "tough-cookie": "~2.5.0",
			  "tunnel-agent": "^0.6.0",
			  "uuid": "^3.3.2"
			},
			"dependencies": {
			  "qs": {
				"version": "6.5.2",
				"resolved": "https://registry.npmjs.org/qs/-/qs-6.5.2.tgz",
				"integrity": "sha512-N5ZAX4/LxJmF+7wN74pUD6qAh9/wnvdQcjq9TZjevvXzSUo7bfmw91saqMjzGS2xq91/odN2dW/WOl7qQHNDGA==",
				"dev": true
			  },
			  "tough-cookie": {
				"version": "2.5.0",
				"resolved": "https://registry.npmjs.org/tough-cookie/-/tough-cookie-2.5.0.tgz",
				"integrity": "sha512-nlLsUzgm1kfLXSXfRZMc1KLAugd4hqJHDTvc2hDIwS3mZAfMEuMbc03SujMF+GEcpaX/qboeycw6iO8JwVv2+g==",
				"dev": true,
				"requires": {
				  "psl": "^1.1.28",
				  "punycode": "^2.1.1"
				}
			  },
			  "uuid": {
				"version": "3.4.0",
				"resolved": "https://registry.npmjs.org/uuid/-/uuid-3.4.0.tgz",
				"integrity": "sha512-HjSDRw6gZE5JMggctHBcjVak08+KEVhSIiDzFnT9S9aegmp85S/bReBVTb4QTFaRNptJ9kuYaNhnbNEOkbKb/A==",
				"dev": true
			  }
			}
		  },
		  "request-promise-core": {
			"version": "1.1.3",
			"resolved": "https://registry.npmjs.org/request-promise-core/-/request-promise-core-1.1.3.tgz",
			"integrity": "sha512-QIs2+ArIGQVp5ZYbWD5ZLCY29D5CfWizP8eWnm8FoGD1TX61veauETVQbrV60662V0oFBkrDOuaBI8XgtuyYAQ==",
			"dev": true,
			"requires": {
			  "lodash": "^4.17.15"
			}
		  },
		  "request-promise-native": {
			"version": "1.0.8",
			"resolved": "https://registry.npmjs.org/request-promise-native/-/request-promise-native-1.0.8.tgz",
			"integrity": "sha512-dapwLGqkHtwL5AEbfenuzjTYg35Jd6KPytsC2/TLkVMz8rm+tNt72MGUWT1RP/aYawMpN6HqbNGBQaRcBtjQMQ==",
			"dev": true,
			"requires": {
			  "request-promise-core": "1.1.3",
			  "stealthy-require": "^1.1.1",
			  "tough-cookie": "^2.3.3"
			},
			"dependencies": {
			  "tough-cookie": {
				"version": "2.5.0",
				"resolved": "https://registry.npmjs.org/tough-cookie/-/tough-cookie-2.5.0.tgz",
				"integrity": "sha512-nlLsUzgm1kfLXSXfRZMc1KLAugd4hqJHDTvc2hDIwS3mZAfMEuMbc03SujMF+GEcpaX/qboeycw6iO8JwVv2+g==",
				"dev": true,
				"requires": {
				  "psl": "^1.1.28",
				  "punycode": "^2.1.1"
				}
			  }
			}
		  },
		  "require-directory": {
			"version": "2.1.1",
			"resolved": "https://registry.npmjs.org/require-directory/-/require-directory-2.1.1.tgz",
			"integrity": "sha1-jGStX9MNqxyXbiNE/+f3kqam30I=",
			"dev": true
		  },
		  "require-main-filename": {
			"version": "2.0.0",
			"resolved": "https://registry.npmjs.org/require-main-filename/-/require-main-filename-2.0.0.tgz",
			"integrity": "sha512-NKN5kMDylKuldxYLSUfrbo5Tuzh4hd+2E8NPPX02mZtn1VuREQToYe/ZdlJy+J3uCpfaiGF05e7B8W0iXbQHmg==",
			"dev": true
		  },
		  "resolve": {
			"version": "1.17.0",
			"resolved": "https://registry.npmjs.org/resolve/-/resolve-1.17.0.tgz",
			"integrity": "sha512-ic+7JYiV8Vi2yzQGFWOkiZD5Z9z7O2Zhm9XMaTxdJExKasieFCr+yXZ/WmXsckHiKl12ar0y6XiXDx3m4RHn1w==",
			"dev": true,
			"requires": {
			  "path-parse": "^1.0.6"
			}
		  },
		  "resolve-cwd": {
			"version": "3.0.0",
			"resolved": "https://registry.npmjs.org/resolve-cwd/-/resolve-cwd-3.0.0.tgz",
			"integrity": "sha512-OrZaX2Mb+rJCpH/6CpSqt9xFVpN++x01XnN2ie9g6P5/3xelLAkXWVADpdz1IHD/KFfEXyE6V0U01OQ3UO2rEg==",
			"dev": true,
			"requires": {
			  "resolve-from": "^5.0.0"
			},
			"dependencies": {
			  "resolve-from": {
				"version": "5.0.0",
				"resolved": "https://registry.npmjs.org/resolve-from/-/resolve-from-5.0.0.tgz",
				"integrity": "sha512-qYg9KP24dD5qka9J47d0aVky0N+b4fTU89LN9iDnjB5waksiC49rvMB0PrUJQGoTmH50XPiqOvAjDfaijGxYZw==",
				"dev": true
			  }
			}
		  },
		  "resolve-from": {
			"version": "4.0.0",
			"resolved": "https://registry.npmjs.org/resolve-from/-/resolve-from-4.0.0.tgz",
			"integrity": "sha512-pb/MYmXstAkysRFx8piNI1tGFNQIFA3vkE3Gq4EuA1dF6gHp/+vgZqsCGJapvy8N3Q+4o7FwvquPJcnZ7RYy4g==",
			"dev": true
		  },
		  "resolve-url": {
			"version": "0.2.1",
			"resolved": "https://registry.npmjs.org/resolve-url/-/resolve-url-0.2.1.tgz",
			"integrity": "sha1-LGN/53yJOv0qZj/iGqkIAGjiBSo=",
			"dev": true
		  },
		  "restore-cursor": {
			"version": "3.1.0",
			"resolved": "https://registry.npmjs.org/restore-cursor/-/restore-cursor-3.1.0.tgz",
			"integrity": "sha512-l+sSefzHpj5qimhFSE5a8nufZYAM3sBSVMAPtYkmC+4EH2anSGaEMXSD0izRQbu9nfyQ9y5JrVmp7E8oZrUjvA==",
			"dev": true,
			"requires": {
			  "onetime": "^5.1.0",
			  "signal-exit": "^3.0.2"
			}
		  },
		  "ret": {
			"version": "0.1.15",
			"resolved": "https://registry.npmjs.org/ret/-/ret-0.1.15.tgz",
			"integrity": "sha512-TTlYpa+OL+vMMNG24xSlQGEJ3B/RzEfUlLct7b5G/ytav+wPrplCpVMFuwzXbkecJrb6IYo1iFb0S9v37754mg==",
			"dev": true
		  },
		  "rimraf": {
			"version": "3.0.2",
			"resolved": "https://registry.npmjs.org/rimraf/-/rimraf-3.0.2.tgz",
			"integrity": "sha512-JZkJMZkAGFFPP2YqXZXPbMlMBgsxzE8ILs4lMIX/2o0L9UBw9O/Y3o6wFw/i9YLapcUJWwqbi3kdxIPdC62TIA==",
			"requires": {
			  "glob": "^7.1.3"
			}
		  },
		  "ripemd160": {
			"version": "2.0.2",
			"resolved": "https://registry.npmjs.org/ripemd160/-/ripemd160-2.0.2.tgz",
			"integrity": "sha512-ii4iagi25WusVoiC4B4lq7pbXfAp3D9v5CwfkY33vffw2+pkDjY1D8GaN7spsxvCSx8dkPqOZCEZyfxcmJG2IA==",
			"dev": true,
			"requires": {
			  "hash-base": "^3.0.0",
			  "inherits": "^2.0.1"
			}
		  },
		  "rsvp": {
			"version": "4.8.5",
			"resolved": "https://registry.npmjs.org/rsvp/-/rsvp-4.8.5.tgz",
			"integrity": "sha512-nfMOlASu9OnRJo1mbEk2cz0D56a1MBNrJ7orjRZQG10XDyuvwksKbuXNp6qa+kbn839HwjwhBzhFmdsaEAfauA==",
			"dev": true
		  },
		  "run-async": {
			"version": "2.4.1",
			"resolved": "https://registry.npmjs.org/run-async/-/run-async-2.4.1.tgz",
			"integrity": "sha512-tvVnVv01b8c1RrA6Ep7JkStj85Guv/YrMcwqYQnwjsAS2cTmmPGBBjAjpCW7RrSodNSoE2/qg9O4bceNvUuDgQ==",
			"dev": true
		  },
		  "run-queue": {
			"version": "1.0.3",
			"resolved": "https://registry.npmjs.org/run-queue/-/run-queue-1.0.3.tgz",
			"integrity": "sha1-6Eg5bwV9Ij8kOGkkYY4laUFh7Ec=",
			"dev": true,
			"requires": {
			  "aproba": "^1.1.1"
			}
		  },
		  "rxjs": {
			"version": "6.5.5",
			"resolved": "https://registry.npmjs.org/rxjs/-/rxjs-6.5.5.tgz",
			"integrity": "sha512-WfQI+1gohdf0Dai/Bbmk5L5ItH5tYqm3ki2c5GdWhKjalzjg93N3avFjVStyZZz+A2Em+ZxKH5bNghw9UeylGQ==",
			"requires": {
			  "tslib": "^1.9.0"
			}
		  },
		  "safe-buffer": {
			"version": "5.1.2",
			"resolved": "https://registry.npmjs.org/safe-buffer/-/safe-buffer-5.1.2.tgz",
			"integrity": "sha512-Gd2UZBJDkXlY7GbJxfsE8/nvKkUEU1G38c1siN6QP6a9PT9MmHB8GnpscSmMJSoF8LOIrt8ud/wPtojys4G6+g=="
		  },
		  "safe-regex": {
			"version": "1.1.0",
			"resolved": "https://registry.npmjs.org/safe-regex/-/safe-regex-1.1.0.tgz",
			"integrity": "sha1-QKNmnzsHfR6UPURinhV91IAjvy4=",
			"dev": true,
			"requires": {
			  "ret": "~0.1.10"
			}
		  },
		  "safer-buffer": {
			"version": "2.1.2",
			"resolved": "https://registry.npmjs.org/safer-buffer/-/safer-buffer-2.1.2.tgz",
			"integrity": "sha512-YZo3K82SD7Riyi0E1EQPojLz7kpepnSQI9IyPbHHg1XXXevb5dJI7tpyN2ADxGcQbHG7vcyRHk0cbwqcQriUtg=="
		  },
		  "sane": {
			"version": "4.1.0",
			"resolved": "https://registry.npmjs.org/sane/-/sane-4.1.0.tgz",
			"integrity": "sha512-hhbzAgTIX8O7SHfp2c8/kREfEn4qO/9q8C9beyY6+tvZ87EpoZ3i1RIEvp27YBswnNbY9mWd6paKVmKbAgLfZA==",
			"dev": true,
			"requires": {
			  "@cnakazawa/watch": "^1.0.3",
			  "anymatch": "^2.0.0",
			  "capture-exit": "^2.0.0",
			  "exec-sh": "^0.3.2",
			  "execa": "^1.0.0",
			  "fb-watchman": "^2.0.0",
			  "micromatch": "^3.1.4",
			  "minimist": "^1.1.1",
			  "walker": "~1.0.5"
			},
			"dependencies": {
			  "anymatch": {
				"version": "2.0.0",
				"resolved": "https://registry.npmjs.org/anymatch/-/anymatch-2.0.0.tgz",
				"integrity": "sha512-5teOsQWABXHHBFP9y3skS5P3d/WfWXpv3FUpy+LorMrNYaT9pI4oLMQX7jzQ2KklNpGpWHzdCXTDT2Y3XGlZBw==",
				"dev": true,
				"requires": {
				  "micromatch": "^3.1.4",
				  "normalize-path": "^2.1.1"
				}
			  },
			  "normalize-path": {
				"version": "2.1.1",
				"resolved": "https://registry.npmjs.org/normalize-path/-/normalize-path-2.1.1.tgz",
				"integrity": "sha1-GrKLVW4Zg2Oowab35vogE3/mrtk=",
				"dev": true,
				"requires": {
				  "remove-trailing-separator": "^1.0.1"
				}
			  }
			}
		  },
		  "saxes": {
			"version": "5.0.1",
			"resolved": "https://registry.npmjs.org/saxes/-/saxes-5.0.1.tgz",
			"integrity": "sha512-5LBh1Tls8c9xgGjw3QrMwETmTMVk0oFgvrFSvWx62llR2hcEInrKNZ2GZCCuuy2lvWrdl5jhbpeqc5hRYKFOcw==",
			"dev": true,
			"requires": {
			  "xmlchars": "^2.2.0"
			}
		  },
		  "schema-utils": {
			"version": "1.0.0",
			"resolved": "https://registry.npmjs.org/schema-utils/-/schema-utils-1.0.0.tgz",
			"integrity": "sha512-i27Mic4KovM/lnGsy8whRCHhc7VicJajAjTrYg11K9zfZXnYIt4k5F+kZkwjnrhKzLic/HLU4j11mjsz2G/75g==",
			"dev": true,
			"requires": {
			  "ajv": "^6.1.0",
			  "ajv-errors": "^1.0.0",
			  "ajv-keywords": "^3.1.0"
			}
		  },
		  "semver": {
			"version": "5.7.1",
			"resolved": "https://registry.npmjs.org/semver/-/semver-5.7.1.tgz",
			"integrity": "sha512-sauaDf/PZdVgrLTNYHRtpXa1iRiKcaebiKQ1BJdpQlWH2lCvexQdX55snPFyK7QzpudqbCI0qXFfOasHdyNDGQ==",
			"dev": true
		  },
		  "send": {
			"version": "0.17.1",
			"resolved": "https://registry.npmjs.org/send/-/send-0.17.1.tgz",
			"integrity": "sha512-BsVKsiGcQMFwT8UxypobUKyv7irCNRHk1T0G680vk88yf6LBByGcZJOTJCrTP2xVN6yI+XjPJcNuE3V4fT9sAg==",
			"requires": {
			  "debug": "2.6.9",
			  "depd": "~1.1.2",
			  "destroy": "~1.0.4",
			  "encodeurl": "~1.0.2",
			  "escape-html": "~1.0.3",
			  "etag": "~1.8.1",
			  "fresh": "0.5.2",
			  "http-errors": "~1.7.2",
			  "mime": "1.6.0",
			  "ms": "2.1.1",
			  "on-finished": "~2.3.0",
			  "range-parser": "~1.2.1",
			  "statuses": "~1.5.0"
			},
			"dependencies": {
			  "debug": {
				"version": "2.6.9",
				"resolved": "https://registry.npmjs.org/debug/-/debug-2.6.9.tgz",
				"integrity": "sha512-bC7ElrdJaJnPbAP+1EotYvqZsb3ecl5wi6Bfi6BJTUcNowp6cvspg0jXznRTKDjm/E7AdgFBVeAPVMNcKGsHMA==",
				"requires": {
				  "ms": "2.0.0"
				},
				"dependencies": {
				  "ms": {
					"version": "2.0.0",
					"resolved": "https://registry.npmjs.org/ms/-/ms-2.0.0.tgz",
					"integrity": "sha1-VgiurfwAvmwpAd9fmGF4jeDVl8g="
				  }
				}
			  },
			  "ms": {
				"version": "2.1.1",
				"resolved": "https://registry.npmjs.org/ms/-/ms-2.1.1.tgz",
				"integrity": "sha512-tgp+dl5cGk28utYktBsrFqA7HKgrhgPsg6Z/EfhWI4gl1Hwq8B/GmY/0oXZ6nF8hDVesS/FpnYaD/kOWhYQvyg=="
			  }
			}
		  },
		  "serialize-javascript": {
			"version": "2.1.2",
			"resolved": "https://registry.npmjs.org/serialize-javascript/-/serialize-javascript-2.1.2.tgz",
			"integrity": "sha512-rs9OggEUF0V4jUSecXazOYsLfu7OGK2qIn3c7IPBiffz32XniEp/TX9Xmc9LQfK2nQ2QKHvZ2oygKUGU0lG4jQ==",
			"dev": true
		  },
		  "serve-static": {
			"version": "1.14.1",
			"resolved": "https://registry.npmjs.org/serve-static/-/serve-static-1.14.1.tgz",
			"integrity": "sha512-JMrvUwE54emCYWlTI+hGrGv5I8dEwmco/00EvkzIIsR7MqrHonbD9pO2MOfFnpFntl7ecpZs+3mW+XbQZu9QCg==",
			"requires": {
			  "encodeurl": "~1.0.2",
			  "escape-html": "~1.0.3",
			  "parseurl": "~1.3.3",
			  "send": "0.17.1"
			}
		  },
		  "set-blocking": {
			"version": "2.0.0",
			"resolved": "https://registry.npmjs.org/set-blocking/-/set-blocking-2.0.0.tgz",
			"integrity": "sha1-BF+XgtARrppoA93TgrJDkrPYkPc=",
			"dev": true
		  },
		  "set-value": {
			"version": "2.0.1",
			"resolved": "https://registry.npmjs.org/set-value/-/set-value-2.0.1.tgz",
			"integrity": "sha512-JxHc1weCN68wRY0fhCoXpyK55m/XPHafOmK4UWD7m2CI14GMcFypt4w/0+NV5f/ZMby2F6S2wwA7fgynh9gWSw==",
			"dev": true,
			"requires": {
			  "extend-shallow": "^2.0.1",
			  "is-extendable": "^0.1.1",
			  "is-plain-object": "^2.0.3",
			  "split-string": "^3.0.1"
			},
			"dependencies": {
			  "extend-shallow": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/extend-shallow/-/extend-shallow-2.0.1.tgz",
				"integrity": "sha1-Ua99YUrZqfYQ6huvu5idaxxWiQ8=",
				"dev": true,
				"requires": {
				  "is-extendable": "^0.1.0"
				}
			  }
			}
		  },
		  "setimmediate": {
			"version": "1.0.5",
			"resolved": "https://registry.npmjs.org/setimmediate/-/setimmediate-1.0.5.tgz",
			"integrity": "sha1-KQy7Iy4waULX1+qbg3Mqt4VvgoU=",
			"dev": true
		  },
		  "setprototypeof": {
			"version": "1.1.1",
			"resolved": "https://registry.npmjs.org/setprototypeof/-/setprototypeof-1.1.1.tgz",
			"integrity": "sha512-JvdAWfbXeIGaZ9cILp38HntZSFSo3mWg6xGcJJsd+d4aRMOqauag1C63dJfDw7OaMYwEbHMOxEZ1lqVRYP2OAw=="
		  },
		  "sha.js": {
			"version": "2.4.11",
			"resolved": "https://registry.npmjs.org/sha.js/-/sha.js-2.4.11.tgz",
			"integrity": "sha512-QMEp5B7cftE7APOjk5Y6xgrbWu+WkLVQwk8JNjZ8nKRciZaByEW6MubieAiToS7+dwvrjGhH8jRXz3MVd0AYqQ==",
			"dev": true,
			"requires": {
			  "inherits": "^2.0.1",
			  "safe-buffer": "^5.0.1"
			}
		  },
		  "shebang-command": {
			"version": "1.2.0",
			"resolved": "https://registry.npmjs.org/shebang-command/-/shebang-command-1.2.0.tgz",
			"integrity": "sha1-RKrGW2lbAzmJaMOfNj/uXer98eo=",
			"dev": true,
			"requires": {
			  "shebang-regex": "^1.0.0"
			}
		  },
		  "shebang-regex": {
			"version": "1.0.0",
			"resolved": "https://registry.npmjs.org/shebang-regex/-/shebang-regex-1.0.0.tgz",
			"integrity": "sha1-2kL0l0DAtC2yypcoVxyxkMmO/qM=",
			"dev": true
		  },
		  "shelljs": {
			"version": "0.8.4",
			"resolved": "https://registry.npmjs.org/shelljs/-/shelljs-0.8.4.tgz",
			"integrity": "sha512-7gk3UZ9kOfPLIAbslLzyWeGiEqx9e3rxwZM0KE6EL8GlGwjym9Mrlx5/p33bWTu9YG6vcS4MBxYZDHYr5lr8BQ==",
			"dev": true,
			"requires": {
			  "glob": "^7.0.0",
			  "interpret": "^1.0.0",
			  "rechoir": "^0.6.2"
			}
		  },
		  "shellwords": {
			"version": "0.1.1",
			"resolved": "https://registry.npmjs.org/shellwords/-/shellwords-0.1.1.tgz",
			"integrity": "sha512-vFwSUfQvqybiICwZY5+DAWIPLKsWO31Q91JSKl3UYv+K5c2QRPzn0qzec6QPu1Qc9eHYItiP3NdJqNVqetYAww==",
			"dev": true,
			"optional": true
		  },
		  "signal-exit": {
			"version": "3.0.2",
			"resolved": "https://registry.npmjs.org/signal-exit/-/signal-exit-3.0.2.tgz",
			"integrity": "sha1-tf3AjxKH6hF4Yo5BXiUTK3NkbG0=",
			"dev": true
		  },
		  "sisteransi": {
			"version": "1.0.5",
			"resolved": "https://registry.npmjs.org/sisteransi/-/sisteransi-1.0.5.tgz",
			"integrity": "sha512-bLGGlR1QxBcynn2d5YmDX4MGjlZvy2MRBDRNHLJ8VI6l6+9FUiyTFNJ0IveOSP0bcXgVDPRcfGqA0pjaqUpfVg==",
			"dev": true
		  },
		  "slash": {
			"version": "3.0.0",
			"resolved": "https://registry.npmjs.org/slash/-/slash-3.0.0.tgz",
			"integrity": "sha512-g9Q1haeby36OSStwb4ntCGGGaKsaVSjQ68fBxoQcutl5fS1vuY18H3wSt3jFyFtrkx+Kz0V1G85A4MyAdDMi2Q==",
			"dev": true
		  },
		  "slice-ansi": {
			"version": "2.1.0",
			"resolved": "https://registry.npmjs.org/slice-ansi/-/slice-ansi-2.1.0.tgz",
			"integrity": "sha512-Qu+VC3EwYLldKa1fCxuuvULvSJOKEgk9pi8dZeCVK7TqBfUNTH4sFkk4joj8afVSfAYgJoSOetjx9QWOJ5mYoQ==",
			"dev": true,
			"requires": {
			  "ansi-styles": "^3.2.0",
			  "astral-regex": "^1.0.0",
			  "is-fullwidth-code-point": "^2.0.0"
			},
			"dependencies": {
			  "is-fullwidth-code-point": {
				"version": "2.0.0",
				"resolved": "https://registry.npmjs.org/is-fullwidth-code-point/-/is-fullwidth-code-point-2.0.0.tgz",
				"integrity": "sha1-o7MKXE8ZkYMWeqq5O+764937ZU8=",
				"dev": true
			  }
			}
		  },
		  "snapdragon": {
			"version": "0.8.2",
			"resolved": "https://registry.npmjs.org/snapdragon/-/snapdragon-0.8.2.tgz",
			"integrity": "sha512-FtyOnWN/wCHTVXOMwvSv26d+ko5vWlIDD6zoUJ7LW8vh+ZBC8QdljveRP+crNrtBwioEUWy/4dMtbBjA4ioNlg==",
			"dev": true,
			"requires": {
			  "base": "^0.11.1",
			  "debug": "^2.2.0",
			  "define-property": "^0.2.5",
			  "extend-shallow": "^2.0.1",
			  "map-cache": "^0.2.2",
			  "source-map": "^0.5.6",
			  "source-map-resolve": "^0.5.0",
			  "use": "^3.1.0"
			},
			"dependencies": {
			  "debug": {
				"version": "2.6.9",
				"resolved": "https://registry.npmjs.org/debug/-/debug-2.6.9.tgz",
				"integrity": "sha512-bC7ElrdJaJnPbAP+1EotYvqZsb3ecl5wi6Bfi6BJTUcNowp6cvspg0jXznRTKDjm/E7AdgFBVeAPVMNcKGsHMA==",
				"dev": true,
				"requires": {
				  "ms": "2.0.0"
				}
			  },
			  "define-property": {
				"version": "0.2.5",
				"resolved": "https://registry.npmjs.org/define-property/-/define-property-0.2.5.tgz",
				"integrity": "sha1-w1se+RjsPJkPmlvFe+BKrOxcgRY=",
				"dev": true,
				"requires": {
				  "is-descriptor": "^0.1.0"
				}
			  },
			  "extend-shallow": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/extend-shallow/-/extend-shallow-2.0.1.tgz",
				"integrity": "sha1-Ua99YUrZqfYQ6huvu5idaxxWiQ8=",
				"dev": true,
				"requires": {
				  "is-extendable": "^0.1.0"
				}
			  },
			  "source-map": {
				"version": "0.5.7",
				"resolved": "https://registry.npmjs.org/source-map/-/source-map-0.5.7.tgz",
				"integrity": "sha1-igOdLRAh0i0eoUyA2OpGi6LvP8w=",
				"dev": true
			  }
			}
		  },
		  "snapdragon-node": {
			"version": "2.1.1",
			"resolved": "https://registry.npmjs.org/snapdragon-node/-/snapdragon-node-2.1.1.tgz",
			"integrity": "sha512-O27l4xaMYt/RSQ5TR3vpWCAB5Kb/czIcqUFOM/C4fYcLnbZUc1PkjTAMjof2pBWaSTwOUd6qUHcFGVGj7aIwnw==",
			"dev": true,
			"requires": {
			  "define-property": "^1.0.0",
			  "isobject": "^3.0.0",
			  "snapdragon-util": "^3.0.1"
			},
			"dependencies": {
			  "define-property": {
				"version": "1.0.0",
				"resolved": "https://registry.npmjs.org/define-property/-/define-property-1.0.0.tgz",
				"integrity": "sha1-dp66rz9KY6rTr56NMEybvnm/sOY=",
				"dev": true,
				"requires": {
				  "is-descriptor": "^1.0.0"
				}
			  },
			  "is-accessor-descriptor": {
				"version": "1.0.0",
				"resolved": "https://registry.npmjs.org/is-accessor-descriptor/-/is-accessor-descriptor-1.0.0.tgz",
				"integrity": "sha512-m5hnHTkcVsPfqx3AKlyttIPb7J+XykHvJP2B9bZDjlhLIoEq4XoK64Vg7boZlVWYK6LUY94dYPEE7Lh0ZkZKcQ==",
				"dev": true,
				"requires": {
				  "kind-of": "^6.0.0"
				}
			  },
			  "is-data-descriptor": {
				"version": "1.0.0",
				"resolved": "https://registry.npmjs.org/is-data-descriptor/-/is-data-descriptor-1.0.0.tgz",
				"integrity": "sha512-jbRXy1FmtAoCjQkVmIVYwuuqDFUbaOeDjmed1tOGPrsMhtJA4rD9tkgA0F1qJ3gRFRXcHYVkdeaP50Q5rE/jLQ==",
				"dev": true,
				"requires": {
				  "kind-of": "^6.0.0"
				}
			  },
			  "is-descriptor": {
				"version": "1.0.2",
				"resolved": "https://registry.npmjs.org/is-descriptor/-/is-descriptor-1.0.2.tgz",
				"integrity": "sha512-2eis5WqQGV7peooDyLmNEPUrps9+SXX5c9pL3xEB+4e9HnGuDa7mB7kHxHw4CbqS9k1T2hOH3miL8n8WtiYVtg==",
				"dev": true,
				"requires": {
				  "is-accessor-descriptor": "^1.0.0",
				  "is-data-descriptor": "^1.0.0",
				  "kind-of": "^6.0.2"
				}
			  }
			}
		  },
		  "snapdragon-util": {
			"version": "3.0.1",
			"resolved": "https://registry.npmjs.org/snapdragon-util/-/snapdragon-util-3.0.1.tgz",
			"integrity": "sha512-mbKkMdQKsjX4BAL4bRYTj21edOf8cN7XHdYUJEe+Zn99hVEYcMvKPct1IqNe7+AZPirn8BCDOQBHQZknqmKlZQ==",
			"dev": true,
			"requires": {
			  "kind-of": "^3.2.0"
			},
			"dependencies": {
			  "kind-of": {
				"version": "3.2.2",
				"resolved": "https://registry.npmjs.org/kind-of/-/kind-of-3.2.2.tgz",
				"integrity": "sha1-MeohpzS6ubuw8yRm2JOupR5KPGQ=",
				"dev": true,
				"requires": {
				  "is-buffer": "^1.1.5"
				}
			  }
			}
		  },
		  "source-list-map": {
			"version": "2.0.1",
			"resolved": "https://registry.npmjs.org/source-list-map/-/source-list-map-2.0.1.tgz",
			"integrity": "sha512-qnQ7gVMxGNxsiL4lEuJwe/To8UnK7fAnmbGEEH8RpLouuKbeEm0lhbQVFIrNSuB+G7tVrAlVsZgETT5nljf+Iw==",
			"dev": true
		  },
		  "source-map": {
			"version": "0.7.3",
			"resolved": "https://registry.npmjs.org/source-map/-/source-map-0.7.3.tgz",
			"integrity": "sha512-CkCj6giN3S+n9qrYiBTX5gystlENnRW5jZeNLHpe6aue+SrHcG5VYwujhW9s4dY31mEGsxBDrHR6oI69fTXsaQ==",
			"dev": true
		  },
		  "source-map-resolve": {
			"version": "0.5.3",
			"resolved": "https://registry.npmjs.org/source-map-resolve/-/source-map-resolve-0.5.3.tgz",
			"integrity": "sha512-Htz+RnsXWk5+P2slx5Jh3Q66vhQj1Cllm0zvnaY98+NFx+Dv2CF/f5O/t8x+KaNdrdIAsruNzoh/KpialbqAnw==",
			"dev": true,
			"requires": {
			  "atob": "^2.1.2",
			  "decode-uri-component": "^0.2.0",
			  "resolve-url": "^0.2.1",
			  "source-map-url": "^0.4.0",
			  "urix": "^0.1.0"
			}
		  },
		  "source-map-support": {
			"version": "0.5.19",
			"resolved": "https://registry.npmjs.org/source-map-support/-/source-map-support-0.5.19.tgz",
			"integrity": "sha512-Wonm7zOCIJzBGQdB+thsPar0kYuCIzYvxZwlBa87yi/Mdjv7Tip2cyVbLj5o0cFPN4EVkuTwb3GDDyUx2DGnGw==",
			"dev": true,
			"requires": {
			  "buffer-from": "^1.0.0",
			  "source-map": "^0.6.0"
			},
			"dependencies": {
			  "source-map": {
				"version": "0.6.1",
				"resolved": "https://registry.npmjs.org/source-map/-/source-map-0.6.1.tgz",
				"integrity": "sha512-UjgapumWlbMhkBgzT7Ykc5YXUT46F0iKu8SGXq0bcwP5dz/h0Plj6enJqjz1Zbq2l5WaqYnrVbwWOWMyF3F47g==",
				"dev": true
			  }
			}
		  },
		  "source-map-url": {
			"version": "0.4.0",
			"resolved": "https://registry.npmjs.org/source-map-url/-/source-map-url-0.4.0.tgz",
			"integrity": "sha1-PpNdfd1zYxuXZZlW1VEo6HtQhKM=",
			"dev": true
		  },
		  "sourcemap-codec": {
			"version": "1.4.8",
			"resolved": "https://registry.npmjs.org/sourcemap-codec/-/sourcemap-codec-1.4.8.tgz",
			"integrity": "sha512-9NykojV5Uih4lgo5So5dtw+f0JgJX30KCNI8gwhz2J9A15wD0Ml6tjHKwf6fTSa6fAdVBdZeNOs9eJ71qCk8vA==",
			"dev": true
		  },
		  "spdx-correct": {
			"version": "3.1.0",
			"resolved": "https://registry.npmjs.org/spdx-correct/-/spdx-correct-3.1.0.tgz",
			"integrity": "sha512-lr2EZCctC2BNR7j7WzJ2FpDznxky1sjfxvvYEyzxNyb6lZXHODmEoJeFu4JupYlkfha1KZpJyoqiJ7pgA1qq8Q==",
			"dev": true,
			"requires": {
			  "spdx-expression-parse": "^3.0.0",
			  "spdx-license-ids": "^3.0.0"
			}
		  },
		  "spdx-exceptions": {
			"version": "2.3.0",
			"resolved": "https://registry.npmjs.org/spdx-exceptions/-/spdx-exceptions-2.3.0.tgz",
			"integrity": "sha512-/tTrYOC7PPI1nUAgx34hUpqXuyJG+DTHJTnIULG4rDygi4xu/tfgmq1e1cIRwRzwZgo4NLySi+ricLkZkw4i5A==",
			"dev": true
		  },
		  "spdx-expression-parse": {
			"version": "3.0.1",
			"resolved": "https://registry.npmjs.org/spdx-expression-parse/-/spdx-expression-parse-3.0.1.tgz",
			"integrity": "sha512-cbqHunsQWnJNE6KhVSMsMeH5H/L9EpymbzqTQ3uLwNCLZ1Q481oWaofqH7nO6V07xlXwY6PhQdQ2IedWx/ZK4Q==",
			"dev": true,
			"requires": {
			  "spdx-exceptions": "^2.1.0",
			  "spdx-license-ids": "^3.0.0"
			}
		  },
		  "spdx-license-ids": {
			"version": "3.0.5",
			"resolved": "https://registry.npmjs.org/spdx-license-ids/-/spdx-license-ids-3.0.5.tgz",
			"integrity": "sha512-J+FWzZoynJEXGphVIS+XEh3kFSjZX/1i9gFBaWQcB+/tmpe2qUsSBABpcxqxnAxFdiUFEgAX1bjYGQvIZmoz9Q==",
			"dev": true
		  },
		  "split-string": {
			"version": "3.1.0",
			"resolved": "https://registry.npmjs.org/split-string/-/split-string-3.1.0.tgz",
			"integrity": "sha512-NzNVhJDYpwceVVii8/Hu6DKfD2G+NrQHlS/V/qgv763EYudVwEcMQNxd2lh+0VrUByXN/oJkl5grOhYWvQUYiw==",
			"dev": true,
			"requires": {
			  "extend-shallow": "^3.0.0"
			}
		  },
		  "sprintf-js": {
			"version": "1.0.3",
			"resolved": "https://registry.npmjs.org/sprintf-js/-/sprintf-js-1.0.3.tgz",
			"integrity": "sha1-BOaSb2YolTVPPdAVIDYzuFcpfiw=",
			"dev": true
		  },
		  "sshpk": {
			"version": "1.16.1",
			"resolved": "https://registry.npmjs.org/sshpk/-/sshpk-1.16.1.tgz",
			"integrity": "sha512-HXXqVUq7+pcKeLqqZj6mHFUMvXtOJt1uoUx09pFW6011inTMxqI8BA8PM95myrIyyKwdnzjdFjLiE6KBPVtJIg==",
			"dev": true,
			"requires": {
			  "asn1": "~0.2.3",
			  "assert-plus": "^1.0.0",
			  "bcrypt-pbkdf": "^1.0.0",
			  "dashdash": "^1.12.0",
			  "ecc-jsbn": "~0.1.1",
			  "getpass": "^0.1.1",
			  "jsbn": "~0.1.0",
			  "safer-buffer": "^2.0.2",
			  "tweetnacl": "~0.14.0"
			}
		  },
		  "ssri": {
			"version": "6.0.1",
			"resolved": "https://registry.npmjs.org/ssri/-/ssri-6.0.1.tgz",
			"integrity": "sha512-3Wge10hNcT1Kur4PDFwEieXSCMCJs/7WvSACcrMYrNp+b8kDL1/0wJch5Ni2WrtwEa2IO8OsVfeKIciKCDx/QA==",
			"dev": true,
			"requires": {
			  "figgy-pudding": "^3.5.1"
			}
		  },
		  "stack-utils": {
			"version": "2.0.2",
			"resolved": "https://registry.npmjs.org/stack-utils/-/stack-utils-2.0.2.tgz",
			"integrity": "sha512-0H7QK2ECz3fyZMzQ8rH0j2ykpfbnd20BFtfg/SqVC2+sCTtcw0aDTGB7dk+de4U4uUeuz6nOtJcrkFFLG1B0Rg==",
			"dev": true,
			"requires": {
			  "escape-string-regexp": "^2.0.0"
			},
			"dependencies": {
			  "escape-string-regexp": {
				"version": "2.0.0",
				"resolved": "https://registry.npmjs.org/escape-string-regexp/-/escape-string-regexp-2.0.0.tgz",
				"integrity": "sha512-UpzcLCXolUWcNu5HtVMHYdXJjArjsF9C0aNnquZYY4uW/Vu0miy5YoWvbV345HauVvcAUnpRuhMMcqTcGOY2+w==",
				"dev": true
			  }
			}
		  },
		  "static-extend": {
			"version": "0.1.2",
			"resolved": "https://registry.npmjs.org/static-extend/-/static-extend-0.1.2.tgz",
			"integrity": "sha1-YICcOcv/VTNyJv1eC1IPNB8ftcY=",
			"dev": true,
			"requires": {
			  "define-property": "^0.2.5",
			  "object-copy": "^0.1.0"
			},
			"dependencies": {
			  "define-property": {
				"version": "0.2.5",
				"resolved": "https://registry.npmjs.org/define-property/-/define-property-0.2.5.tgz",
				"integrity": "sha1-w1se+RjsPJkPmlvFe+BKrOxcgRY=",
				"dev": true,
				"requires": {
				  "is-descriptor": "^0.1.0"
				}
			  }
			}
		  },
		  "statuses": {
			"version": "1.5.0",
			"resolved": "https://registry.npmjs.org/statuses/-/statuses-1.5.0.tgz",
			"integrity": "sha1-Fhx9rBd2Wf2YEfQ3cfqZOBR4Yow="
		  },
		  "stealthy-require": {
			"version": "1.1.1",
			"resolved": "https://registry.npmjs.org/stealthy-require/-/stealthy-require-1.1.1.tgz",
			"integrity": "sha1-NbCYdbT/SfJqd35QmzCQoyJr8ks=",
			"dev": true
		  },
		  "stream-browserify": {
			"version": "2.0.2",
			"resolved": "https://registry.npmjs.org/stream-browserify/-/stream-browserify-2.0.2.tgz",
			"integrity": "sha512-nX6hmklHs/gr2FuxYDltq8fJA1GDlxKQCz8O/IM4atRqBH8OORmBNgfvW5gG10GT/qQ9u0CzIvr2X5Pkt6ntqg==",
			"dev": true,
			"requires": {
			  "inherits": "~2.0.1",
			  "readable-stream": "^2.0.2"
			},
			"dependencies": {
			  "isarray": {
				"version": "1.0.0",
				"resolved": "https://registry.npmjs.org/isarray/-/isarray-1.0.0.tgz",
				"integrity": "sha1-u5NdSFgsuhaMBoNJV6VKPgcSTxE=",
				"dev": true
			  },
			  "readable-stream": {
				"version": "2.3.7",
				"resolved": "https://registry.npmjs.org/readable-stream/-/readable-stream-2.3.7.tgz",
				"integrity": "sha512-Ebho8K4jIbHAxnuxi7o42OrZgF/ZTNcsZj6nRKyUmkhLFq8CHItp/fy6hQZuZmP/n3yZ9VBUbp4zz/mX8hmYPw==",
				"dev": true,
				"requires": {
				  "core-util-is": "~1.0.0",
				  "inherits": "~2.0.3",
				  "isarray": "~1.0.0",
				  "process-nextick-args": "~2.0.0",
				  "safe-buffer": "~5.1.1",
				  "string_decoder": "~1.1.1",
				  "util-deprecate": "~1.0.1"
				}
			  },
			  "string_decoder": {
				"version": "1.1.1",
				"resolved": "https://registry.npmjs.org/string_decoder/-/string_decoder-1.1.1.tgz",
				"integrity": "sha512-n/ShnvDi6FHbbVfviro+WojiFzv+s8MPMHBczVePfUpDJLwoLT0ht1l4YwBCbi8pJAveEEdnkHyPyTP/mzRfwg==",
				"dev": true,
				"requires": {
				  "safe-buffer": "~5.1.0"
				}
			  }
			}
		  },
		  "stream-each": {
			"version": "1.2.3",
			"resolved": "https://registry.npmjs.org/stream-each/-/stream-each-1.2.3.tgz",
			"integrity": "sha512-vlMC2f8I2u/bZGqkdfLQW/13Zihpej/7PmSiMQsbYddxuTsJp8vRe2x2FvVExZg7FaOds43ROAuFJwPR4MTZLw==",
			"dev": true,
			"requires": {
			  "end-of-stream": "^1.1.0",
			  "stream-shift": "^1.0.0"
			}
		  },
		  "stream-http": {
			"version": "2.8.3",
			"resolved": "https://registry.npmjs.org/stream-http/-/stream-http-2.8.3.tgz",
			"integrity": "sha512-+TSkfINHDo4J+ZobQLWiMouQYB+UVYFttRA94FpEzzJ7ZdqcL4uUUQ7WkdkI4DSozGmgBUE/a47L+38PenXhUw==",
			"dev": true,
			"requires": {
			  "builtin-status-codes": "^3.0.0",
			  "inherits": "^2.0.1",
			  "readable-stream": "^2.3.6",
			  "to-arraybuffer": "^1.0.0",
			  "xtend": "^4.0.0"
			},
			"dependencies": {
			  "isarray": {
				"version": "1.0.0",
				"resolved": "https://registry.npmjs.org/isarray/-/isarray-1.0.0.tgz",
				"integrity": "sha1-u5NdSFgsuhaMBoNJV6VKPgcSTxE=",
				"dev": true
			  },
			  "readable-stream": {
				"version": "2.3.7",
				"resolved": "https://registry.npmjs.org/readable-stream/-/readable-stream-2.3.7.tgz",
				"integrity": "sha512-Ebho8K4jIbHAxnuxi7o42OrZgF/ZTNcsZj6nRKyUmkhLFq8CHItp/fy6hQZuZmP/n3yZ9VBUbp4zz/mX8hmYPw==",
				"dev": true,
				"requires": {
				  "core-util-is": "~1.0.0",
				  "inherits": "~2.0.3",
				  "isarray": "~1.0.0",
				  "process-nextick-args": "~2.0.0",
				  "safe-buffer": "~5.1.1",
				  "string_decoder": "~1.1.1",
				  "util-deprecate": "~1.0.1"
				}
			  },
			  "string_decoder": {
				"version": "1.1.1",
				"resolved": "https://registry.npmjs.org/string_decoder/-/string_decoder-1.1.1.tgz",
				"integrity": "sha512-n/ShnvDi6FHbbVfviro+WojiFzv+s8MPMHBczVePfUpDJLwoLT0ht1l4YwBCbi8pJAveEEdnkHyPyTP/mzRfwg==",
				"dev": true,
				"requires": {
				  "safe-buffer": "~5.1.0"
				}
			  }
			}
		  },
		  "stream-shift": {
			"version": "1.0.1",
			"resolved": "https://registry.npmjs.org/stream-shift/-/stream-shift-1.0.1.tgz",
			"integrity": "sha512-AiisoFqQ0vbGcZgQPY1cdP2I76glaVA/RauYR4G4thNFgkTqr90yXTo4LYX60Jl+sIlPNHHdGSwo01AvbKUSVQ==",
			"dev": true
		  },
		  "streamsearch": {
			"version": "0.1.2",
			"resolved": "https://registry.npmjs.org/streamsearch/-/streamsearch-0.1.2.tgz",
			"integrity": "sha1-gIudDlb8Jz2Am6VzOOkpkZoanxo="
		  },
		  "string-length": {
			"version": "4.0.1",
			"resolved": "https://registry.npmjs.org/string-length/-/string-length-4.0.1.tgz",
			"integrity": "sha512-PKyXUd0LK0ePjSOnWn34V2uD6acUWev9uy0Ft05k0E8xRW+SKcA0F7eMr7h5xlzfn+4O3N+55rduYyet3Jk+jw==",
			"dev": true,
			"requires": {
			  "char-regex": "^1.0.2",
			  "strip-ansi": "^6.0.0"
			},
			"dependencies": {
			  "ansi-regex": {
				"version": "5.0.0",
				"resolved": "https://registry.npmjs.org/ansi-regex/-/ansi-regex-5.0.0.tgz",
				"integrity": "sha512-bY6fj56OUQ0hU1KjFNDQuJFezqKdrAyFdIevADiqrWHwSlbmBNMHp5ak2f40Pm8JTFyM2mqxkG6ngkHO11f/lg==",
				"dev": true
			  },
			  "strip-ansi": {
				"version": "6.0.0",
				"resolved": "https://registry.npmjs.org/strip-ansi/-/strip-ansi-6.0.0.tgz",
				"integrity": "sha512-AuvKTrTfQNYNIctbR1K/YGTR1756GycPsg7b9bdV9Duqur4gv6aKqHXah67Z8ImS7WEz5QVcOtlfW2rZEugt6w==",
				"dev": true,
				"requires": {
				  "ansi-regex": "^5.0.0"
				}
			  }
			}
		  },
		  "string-width": {
			"version": "4.2.0",
			"resolved": "https://registry.npmjs.org/string-width/-/string-width-4.2.0.tgz",
			"integrity": "sha512-zUz5JD+tgqtuDjMhwIg5uFVV3dtqZ9yQJlZVfq4I01/K5Paj5UHj7VyrQOJvzawSVlKpObApbfD0Ed6yJc+1eg==",
			"dev": true,
			"requires": {
			  "emoji-regex": "^8.0.0",
			  "is-fullwidth-code-point": "^3.0.0",
			  "strip-ansi": "^6.0.0"
			},
			"dependencies": {
			  "ansi-regex": {
				"version": "5.0.0",
				"resolved": "https://registry.npmjs.org/ansi-regex/-/ansi-regex-5.0.0.tgz",
				"integrity": "sha512-bY6fj56OUQ0hU1KjFNDQuJFezqKdrAyFdIevADiqrWHwSlbmBNMHp5ak2f40Pm8JTFyM2mqxkG6ngkHO11f/lg==",
				"dev": true
			  },
			  "strip-ansi": {
				"version": "6.0.0",
				"resolved": "https://registry.npmjs.org/strip-ansi/-/strip-ansi-6.0.0.tgz",
				"integrity": "sha512-AuvKTrTfQNYNIctbR1K/YGTR1756GycPsg7b9bdV9Duqur4gv6aKqHXah67Z8ImS7WEz5QVcOtlfW2rZEugt6w==",
				"dev": true,
				"requires": {
				  "ansi-regex": "^5.0.0"
				}
			  }
			}
		  },
		  "string.prototype.trimend": {
			"version": "1.0.1",
			"resolved": "https://registry.npmjs.org/string.prototype.trimend/-/string.prototype.trimend-1.0.1.tgz",
			"integrity": "sha512-LRPxFUaTtpqYsTeNKaFOw3R4bxIzWOnbQ837QfBylo8jIxtcbK/A/sMV7Q+OAV/vWo+7s25pOE10KYSjaSO06g==",
			"dev": true,
			"requires": {
			  "define-properties": "^1.1.3",
			  "es-abstract": "^1.17.5"
			}
		  },
		  "string.prototype.trimleft": {
			"version": "2.1.2",
			"resolved": "https://registry.npmjs.org/string.prototype.trimleft/-/string.prototype.trimleft-2.1.2.tgz",
			"integrity": "sha512-gCA0tza1JBvqr3bfAIFJGqfdRTyPae82+KTnm3coDXkZN9wnuW3HjGgN386D7hfv5CHQYCI022/rJPVlqXyHSw==",
			"dev": true,
			"requires": {
			  "define-properties": "^1.1.3",
			  "es-abstract": "^1.17.5",
			  "string.prototype.trimstart": "^1.0.0"
			}
		  },
		  "string.prototype.trimright": {
			"version": "2.1.2",
			"resolved": "https://registry.npmjs.org/string.prototype.trimright/-/string.prototype.trimright-2.1.2.tgz",
			"integrity": "sha512-ZNRQ7sY3KroTaYjRS6EbNiiHrOkjihL9aQE/8gfQ4DtAC/aEBRHFJa44OmoWxGGqXuJlfKkZW4WcXErGr+9ZFg==",
			"dev": true,
			"requires": {
			  "define-properties": "^1.1.3",
			  "es-abstract": "^1.17.5",
			  "string.prototype.trimend": "^1.0.0"
			}
		  },
		  "string.prototype.trimstart": {
			"version": "1.0.1",
			"resolved": "https://registry.npmjs.org/string.prototype.trimstart/-/string.prototype.trimstart-1.0.1.tgz",
			"integrity": "sha512-XxZn+QpvrBI1FOcg6dIpxUPgWCPuNXvMD72aaRaUQv1eD4e/Qy8i/hFTe0BUmD60p/QA6bh1avmuPTfNjqVWRw==",
			"dev": true,
			"requires": {
			  "define-properties": "^1.1.3",
			  "es-abstract": "^1.17.5"
			}
		  },
		  "string_decoder": {
			"version": "0.10.31",
			"resolved": "https://registry.npmjs.org/string_decoder/-/string_decoder-0.10.31.tgz",
			"integrity": "sha1-YuIDvEF2bGwoyfyEMB2rHFMQ+pQ="
		  },
		  "strip-ansi": {
			"version": "5.2.0",
			"resolved": "https://registry.npmjs.org/strip-ansi/-/strip-ansi-5.2.0.tgz",
			"integrity": "sha512-DuRs1gKbBqsMKIZlrffwlug8MHkcnpjs5VPmL1PAh+mA30U0DTotfDZ0d2UUsXpPmPmMMJ6W773MaA3J+lbiWA==",
			"dev": true,
			"requires": {
			  "ansi-regex": "^4.1.0"
			},
			"dependencies": {
			  "ansi-regex": {
				"version": "4.1.0",
				"resolved": "https://registry.npmjs.org/ansi-regex/-/ansi-regex-4.1.0.tgz",
				"integrity": "sha512-1apePfXM1UOSqw0o9IiFAovVz9M5S1Dg+4TrDwfMewQ6p/rmMueb7tWZjQ1rx4Loy1ArBggoqGpfqqdI4rondg==",
				"dev": true
			  }
			}
		  },
		  "strip-bom": {
			"version": "3.0.0",
			"resolved": "https://registry.npmjs.org/strip-bom/-/strip-bom-3.0.0.tgz",
			"integrity": "sha1-IzTBjpx1n3vdVv3vfprj1YjmjtM=",
			"dev": true
		  },
		  "strip-eof": {
			"version": "1.0.0",
			"resolved": "https://registry.npmjs.org/strip-eof/-/strip-eof-1.0.0.tgz",
			"integrity": "sha1-u0P/VZim6wXYm1n80SnJgzE2Br8=",
			"dev": true
		  },
		  "strip-final-newline": {
			"version": "2.0.0",
			"resolved": "https://registry.npmjs.org/strip-final-newline/-/strip-final-newline-2.0.0.tgz",
			"integrity": "sha512-BrpvfNAE3dcvq7ll3xVumzjKjZQ5tI1sEUIKr3Uoks0XUl45St3FlatVqef9prk4jRDzhW6WZg+3bk93y6pLjA==",
			"dev": true
		  },
		  "strip-json-comments": {
			"version": "3.1.0",
			"resolved": "https://registry.npmjs.org/strip-json-comments/-/strip-json-comments-3.1.0.tgz",
			"integrity": "sha512-e6/d0eBu7gHtdCqFt0xJr642LdToM5/cN4Qb9DbHjVx1CP5RyeM+zH7pbecEmDv/lBqb0QH+6Uqq75rxFPkM0w==",
			"dev": true
		  },
		  "superagent": {
			"version": "3.8.3",
			"resolved": "https://registry.npmjs.org/superagent/-/superagent-3.8.3.tgz",
			"integrity": "sha512-GLQtLMCoEIK4eDv6OGtkOoSMt3D+oq0y3dsxMuYuDvaNUvuT8eFBuLmfR0iYYzHC1e8hpzC6ZsxbuP6DIalMFA==",
			"dev": true,
			"requires": {
			  "component-emitter": "^1.2.0",
			  "cookiejar": "^2.1.0",
			  "debug": "^3.1.0",
			  "extend": "^3.0.0",
			  "form-data": "^2.3.1",
			  "formidable": "^1.2.0",
			  "methods": "^1.1.1",
			  "mime": "^1.4.1",
			  "qs": "^6.5.1",
			  "readable-stream": "^2.3.5"
			},
			"dependencies": {
			  "isarray": {
				"version": "1.0.0",
				"resolved": "https://registry.npmjs.org/isarray/-/isarray-1.0.0.tgz",
				"integrity": "sha1-u5NdSFgsuhaMBoNJV6VKPgcSTxE=",
				"dev": true
			  },
			  "readable-stream": {
				"version": "2.3.6",
				"resolved": "https://registry.npmjs.org/readable-stream/-/readable-stream-2.3.6.tgz",
				"integrity": "sha512-tQtKA9WIAhBF3+VLAseyMqZeBjW0AHJoxOtYqSUZNJxauErmLbVm2FW1y+J/YA9dUrAC39ITejlZWhVIwawkKw==",
				"dev": true,
				"requires": {
				  "core-util-is": "~1.0.0",
				  "inherits": "~2.0.3",
				  "isarray": "~1.0.0",
				  "process-nextick-args": "~2.0.0",
				  "safe-buffer": "~5.1.1",
				  "string_decoder": "~1.1.1",
				  "util-deprecate": "~1.0.1"
				}
			  },
			  "string_decoder": {
				"version": "1.1.1",
				"resolved": "https://registry.npmjs.org/string_decoder/-/string_decoder-1.1.1.tgz",
				"integrity": "sha512-n/ShnvDi6FHbbVfviro+WojiFzv+s8MPMHBczVePfUpDJLwoLT0ht1l4YwBCbi8pJAveEEdnkHyPyTP/mzRfwg==",
				"dev": true,
				"requires": {
				  "safe-buffer": "~5.1.0"
				}
			  }
			}
		  },
		  "supertest": {
			"version": "4.0.2",
			"resolved": "https://registry.npmjs.org/supertest/-/supertest-4.0.2.tgz",
			"integrity": "sha512-1BAbvrOZsGA3YTCWqbmh14L0YEq0EGICX/nBnfkfVJn7SrxQV1I3pMYjSzG9y/7ZU2V9dWqyqk2POwxlb09duQ==",
			"dev": true,
			"requires": {
			  "methods": "^1.1.2",
			  "superagent": "^3.8.3"
			}
		  },
		  "supports-color": {
			"version": "5.5.0",
			"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-5.5.0.tgz",
			"integrity": "sha512-QjVjwdXIt408MIiAqCX4oUKsgU2EqAGzs2Ppkm4aQYbjm+ZEWEcW4SfFNTr4uMNZma0ey4f5lgLrkB0aX0QMow==",
			"requires": {
			  "has-flag": "^3.0.0"
			}
		  },
		  "supports-hyperlinks": {
			"version": "2.1.0",
			"resolved": "https://registry.npmjs.org/supports-hyperlinks/-/supports-hyperlinks-2.1.0.tgz",
			"integrity": "sha512-zoE5/e+dnEijk6ASB6/qrK+oYdm2do1hjoLWrqUC/8WEIW1gbxFcKuBof7sW8ArN6e+AYvsE8HBGiVRWL/F5CA==",
			"dev": true,
			"requires": {
			  "has-flag": "^4.0.0",
			  "supports-color": "^7.0.0"
			},
			"dependencies": {
			  "has-flag": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/has-flag/-/has-flag-4.0.0.tgz",
				"integrity": "sha512-EykJT/Q1KjTWctppgIAgfSO0tKVuZUjhgMr17kqTumMl6Afv3EISleU7qZUzoXDFTAHTDC4NOoG/ZxU3EvlMPQ==",
				"dev": true
			  },
			  "supports-color": {
				"version": "7.1.0",
				"resolved": "https://registry.npmjs.org/supports-color/-/supports-color-7.1.0.tgz",
				"integrity": "sha512-oRSIpR8pxT1Wr2FquTNnGet79b3BWljqOuoW/h4oBhxJ/HUbX5nX6JSruTkvXDCFMwDPvsaTTbvMLKZWSy0R5g==",
				"dev": true,
				"requires": {
				  "has-flag": "^4.0.0"
				}
			  }
			}
		  },
		  "symbol-observable": {
			"version": "1.2.0",
			"resolved": "https://registry.npmjs.org/symbol-observable/-/symbol-observable-1.2.0.tgz",
			"integrity": "sha512-e900nM8RRtGhlV36KGEU9k65K3mPb1WV70OdjfxlG2EAuM1noi/E/BaW/uMhL7bPEssK8QV57vN3esixjUvcXQ==",
			"dev": true
		  },
		  "symbol-tree": {
			"version": "3.2.4",
			"resolved": "https://registry.npmjs.org/symbol-tree/-/symbol-tree-3.2.4.tgz",
			"integrity": "sha512-9QNk5KwDF+Bvz+PyObkmSYjI5ksVUYtjW7AU22r2NKcfLJcXp96hkDWU3+XndOsUb+AQ9QhfzfCT2O+CNWT5Tw==",
			"dev": true
		  },
		  "table": {
			"version": "5.4.6",
			"resolved": "https://registry.npmjs.org/table/-/table-5.4.6.tgz",
			"integrity": "sha512-wmEc8m4fjnob4gt5riFRtTu/6+4rSe12TpAELNSqHMfF3IqnA+CH37USM6/YR3qRZv7e56kAEAtd6nKZaxe0Ug==",
			"dev": true,
			"requires": {
			  "ajv": "^6.10.2",
			  "lodash": "^4.17.14",
			  "slice-ansi": "^2.1.0",
			  "string-width": "^3.0.0"
			},
			"dependencies": {
			  "emoji-regex": {
				"version": "7.0.3",
				"resolved": "https://registry.npmjs.org/emoji-regex/-/emoji-regex-7.0.3.tgz",
				"integrity": "sha512-CwBLREIQ7LvYFB0WyRvwhq5N5qPhc6PMjD6bYggFlI5YyDgl+0vxq5VHbMOFqLg7hfWzmu8T5Z1QofhmTIhItA==",
				"dev": true
			  },
			  "is-fullwidth-code-point": {
				"version": "2.0.0",
				"resolved": "https://registry.npmjs.org/is-fullwidth-code-point/-/is-fullwidth-code-point-2.0.0.tgz",
				"integrity": "sha1-o7MKXE8ZkYMWeqq5O+764937ZU8=",
				"dev": true
			  },
			  "string-width": {
				"version": "3.1.0",
				"resolved": "https://registry.npmjs.org/string-width/-/string-width-3.1.0.tgz",
				"integrity": "sha512-vafcv6KjVZKSgz06oM/H6GDBrAtz8vdhQakGjFIvNrHA6y3HCF1CInLy+QLq8dTJPQ1b+KDUqDFctkdRW44e1w==",
				"dev": true,
				"requires": {
				  "emoji-regex": "^7.0.1",
				  "is-fullwidth-code-point": "^2.0.0",
				  "strip-ansi": "^5.1.0"
				}
			  }
			}
		  },
		  "tapable": {
			"version": "1.1.3",
			"resolved": "https://registry.npmjs.org/tapable/-/tapable-1.1.3.tgz",
			"integrity": "sha512-4WK/bYZmj8xLr+HUCODHGF1ZFzsYffasLUgEiMBY4fgtltdO6B4WJtlSbPaDTLpYTcGVwM2qLnFTICEcNxs3kA==",
			"dev": true
		  },
		  "terminal-link": {
			"version": "2.1.1",
			"resolved": "https://registry.npmjs.org/terminal-link/-/terminal-link-2.1.1.tgz",
			"integrity": "sha512-un0FmiRUQNr5PJqy9kP7c40F5BOfpGlYTrxonDChEZB7pzZxRNp/bt+ymiy9/npwXya9KH99nJ/GXFIiUkYGFQ==",
			"dev": true,
			"requires": {
			  "ansi-escapes": "^4.2.1",
			  "supports-hyperlinks": "^2.0.0"
			}
		  },
		  "terser": {
			"version": "4.7.0",
			"resolved": "https://registry.npmjs.org/terser/-/terser-4.7.0.tgz",
			"integrity": "sha512-Lfb0RiZcjRDXCC3OSHJpEkxJ9Qeqs6mp2v4jf2MHfy8vGERmVDuvjXdd/EnP5Deme5F2yBRBymKmKHCBg2echw==",
			"dev": true,
			"requires": {
			  "commander": "^2.20.0",
			  "source-map": "~0.6.1",
			  "source-map-support": "~0.5.12"
			},
			"dependencies": {
			  "commander": {
				"version": "2.20.3",
				"resolved": "https://registry.npmjs.org/commander/-/commander-2.20.3.tgz",
				"integrity": "sha512-GpVkmM8vF2vQUkj2LvZmD35JxeJOLCwJ9cUkugyk2nuhbv3+mJvpLYYt+0+USMxE+oj+ey/lJEnhZw75x/OMcQ==",
				"dev": true
			  },
			  "source-map": {
				"version": "0.6.1",
				"resolved": "https://registry.npmjs.org/source-map/-/source-map-0.6.1.tgz",
				"integrity": "sha512-UjgapumWlbMhkBgzT7Ykc5YXUT46F0iKu8SGXq0bcwP5dz/h0Plj6enJqjz1Zbq2l5WaqYnrVbwWOWMyF3F47g==",
				"dev": true
			  }
			}
		  },
		  "terser-webpack-plugin": {
			"version": "1.4.3",
			"resolved": "https://registry.npmjs.org/terser-webpack-plugin/-/terser-webpack-plugin-1.4.3.tgz",
			"integrity": "sha512-QMxecFz/gHQwteWwSo5nTc6UaICqN1bMedC5sMtUc7y3Ha3Q8y6ZO0iCR8pq4RJC8Hjf0FEPEHZqcMB/+DFCrA==",
			"dev": true,
			"requires": {
			  "cacache": "^12.0.2",
			  "find-cache-dir": "^2.1.0",
			  "is-wsl": "^1.1.0",
			  "schema-utils": "^1.0.0",
			  "serialize-javascript": "^2.1.2",
			  "source-map": "^0.6.1",
			  "terser": "^4.1.2",
			  "webpack-sources": "^1.4.0",
			  "worker-farm": "^1.7.0"
			},
			"dependencies": {
			  "source-map": {
				"version": "0.6.1",
				"resolved": "https://registry.npmjs.org/source-map/-/source-map-0.6.1.tgz",
				"integrity": "sha512-UjgapumWlbMhkBgzT7Ykc5YXUT46F0iKu8SGXq0bcwP5dz/h0Plj6enJqjz1Zbq2l5WaqYnrVbwWOWMyF3F47g==",
				"dev": true
			  }
			}
		  },
		  "test-exclude": {
			"version": "6.0.0",
			"resolved": "https://registry.npmjs.org/test-exclude/-/test-exclude-6.0.0.tgz",
			"integrity": "sha512-cAGWPIyOHU6zlmg88jwm7VRyXnMN7iV68OGAbYDk/Mh/xC/pzVPlQtY6ngoIH/5/tciuhGfvESU8GrHrcxD56w==",
			"dev": true,
			"requires": {
			  "@istanbuljs/schema": "^0.1.2",
			  "glob": "^7.1.4",
			  "minimatch": "^3.0.4"
			}
		  },
		  "text-table": {
			"version": "0.2.0",
			"resolved": "https://registry.npmjs.org/text-table/-/text-table-0.2.0.tgz",
			"integrity": "sha1-f17oI66AUgfACvLfSoTsP8+lcLQ=",
			"dev": true
		  },
		  "throat": {
			"version": "5.0.0",
			"resolved": "https://registry.npmjs.org/throat/-/throat-5.0.0.tgz",
			"integrity": "sha512-fcwX4mndzpLQKBS1DVYhGAcYaYt7vsHNIvQV+WXMvnow5cgjPphq5CaayLaGsjRdSCKZFNGt7/GYAuXaNOiYCA==",
			"dev": true
		  },
		  "through": {
			"version": "2.3.8",
			"resolved": "https://registry.npmjs.org/through/-/through-2.3.8.tgz",
			"integrity": "sha1-DdTJ/6q8NXlgsbckEV1+Doai4fU=",
			"dev": true
		  },
		  "through2": {
			"version": "2.0.5",
			"resolved": "https://registry.npmjs.org/through2/-/through2-2.0.5.tgz",
			"integrity": "sha512-/mrRod8xqpA+IHSLyGCQ2s8SPHiCDEeQJSep1jqLYeEUClOFG2Qsh+4FU6G9VeqpZnGW/Su8LQGc4YKni5rYSQ==",
			"dev": true,
			"requires": {
			  "readable-stream": "~2.3.6",
			  "xtend": "~4.0.1"
			},
			"dependencies": {
			  "isarray": {
				"version": "1.0.0",
				"resolved": "https://registry.npmjs.org/isarray/-/isarray-1.0.0.tgz",
				"integrity": "sha1-u5NdSFgsuhaMBoNJV6VKPgcSTxE=",
				"dev": true
			  },
			  "readable-stream": {
				"version": "2.3.7",
				"resolved": "https://registry.npmjs.org/readable-stream/-/readable-stream-2.3.7.tgz",
				"integrity": "sha512-Ebho8K4jIbHAxnuxi7o42OrZgF/ZTNcsZj6nRKyUmkhLFq8CHItp/fy6hQZuZmP/n3yZ9VBUbp4zz/mX8hmYPw==",
				"dev": true,
				"requires": {
				  "core-util-is": "~1.0.0",
				  "inherits": "~2.0.3",
				  "isarray": "~1.0.0",
				  "process-nextick-args": "~2.0.0",
				  "safe-buffer": "~5.1.1",
				  "string_decoder": "~1.1.1",
				  "util-deprecate": "~1.0.1"
				}
			  },
			  "string_decoder": {
				"version": "1.1.1",
				"resolved": "https://registry.npmjs.org/string_decoder/-/string_decoder-1.1.1.tgz",
				"integrity": "sha512-n/ShnvDi6FHbbVfviro+WojiFzv+s8MPMHBczVePfUpDJLwoLT0ht1l4YwBCbi8pJAveEEdnkHyPyTP/mzRfwg==",
				"dev": true,
				"requires": {
				  "safe-buffer": "~5.1.0"
				}
			  }
			}
		  },
		  "timers-browserify": {
			"version": "2.0.11",
			"resolved": "https://registry.npmjs.org/timers-browserify/-/timers-browserify-2.0.11.tgz",
			"integrity": "sha512-60aV6sgJ5YEbzUdn9c8kYGIqOubPoUdqQCul3SBAsRCZ40s6Y5cMcrW4dt3/k/EsbLVJNl9n6Vz3fTc+k2GeKQ==",
			"dev": true,
			"requires": {
			  "setimmediate": "^1.0.4"
			}
		  },
		  "timers-ext": {
			"version": "0.1.7",
			"resolved": "https://registry.npmjs.org/timers-ext/-/timers-ext-0.1.7.tgz",
			"integrity": "sha512-b85NUNzTSdodShTIbky6ZF02e8STtVVfD+fu4aXXShEELpozH+bCpJLYMPZbsABN2wDH7fJpqIoXxJpzbf0NqQ==",
			"requires": {
			  "es5-ext": "~0.10.46",
			  "next-tick": "1"
			}
		  },
		  "tmp": {
			"version": "0.0.33",
			"resolved": "https://registry.npmjs.org/tmp/-/tmp-0.0.33.tgz",
			"integrity": "sha512-jRCJlojKnZ3addtTOjdIqoRuPEKBvNXcGYqzO6zWZX8KfKEpnGY5jfggJQ3EjKuu8D4bJRr0y+cYJFmYbImXGw==",
			"dev": true,
			"requires": {
			  "os-tmpdir": "~1.0.2"
			}
		  },
		  "tmpl": {
			"version": "1.0.4",
			"resolved": "https://registry.npmjs.org/tmpl/-/tmpl-1.0.4.tgz",
			"integrity": "sha1-I2QN17QtAEM5ERQIIOXPRA5SHdE=",
			"dev": true
		  },
		  "to-arraybuffer": {
			"version": "1.0.1",
			"resolved": "https://registry.npmjs.org/to-arraybuffer/-/to-arraybuffer-1.0.1.tgz",
			"integrity": "sha1-fSKbH8xjfkZsoIEYCDanqr/4P0M=",
			"dev": true
		  },
		  "to-fast-properties": {
			"version": "2.0.0",
			"resolved": "https://registry.npmjs.org/to-fast-properties/-/to-fast-properties-2.0.0.tgz",
			"integrity": "sha1-3F5pjL0HkmW8c+A3doGk5Og/YW4=",
			"dev": true
		  },
		  "to-object-path": {
			"version": "0.3.0",
			"resolved": "https://registry.npmjs.org/to-object-path/-/to-object-path-0.3.0.tgz",
			"integrity": "sha1-KXWIt7Dn4KwI4E5nL4XB9JmeF68=",
			"dev": true,
			"requires": {
			  "kind-of": "^3.0.2"
			},
			"dependencies": {
			  "kind-of": {
				"version": "3.2.2",
				"resolved": "https://registry.npmjs.org/kind-of/-/kind-of-3.2.2.tgz",
				"integrity": "sha1-MeohpzS6ubuw8yRm2JOupR5KPGQ=",
				"dev": true,
				"requires": {
				  "is-buffer": "^1.1.5"
				}
			  }
			}
		  },
		  "to-regex": {
			"version": "3.0.2",
			"resolved": "https://registry.npmjs.org/to-regex/-/to-regex-3.0.2.tgz",
			"integrity": "sha512-FWtleNAtZ/Ki2qtqej2CXTOayOH9bHDQF+Q48VpWyDXjbYxA4Yz8iDB31zXOBUlOHHKidDbqGVrTUvQMPmBGBw==",
			"dev": true,
			"requires": {
			  "define-property": "^2.0.2",
			  "extend-shallow": "^3.0.2",
			  "regex-not": "^1.0.2",
			  "safe-regex": "^1.1.0"
			}
		  },
		  "to-regex-range": {
			"version": "5.0.1",
			"resolved": "https://registry.npmjs.org/to-regex-range/-/to-regex-range-5.0.1.tgz",
			"integrity": "sha512-65P7iz6X5yEr1cwcgvQxbbIw7Uk3gOy5dIdtZ4rDveLqhrdJP+Li/Hx6tyK0NEb+2GCyneCMJiGqrADCSNk8sQ==",
			"dev": true,
			"requires": {
			  "is-number": "^7.0.0"
			}
		  },
		  "toidentifier": {
			"version": "1.0.0",
			"resolved": "https://registry.npmjs.org/toidentifier/-/toidentifier-1.0.0.tgz",
			"integrity": "sha512-yaOH/Pk/VEhBWWTlhI+qXxDFXlejDGcQipMlyxda9nthulaxLZUNcUqFxokp0vcYnvteJln5FNQDRrxj3YcbVw=="
		  },
		  "tough-cookie": {
			"version": "3.0.1",
			"resolved": "https://registry.npmjs.org/tough-cookie/-/tough-cookie-3.0.1.tgz",
			"integrity": "sha512-yQyJ0u4pZsv9D4clxO69OEjLWYw+jbgspjTue4lTQZLfV0c5l1VmK2y1JK8E9ahdpltPOaAThPcp5nKPUgSnsg==",
			"dev": true,
			"requires": {
			  "ip-regex": "^2.1.0",
			  "psl": "^1.1.28",
			  "punycode": "^2.1.1"
			}
		  },
		  "tr46": {
			"version": "2.0.2",
			"resolved": "https://registry.npmjs.org/tr46/-/tr46-2.0.2.tgz",
			"integrity": "sha512-3n1qG+/5kg+jrbTzwAykB5yRYtQCTqOGKq5U5PE3b0a1/mzo6snDhjGS0zJVJunO0NrT3Dg1MLy5TjWP/UJppg==",
			"dev": true,
			"requires": {
			  "punycode": "^2.1.1"
			}
		  },
		  "tree-kill": {
			"version": "1.2.2",
			"resolved": "https://registry.npmjs.org/tree-kill/-/tree-kill-1.2.2.tgz",
			"integrity": "sha512-L0Orpi8qGpRG//Nd+H90vFB+3iHnue1zSSGmNOOCh1GLJ7rUKVwV2HvijphGQS2UmhUZewS9VgvxYIdgr+fG1A==",
			"dev": true
		  },
		  "ts-jest": {
			"version": "26.1.0",
			"resolved": "https://registry.npmjs.org/ts-jest/-/ts-jest-26.1.0.tgz",
			"integrity": "sha512-JbhQdyDMYN5nfKXaAwCIyaWLGwevcT2/dbqRPsQeh6NZPUuXjZQZEfeLb75tz0ubCIgEELNm6xAzTe5NXs5Y4Q==",
			"dev": true,
			"requires": {
			  "bs-logger": "0.x",
			  "buffer-from": "1.x",
			  "fast-json-stable-stringify": "2.x",
			  "json5": "2.x",
			  "lodash.memoize": "4.x",
			  "make-error": "1.x",
			  "micromatch": "4.x",
			  "mkdirp": "1.x",
			  "semver": "7.x",
			  "yargs-parser": "18.x"
			},
			"dependencies": {
			  "json5": {
				"version": "2.1.3",
				"resolved": "https://registry.npmjs.org/json5/-/json5-2.1.3.tgz",
				"integrity": "sha512-KXPvOm8K9IJKFM0bmdn8QXh7udDh1g/giieX0NLCaMnb4hEiVFqnop2ImTXCc5e0/oHz3LTqmHGtExn5hfMkOA==",
				"dev": true,
				"requires": {
				  "minimist": "^1.2.5"
				}
			  },
			  "micromatch": {
				"version": "4.0.2",
				"resolved": "https://registry.npmjs.org/micromatch/-/micromatch-4.0.2.tgz",
				"integrity": "sha512-y7FpHSbMUMoyPbYUSzO6PaZ6FyRnQOpHuKwbo1G+Knck95XVU4QAiKdGEnj5wwoS7PlOgthX/09u5iFJ+aYf5Q==",
				"dev": true,
				"requires": {
				  "braces": "^3.0.1",
				  "picomatch": "^2.0.5"
				}
			  },
			  "mkdirp": {
				"version": "1.0.4",
				"resolved": "https://registry.npmjs.org/mkdirp/-/mkdirp-1.0.4.tgz",
				"integrity": "sha512-vVqVZQyf3WLx2Shd0qJ9xuvqgAyKPLAiqITEtqW0oIUjzo3PePDd6fW9iFz30ef7Ysp/oiWqbhszeGWW2T6Gzw==",
				"dev": true
			  },
			  "semver": {
				"version": "7.3.2",
				"resolved": "https://registry.npmjs.org/semver/-/semver-7.3.2.tgz",
				"integrity": "sha512-OrOb32TeeambH6UrhtShmF7CRDqhL6/5XpPNp2DuRH6+9QLw/orhp72j87v8Qa1ScDkvrrBNpZcDejAirJmfXQ==",
				"dev": true
			  }
			}
		  },
		  "ts-loader": {
			"version": "7.0.5",
			"resolved": "https://registry.npmjs.org/ts-loader/-/ts-loader-7.0.5.tgz",
			"integrity": "sha512-zXypEIT6k3oTc+OZNx/cqElrsbBtYqDknf48OZos0NQ3RTt045fBIU8RRSu+suObBzYB355aIPGOe/3kj9h7Ig==",
			"dev": true,
			"requires": {
			  "chalk": "^2.3.0",
			  "enhanced-resolve": "^4.0.0",
			  "loader-utils": "^1.0.2",
			  "micromatch": "^4.0.0",
			  "semver": "^6.0.0"
			},
			"dependencies": {
			  "micromatch": {
				"version": "4.0.2",
				"resolved": "https://registry.npmjs.org/micromatch/-/micromatch-4.0.2.tgz",
				"integrity": "sha512-y7FpHSbMUMoyPbYUSzO6PaZ6FyRnQOpHuKwbo1G+Knck95XVU4QAiKdGEnj5wwoS7PlOgthX/09u5iFJ+aYf5Q==",
				"dev": true,
				"requires": {
				  "braces": "^3.0.1",
				  "picomatch": "^2.0.5"
				}
			  },
			  "semver": {
				"version": "6.3.0",
				"resolved": "https://registry.npmjs.org/semver/-/semver-6.3.0.tgz",
				"integrity": "sha512-b39TBaTSfV6yBrapU89p5fKekE2m/NwnDocOVruQFS1/veMgdzuPcnOM34M6CwxW8jH/lxEa5rBoDeUwu5HHTw==",
				"dev": true
			  }
			}
		  },
		  "ts-node": {
			"version": "8.10.2",
			"resolved": "https://registry.npmjs.org/ts-node/-/ts-node-8.10.2.tgz",
			"integrity": "sha512-ISJJGgkIpDdBhWVu3jufsWpK3Rzo7bdiIXJjQc0ynKxVOVcg2oIrf2H2cejminGrptVc6q6/uynAHNCuWGbpVA==",
			"dev": true,
			"requires": {
			  "arg": "^4.1.0",
			  "diff": "^4.0.1",
			  "make-error": "^1.1.1",
			  "source-map-support": "^0.5.17",
			  "yn": "3.1.1"
			}
		  },
		  "tsconfig-paths": {
			"version": "3.9.0",
			"resolved": "https://registry.npmjs.org/tsconfig-paths/-/tsconfig-paths-3.9.0.tgz",
			"integrity": "sha512-dRcuzokWhajtZWkQsDVKbWyY+jgcLC5sqJhg2PSgf4ZkH2aHPvaOY8YWGhmjb68b5qqTfasSsDO9k7RUiEmZAw==",
			"dev": true,
			"requires": {
			  "@types/json5": "^0.0.29",
			  "json5": "^1.0.1",
			  "minimist": "^1.2.0",
			  "strip-bom": "^3.0.0"
			},
			"dependencies": {
			  "json5": {
				"version": "1.0.1",
				"resolved": "https://registry.npmjs.org/json5/-/json5-1.0.1.tgz",
				"integrity": "sha512-aKS4WQjPenRxiQsC93MNfjx+nbF4PAdYzmd/1JIj8HYzqfbu86beTuNgXDzPknWk0n0uARlyewZo4s++ES36Ow==",
				"dev": true,
				"requires": {
				  "minimist": "^1.2.0"
				}
			  },
			  "minimist": {
				"version": "1.2.5",
				"resolved": "https://registry.npmjs.org/minimist/-/minimist-1.2.5.tgz",
				"integrity": "sha512-FM9nNUYrRBAELZQT3xeZQ7fmMOBg6nWNmJKTcgsJeaLstP/UODVpGsr5OhXhhXg6f+qtJ8uiZ+PUxkDWcgIXLw==",
				"dev": true
			  }
			}
		  },
		  "tsconfig-paths-webpack-plugin": {
			"version": "3.2.0",
			"resolved": "https://registry.npmjs.org/tsconfig-paths-webpack-plugin/-/tsconfig-paths-webpack-plugin-3.2.0.tgz",
			"integrity": "sha512-S/gOOPOkV8rIL4LurZ1vUdYCVgo15iX9ZMJ6wx6w2OgcpT/G4wMyHB6WM+xheSqGMrWKuxFul+aXpCju3wmj/g==",
			"dev": true,
			"requires": {
			  "chalk": "^2.3.0",
			  "enhanced-resolve": "^4.0.0",
			  "tsconfig-paths": "^3.4.0"
			}
		  },
		  "tslib": {
			"version": "1.10.0",
			"resolved": "https://registry.npmjs.org/tslib/-/tslib-1.10.0.tgz",
			"integrity": "sha512-qOebF53frne81cf0S9B41ByenJ3/IuH8yJKngAX35CmiZySA0khhkovshKK+jGCaMnVomla7gVlIcc3EvKPbTQ=="
		  },
		  "tsutils": {
			"version": "3.17.1",
			"resolved": "https://registry.npmjs.org/tsutils/-/tsutils-3.17.1.tgz",
			"integrity": "sha512-kzeQ5B8H3w60nFY2g8cJIuH7JDpsALXySGtwGJ0p2LSjLgay3NdIpqq5SoOBe46bKDW2iq25irHCr8wjomUS2g==",
			"dev": true,
			"requires": {
			  "tslib": "^1.8.1"
			}
		  },
		  "tty-browserify": {
			"version": "0.0.0",
			"resolved": "https://registry.npmjs.org/tty-browserify/-/tty-browserify-0.0.0.tgz",
			"integrity": "sha1-oVe6QC2iTpv5V/mqadUk7tQpAaY=",
			"dev": true
		  },
		  "tunnel-agent": {
			"version": "0.6.0",
			"resolved": "https://registry.npmjs.org/tunnel-agent/-/tunnel-agent-0.6.0.tgz",
			"integrity": "sha1-J6XeoGs2sEoKmWZ3SykIaPD8QP0=",
			"dev": true,
			"requires": {
			  "safe-buffer": "^5.0.1"
			}
		  },
		  "tweetnacl": {
			"version": "0.14.5",
			"resolved": "https://registry.npmjs.org/tweetnacl/-/tweetnacl-0.14.5.tgz",
			"integrity": "sha1-WuaBd/GS1EViadEIr6k/+HQ/T2Q=",
			"dev": true
		  },
		  "type": {
			"version": "1.2.0",
			"resolved": "https://registry.npmjs.org/type/-/type-1.2.0.tgz",
			"integrity": "sha512-+5nt5AAniqsCnu2cEQQdpzCAh33kVx8n0VoFidKpB1dVVLAN/F+bgVOqOJqOnEnrhp222clB5p3vUlD+1QAnfg=="
		  },
		  "type-check": {
			"version": "0.4.0",
			"resolved": "https://registry.npmjs.org/type-check/-/type-check-0.4.0.tgz",
			"integrity": "sha512-XleUoc9uwGXqjWwXaUTZAmzMcFZ5858QA2vvx1Ur5xIcixXIP+8LnFDgRplU30us6teqdlskFfu+ae4K79Ooew==",
			"dev": true,
			"requires": {
			  "prelude-ls": "^1.2.1"
			}
		  },
		  "type-detect": {
			"version": "4.0.8",
			"resolved": "https://registry.npmjs.org/type-detect/-/type-detect-4.0.8.tgz",
			"integrity": "sha512-0fr/mIH1dlO+x7TlcMy+bIDqKPsw/70tVyeHW787goQjhmqaZe10uwLujubK9q9Lg6Fiho1KUKDYz0Z7k7g5/g==",
			"dev": true
		  },
		  "type-fest": {
			"version": "0.11.0",
			"resolved": "https://registry.npmjs.org/type-fest/-/type-fest-0.11.0.tgz",
			"integrity": "sha512-OdjXJxnCN1AvyLSzeKIgXTXxV+99ZuXl3Hpo9XpJAv9MBcHrrJOQ5kV7ypXOuQie+AmWG25hLbiKdwYTifzcfQ==",
			"dev": true
		  },
		  "type-is": {
			"version": "1.6.18",
			"resolved": "https://registry.npmjs.org/type-is/-/type-is-1.6.18.tgz",
			"integrity": "sha512-TkRKr9sUTxEH8MdfuCSP7VizJyzRNMjj2J2do2Jr3Kym598JVdEksuzPQCnlFPW4ky9Q+iA+ma9BGm06XQBy8g==",
			"requires": {
			  "media-typer": "0.3.0",
			  "mime-types": "~2.1.24"
			}
		  },
		  "typedarray": {
			"version": "0.0.6",
			"resolved": "https://registry.npmjs.org/typedarray/-/typedarray-0.0.6.tgz",
			"integrity": "sha1-hnrHTjhkGHsdPUfZlqeOxciDB3c="
		  },
		  "typedarray-to-buffer": {
			"version": "3.1.5",
			"resolved": "https://registry.npmjs.org/typedarray-to-buffer/-/typedarray-to-buffer-3.1.5.tgz",
			"integrity": "sha512-zdu8XMNEDepKKR+XYOXAVPtWui0ly0NtohUscw+UmaHiAWT8hrV1rr//H6V+0DvJ3OQ19S979M0laLfX8rm82Q==",
			"dev": true,
			"requires": {
			  "is-typedarray": "^1.0.0"
			}
		  },
		  "typescript": {
			"version": "3.9.3",
			"resolved": "https://registry.npmjs.org/typescript/-/typescript-3.9.3.tgz",
			"integrity": "sha512-D/wqnB2xzNFIcoBG9FG8cXRDjiqSTbG2wd8DMZeQyJlP1vfTkIxH4GKveWaEBYySKIg+USu+E+EDIR47SqnaMQ==",
			"dev": true
		  },
		  "union-value": {
			"version": "1.0.1",
			"resolved": "https://registry.npmjs.org/union-value/-/union-value-1.0.1.tgz",
			"integrity": "sha512-tJfXmxMeWYnczCVs7XAEvIV7ieppALdyepWMkHkwciRpZraG/xwT+s2JN8+pr1+8jCRf80FFzvr+MpQeeoF4Xg==",
			"dev": true,
			"requires": {
			  "arr-union": "^3.1.0",
			  "get-value": "^2.0.6",
			  "is-extendable": "^0.1.1",
			  "set-value": "^2.0.1"
			}
		  },
		  "unique-filename": {
			"version": "1.1.1",
			"resolved": "https://registry.npmjs.org/unique-filename/-/unique-filename-1.1.1.tgz",
			"integrity": "sha512-Vmp0jIp2ln35UTXuryvjzkjGdRyf9b2lTXuSYUiPmzRcl3FDtYqAwOnTJkAngD9SWhnoJzDbTKwaOrZ+STtxNQ==",
			"dev": true,
			"requires": {
			  "unique-slug": "^2.0.0"
			}
		  },
		  "unique-slug": {
			"version": "2.0.2",
			"resolved": "https://registry.npmjs.org/unique-slug/-/unique-slug-2.0.2.tgz",
			"integrity": "sha512-zoWr9ObaxALD3DOPfjPSqxt4fnZiWblxHIgeWqW8x7UqDzEtHEQLzji2cuJYQFCU6KmoJikOYAZlrTHHebjx2w==",
			"dev": true,
			"requires": {
			  "imurmurhash": "^0.1.4"
			}
		  },
		  "universalify": {
			"version": "0.1.2",
			"resolved": "https://registry.npmjs.org/universalify/-/universalify-0.1.2.tgz",
			"integrity": "sha512-rBJeI5CXAlmy1pV+617WB9J63U6XcazHHF2f2dbJix4XzpUF0RS3Zbj0FGIOCAva5P/d/GBOYaACQ1w+0azUkg==",
			"dev": true
		  },
		  "unpipe": {
			"version": "1.0.0",
			"resolved": "https://registry.npmjs.org/unpipe/-/unpipe-1.0.0.tgz",
			"integrity": "sha1-sr9O6FFKrmFltIF4KdIbLvSZBOw="
		  },
		  "unset-value": {
			"version": "1.0.0",
			"resolved": "https://registry.npmjs.org/unset-value/-/unset-value-1.0.0.tgz",
			"integrity": "sha1-g3aHP30jNRef+x5vw6jtDfyKtVk=",
			"dev": true,
			"requires": {
			  "has-value": "^0.3.1",
			  "isobject": "^3.0.0"
			},
			"dependencies": {
			  "has-value": {
				"version": "0.3.1",
				"resolved": "https://registry.npmjs.org/has-value/-/has-value-0.3.1.tgz",
				"integrity": "sha1-ex9YutpiyoJ+wKIHgCVlSEWZXh8=",
				"dev": true,
				"requires": {
				  "get-value": "^2.0.3",
				  "has-values": "^0.1.4",
				  "isobject": "^2.0.0"
				},
				"dependencies": {
				  "isobject": {
					"version": "2.1.0",
					"resolved": "https://registry.npmjs.org/isobject/-/isobject-2.1.0.tgz",
					"integrity": "sha1-8GVWEJaj8dou9GJy+BXIQNh+DIk=",
					"dev": true,
					"requires": {
					  "isarray": "1.0.0"
					}
				  }
				}
			  },
			  "has-values": {
				"version": "0.1.4",
				"resolved": "https://registry.npmjs.org/has-values/-/has-values-0.1.4.tgz",
				"integrity": "sha1-bWHeldkd/Km5oCCJrThL/49it3E=",
				"dev": true
			  },
			  "isarray": {
				"version": "1.0.0",
				"resolved": "https://registry.npmjs.org/isarray/-/isarray-1.0.0.tgz",
				"integrity": "sha1-u5NdSFgsuhaMBoNJV6VKPgcSTxE=",
				"dev": true
			  }
			}
		  },
		  "upath": {
			"version": "1.2.0",
			"resolved": "https://registry.npmjs.org/upath/-/upath-1.2.0.tgz",
			"integrity": "sha512-aZwGpamFO61g3OlfT7OQCHqhGnW43ieH9WZeP7QxN/G/jS4jfqUkZxoryvJgVPEcrl5NL/ggHsSmLMHuH64Lhg==",
			"dev": true,
			"optional": true
		  },
		  "uri-js": {
			"version": "4.2.2",
			"resolved": "https://registry.npmjs.org/uri-js/-/uri-js-4.2.2.tgz",
			"integrity": "sha512-KY9Frmirql91X2Qgjry0Wd4Y+YTdrdZheS8TFwvkbLWf/G5KNJDCh6pKL5OZctEW4+0Baa5idK2ZQuELRwPznQ==",
			"dev": true,
			"requires": {
			  "punycode": "^2.1.0"
			}
		  },
		  "urix": {
			"version": "0.1.0",
			"resolved": "https://registry.npmjs.org/urix/-/urix-0.1.0.tgz",
			"integrity": "sha1-2pN/emLiH+wf0Y1Js1wpNQZ6bHI=",
			"dev": true
		  },
		  "url": {
			"version": "0.11.0",
			"resolved": "https://registry.npmjs.org/url/-/url-0.11.0.tgz",
			"integrity": "sha1-ODjpfPxgUh63PFJajlW/3Z4uKPE=",
			"dev": true,
			"requires": {
			  "punycode": "1.3.2",
			  "querystring": "0.2.0"
			},
			"dependencies": {
			  "punycode": {
				"version": "1.3.2",
				"resolved": "https://registry.npmjs.org/punycode/-/punycode-1.3.2.tgz",
				"integrity": "sha1-llOgNvt8HuQjQvIyXM7v6jkmxI0=",
				"dev": true
			  }
			}
		  },
		  "use": {
			"version": "3.1.1",
			"resolved": "https://registry.npmjs.org/use/-/use-3.1.1.tgz",
			"integrity": "sha512-cwESVXlO3url9YWlFW/TA9cshCEhtu7IKJ/p5soJ/gGpj7vbvFrAY/eIioQ6Dw23KjZhYgiIo8HOs1nQ2vr/oQ==",
			"dev": true
		  },
		  "util": {
			"version": "0.11.1",
			"resolved": "https://registry.npmjs.org/util/-/util-0.11.1.tgz",
			"integrity": "sha512-HShAsny+zS2TZfaXxD9tYj4HQGlBezXZMZuM/S5PKLLoZkShZiGk9o5CzukI1LVHZvjdvZ2Sj1aW/Ndn2NB/HQ==",
			"dev": true,
			"requires": {
			  "inherits": "2.0.3"
			}
		  },
		  "util-deprecate": {
			"version": "1.0.2",
			"resolved": "https://registry.npmjs.org/util-deprecate/-/util-deprecate-1.0.2.tgz",
			"integrity": "sha1-RQ1Nyfpw3nMnYvvS1KKJgUGaDM8="
		  },
		  "utils-merge": {
			"version": "1.0.1",
			"resolved": "https://registry.npmjs.org/utils-merge/-/utils-merge-1.0.1.tgz",
			"integrity": "sha1-n5VxD1CiZ5R7LMwSR0HBAoQn5xM="
		  },
		  "uuid": {
			"version": "7.0.3",
			"resolved": "https://registry.npmjs.org/uuid/-/uuid-7.0.3.tgz",
			"integrity": "sha512-DPSke0pXhTZgoF/d+WSt2QaKMCFSfx7QegxEWT+JOuHF5aWrKEn0G+ztjuJg/gG8/ItK+rbPCD/yNv8yyih6Cg==",
			"dev": true,
			"optional": true
		  },
		  "v8-compile-cache": {
			"version": "2.1.1",
			"resolved": "https://registry.npmjs.org/v8-compile-cache/-/v8-compile-cache-2.1.1.tgz",
			"integrity": "sha512-8OQ9CL+VWyt3JStj7HX7/ciTL2V3Rl1Wf5OL+SNTm0yK1KvtReVulksyeRnCANHHuUxHlQig+JJDlUhBt1NQDQ==",
			"dev": true
		  },
		  "v8-to-istanbul": {
			"version": "4.1.4",
			"resolved": "https://registry.npmjs.org/v8-to-istanbul/-/v8-to-istanbul-4.1.4.tgz",
			"integrity": "sha512-Rw6vJHj1mbdK8edjR7+zuJrpDtKIgNdAvTSAcpYfgMIw+u2dPDntD3dgN4XQFLU2/fvFQdzj+EeSGfd/jnY5fQ==",
			"dev": true,
			"requires": {
			  "@types/istanbul-lib-coverage": "^2.0.1",
			  "convert-source-map": "^1.6.0",
			  "source-map": "^0.7.3"
			}
		  },
		  "validate-npm-package-license": {
			"version": "3.0.4",
			"resolved": "https://registry.npmjs.org/validate-npm-package-license/-/validate-npm-package-license-3.0.4.tgz",
			"integrity": "sha512-DpKm2Ui/xN7/HQKCtpZxoRWBhZ9Z0kqtygG8XCgNQ8ZlDnxuQmWhj566j8fN4Cu3/JmbhsDo7fcAJq4s9h27Ew==",
			"dev": true,
			"requires": {
			  "spdx-correct": "^3.0.0",
			  "spdx-expression-parse": "^3.0.0"
			}
		  },
		  "vary": {
			"version": "1.1.2",
			"resolved": "https://registry.npmjs.org/vary/-/vary-1.1.2.tgz",
			"integrity": "sha1-IpnwLG3tMNSllhsLn3RSShj2NPw="
		  },
		  "verror": {
			"version": "1.10.0",
			"resolved": "https://registry.npmjs.org/verror/-/verror-1.10.0.tgz",
			"integrity": "sha1-OhBcoXBTr1XW4nDB+CiGguGNpAA=",
			"dev": true,
			"requires": {
			  "assert-plus": "^1.0.0",
			  "core-util-is": "1.0.2",
			  "extsprintf": "^1.2.0"
			}
		  },
		  "vm-browserify": {
			"version": "1.1.2",
			"resolved": "https://registry.npmjs.org/vm-browserify/-/vm-browserify-1.1.2.tgz",
			"integrity": "sha512-2ham8XPWTONajOR0ohOKOHXkm3+gaBmGut3SRuu75xLd/RRaY6vqgh8NBYYk7+RW3u5AtzPQZG8F10LHkl0lAQ==",
			"dev": true
		  },
		  "w3c-hr-time": {
			"version": "1.0.2",
			"resolved": "https://registry.npmjs.org/w3c-hr-time/-/w3c-hr-time-1.0.2.tgz",
			"integrity": "sha512-z8P5DvDNjKDoFIHK7q8r8lackT6l+jo/Ye3HOle7l9nICP9lf1Ci25fy9vHd0JOWewkIFzXIEig3TdKT7JQ5fQ==",
			"dev": true,
			"requires": {
			  "browser-process-hrtime": "^1.0.0"
			}
		  },
		  "w3c-xmlserializer": {
			"version": "2.0.0",
			"resolved": "https://registry.npmjs.org/w3c-xmlserializer/-/w3c-xmlserializer-2.0.0.tgz",
			"integrity": "sha512-4tzD0mF8iSiMiNs30BiLO3EpfGLZUT2MSX/G+o7ZywDzliWQ3OPtTZ0PTC3B3ca1UAf4cJMHB+2Bf56EriJuRA==",
			"dev": true,
			"requires": {
			  "xml-name-validator": "^3.0.0"
			}
		  },
		  "walker": {
			"version": "1.0.7",
			"resolved": "https://registry.npmjs.org/walker/-/walker-1.0.7.tgz",
			"integrity": "sha1-L3+bj9ENZ3JisYqITijRlhjgKPs=",
			"dev": true,
			"requires": {
			  "makeerror": "1.0.x"
			}
		  },
		  "watchpack": {
			"version": "1.7.2",
			"resolved": "https://registry.npmjs.org/watchpack/-/watchpack-1.7.2.tgz",
			"integrity": "sha512-ymVbbQP40MFTp+cNMvpyBpBtygHnPzPkHqoIwRRj/0B8KhqQwV8LaKjtbaxF2lK4vl8zN9wCxS46IFCU5K4W0g==",
			"dev": true,
			"requires": {
			  "chokidar": "^3.4.0",
			  "graceful-fs": "^4.1.2",
			  "neo-async": "^2.5.0",
			  "watchpack-chokidar2": "^2.0.0"
			}
		  },
		  "watchpack-chokidar2": {
			"version": "2.0.0",
			"resolved": "https://registry.npmjs.org/watchpack-chokidar2/-/watchpack-chokidar2-2.0.0.tgz",
			"integrity": "sha512-9TyfOyN/zLUbA288wZ8IsMZ+6cbzvsNyEzSBp6e/zkifi6xxbl8SmQ/CxQq32k8NNqrdVEVUVSEf56L4rQ/ZxA==",
			"dev": true,
			"optional": true,
			"requires": {
			  "chokidar": "^2.1.8"
			},
			"dependencies": {
			  "anymatch": {
				"version": "2.0.0",
				"resolved": "https://registry.npmjs.org/anymatch/-/anymatch-2.0.0.tgz",
				"integrity": "sha512-5teOsQWABXHHBFP9y3skS5P3d/WfWXpv3FUpy+LorMrNYaT9pI4oLMQX7jzQ2KklNpGpWHzdCXTDT2Y3XGlZBw==",
				"dev": true,
				"optional": true,
				"requires": {
				  "micromatch": "^3.1.4",
				  "normalize-path": "^2.1.1"
				},
				"dependencies": {
				  "normalize-path": {
					"version": "2.1.1",
					"resolved": "https://registry.npmjs.org/normalize-path/-/normalize-path-2.1.1.tgz",
					"integrity": "sha1-GrKLVW4Zg2Oowab35vogE3/mrtk=",
					"dev": true,
					"optional": true,
					"requires": {
					  "remove-trailing-separator": "^1.0.1"
					}
				  }
				}
			  },
			  "binary-extensions": {
				"version": "1.13.1",
				"resolved": "https://registry.npmjs.org/binary-extensions/-/binary-extensions-1.13.1.tgz",
				"integrity": "sha512-Un7MIEDdUC5gNpcGDV97op1Ywk748MpHcFTHoYs6qnj1Z3j7I53VG3nwZhKzoBZmbdRNnb6WRdFlwl7tSDuZGw==",
				"dev": true,
				"optional": true
			  },
			  "braces": {
				"version": "2.3.2",
				"resolved": "https://registry.npmjs.org/braces/-/braces-2.3.2.tgz",
				"integrity": "sha512-aNdbnj9P8PjdXU4ybaWLK2IF3jc/EoDYbC7AazW6to3TRsfXxscC9UXOB5iDiEQrkyIbWp2SLQda4+QAa7nc3w==",
				"dev": true,
				"optional": true,
				"requires": {
				  "arr-flatten": "^1.1.0",
				  "array-unique": "^0.3.2",
				  "extend-shallow": "^2.0.1",
				  "fill-range": "^4.0.0",
				  "isobject": "^3.0.1",
				  "repeat-element": "^1.1.2",
				  "snapdragon": "^0.8.1",
				  "snapdragon-node": "^2.0.1",
				  "split-string": "^3.0.2",
				  "to-regex": "^3.0.1"
				}
			  },
			  "chokidar": {
				"version": "2.1.8",
				"resolved": "https://registry.npmjs.org/chokidar/-/chokidar-2.1.8.tgz",
				"integrity": "sha512-ZmZUazfOzf0Nve7duiCKD23PFSCs4JPoYyccjUFF3aQkQadqBhfzhjkwBH2mNOG9cTBwhamM37EIsIkZw3nRgg==",
				"dev": true,
				"optional": true,
				"requires": {
				  "anymatch": "^2.0.0",
				  "async-each": "^1.0.1",
				  "braces": "^2.3.2",
				  "fsevents": "^1.2.7",
				  "glob-parent": "^3.1.0",
				  "inherits": "^2.0.3",
				  "is-binary-path": "^1.0.0",
				  "is-glob": "^4.0.0",
				  "normalize-path": "^3.0.0",
				  "path-is-absolute": "^1.0.0",
				  "readdirp": "^2.2.1",
				  "upath": "^1.1.1"
				}
			  },
			  "extend-shallow": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/extend-shallow/-/extend-shallow-2.0.1.tgz",
				"integrity": "sha1-Ua99YUrZqfYQ6huvu5idaxxWiQ8=",
				"dev": true,
				"optional": true,
				"requires": {
				  "is-extendable": "^0.1.0"
				}
			  },
			  "fill-range": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/fill-range/-/fill-range-4.0.0.tgz",
				"integrity": "sha1-1USBHUKPmOsGpj3EAtJAPDKMOPc=",
				"dev": true,
				"optional": true,
				"requires": {
				  "extend-shallow": "^2.0.1",
				  "is-number": "^3.0.0",
				  "repeat-string": "^1.6.1",
				  "to-regex-range": "^2.1.0"
				}
			  },
			  "fsevents": {
				"version": "1.2.13",
				"resolved": "https://registry.npmjs.org/fsevents/-/fsevents-1.2.13.tgz",
				"integrity": "sha512-oWb1Z6mkHIskLzEJ/XWX0srkpkTQ7vaopMQkyaEIoq0fmtFVxOthb8cCxeT+p3ynTdkk/RZwbgG4brR5BeWECw==",
				"dev": true,
				"optional": true,
				"requires": {
				  "bindings": "^1.5.0",
				  "nan": "^2.12.1"
				}
			  },
			  "glob-parent": {
				"version": "3.1.0",
				"resolved": "https://registry.npmjs.org/glob-parent/-/glob-parent-3.1.0.tgz",
				"integrity": "sha1-nmr2KZ2NO9K9QEMIMr0RPfkGxa4=",
				"dev": true,
				"optional": true,
				"requires": {
				  "is-glob": "^3.1.0",
				  "path-dirname": "^1.0.0"
				},
				"dependencies": {
				  "is-glob": {
					"version": "3.1.0",
					"resolved": "https://registry.npmjs.org/is-glob/-/is-glob-3.1.0.tgz",
					"integrity": "sha1-e6WuJCF4BKxwcHuWkiVnSGzD6Eo=",
					"dev": true,
					"optional": true,
					"requires": {
					  "is-extglob": "^2.1.0"
					}
				  }
				}
			  },
			  "is-binary-path": {
				"version": "1.0.1",
				"resolved": "https://registry.npmjs.org/is-binary-path/-/is-binary-path-1.0.1.tgz",
				"integrity": "sha1-dfFmQrSA8YenEcgUFh/TpKdlWJg=",
				"dev": true,
				"optional": true,
				"requires": {
				  "binary-extensions": "^1.0.0"
				}
			  },
			  "is-number": {
				"version": "3.0.0",
				"resolved": "https://registry.npmjs.org/is-number/-/is-number-3.0.0.tgz",
				"integrity": "sha1-JP1iAaR4LPUFYcgQJ2r8fRLXEZU=",
				"dev": true,
				"optional": true,
				"requires": {
				  "kind-of": "^3.0.2"
				}
			  },
			  "isarray": {
				"version": "1.0.0",
				"resolved": "https://registry.npmjs.org/isarray/-/isarray-1.0.0.tgz",
				"integrity": "sha1-u5NdSFgsuhaMBoNJV6VKPgcSTxE=",
				"dev": true,
				"optional": true
			  },
			  "kind-of": {
				"version": "3.2.2",
				"resolved": "https://registry.npmjs.org/kind-of/-/kind-of-3.2.2.tgz",
				"integrity": "sha1-MeohpzS6ubuw8yRm2JOupR5KPGQ=",
				"dev": true,
				"optional": true,
				"requires": {
				  "is-buffer": "^1.1.5"
				}
			  },
			  "readable-stream": {
				"version": "2.3.7",
				"resolved": "https://registry.npmjs.org/readable-stream/-/readable-stream-2.3.7.tgz",
				"integrity": "sha512-Ebho8K4jIbHAxnuxi7o42OrZgF/ZTNcsZj6nRKyUmkhLFq8CHItp/fy6hQZuZmP/n3yZ9VBUbp4zz/mX8hmYPw==",
				"dev": true,
				"optional": true,
				"requires": {
				  "core-util-is": "~1.0.0",
				  "inherits": "~2.0.3",
				  "isarray": "~1.0.0",
				  "process-nextick-args": "~2.0.0",
				  "safe-buffer": "~5.1.1",
				  "string_decoder": "~1.1.1",
				  "util-deprecate": "~1.0.1"
				}
			  },
			  "readdirp": {
				"version": "2.2.1",
				"resolved": "https://registry.npmjs.org/readdirp/-/readdirp-2.2.1.tgz",
				"integrity": "sha512-1JU/8q+VgFZyxwrJ+SVIOsh+KywWGpds3NTqikiKpDMZWScmAYyKIgqkO+ARvNWJfXeXR1zxz7aHF4u4CyH6vQ==",
				"dev": true,
				"optional": true,
				"requires": {
				  "graceful-fs": "^4.1.11",
				  "micromatch": "^3.1.10",
				  "readable-stream": "^2.0.2"
				}
			  },
			  "string_decoder": {
				"version": "1.1.1",
				"resolved": "https://registry.npmjs.org/string_decoder/-/string_decoder-1.1.1.tgz",
				"integrity": "sha512-n/ShnvDi6FHbbVfviro+WojiFzv+s8MPMHBczVePfUpDJLwoLT0ht1l4YwBCbi8pJAveEEdnkHyPyTP/mzRfwg==",
				"dev": true,
				"optional": true,
				"requires": {
				  "safe-buffer": "~5.1.0"
				}
			  },
			  "to-regex-range": {
				"version": "2.1.1",
				"resolved": "https://registry.npmjs.org/to-regex-range/-/to-regex-range-2.1.1.tgz",
				"integrity": "sha1-fIDBe53+vlmeJzZ+DU3VWQFB2zg=",
				"dev": true,
				"optional": true,
				"requires": {
				  "is-number": "^3.0.0",
				  "repeat-string": "^1.6.1"
				}
			  }
			}
		  },
		  "wcwidth": {
			"version": "1.0.1",
			"resolved": "https://registry.npmjs.org/wcwidth/-/wcwidth-1.0.1.tgz",
			"integrity": "sha1-8LDc+RW8X/FSivrbLA4XtTLaL+g=",
			"dev": true,
			"requires": {
			  "defaults": "^1.0.3"
			}
		  },
		  "webidl-conversions": {
			"version": "6.1.0",
			"resolved": "https://registry.npmjs.org/webidl-conversions/-/webidl-conversions-6.1.0.tgz",
			"integrity": "sha512-qBIvFLGiBpLjfwmYAaHPXsn+ho5xZnGvyGvsarywGNc8VyQJUMHJ8OBKGGrPER0okBeMDaan4mNBlgBROxuI8w==",
			"dev": true
		  },
		  "webpack": {
			"version": "4.43.0",
			"resolved": "https://registry.npmjs.org/webpack/-/webpack-4.43.0.tgz",
			"integrity": "sha512-GW1LjnPipFW2Y78OOab8NJlCflB7EFskMih2AHdvjbpKMeDJqEgSx24cXXXiPS65+WSwVyxtDsJH6jGX2czy+g==",
			"dev": true,
			"requires": {
			  "@webassemblyjs/ast": "1.9.0",
			  "@webassemblyjs/helper-module-context": "1.9.0",
			  "@webassemblyjs/wasm-edit": "1.9.0",
			  "@webassemblyjs/wasm-parser": "1.9.0",
			  "acorn": "^6.4.1",
			  "ajv": "^6.10.2",
			  "ajv-keywords": "^3.4.1",
			  "chrome-trace-event": "^1.0.2",
			  "enhanced-resolve": "^4.1.0",
			  "eslint-scope": "^4.0.3",
			  "json-parse-better-errors": "^1.0.2",
			  "loader-runner": "^2.4.0",
			  "loader-utils": "^1.2.3",
			  "memory-fs": "^0.4.1",
			  "micromatch": "^3.1.10",
			  "mkdirp": "^0.5.3",
			  "neo-async": "^2.6.1",
			  "node-libs-browser": "^2.2.1",
			  "schema-utils": "^1.0.0",
			  "tapable": "^1.1.3",
			  "terser-webpack-plugin": "^1.4.3",
			  "watchpack": "^1.6.1",
			  "webpack-sources": "^1.4.1"
			},
			"dependencies": {
			  "isarray": {
				"version": "1.0.0",
				"resolved": "https://registry.npmjs.org/isarray/-/isarray-1.0.0.tgz",
				"integrity": "sha1-u5NdSFgsuhaMBoNJV6VKPgcSTxE=",
				"dev": true
			  },
			  "memory-fs": {
				"version": "0.4.1",
				"resolved": "https://registry.npmjs.org/memory-fs/-/memory-fs-0.4.1.tgz",
				"integrity": "sha1-OpoguEYlI+RHz7x+i7gO1me/xVI=",
				"dev": true,
				"requires": {
				  "errno": "^0.1.3",
				  "readable-stream": "^2.0.1"
				}
			  },
			  "readable-stream": {
				"version": "2.3.7",
				"resolved": "https://registry.npmjs.org/readable-stream/-/readable-stream-2.3.7.tgz",
				"integrity": "sha512-Ebho8K4jIbHAxnuxi7o42OrZgF/ZTNcsZj6nRKyUmkhLFq8CHItp/fy6hQZuZmP/n3yZ9VBUbp4zz/mX8hmYPw==",
				"dev": true,
				"requires": {
				  "core-util-is": "~1.0.0",
				  "inherits": "~2.0.3",
				  "isarray": "~1.0.0",
				  "process-nextick-args": "~2.0.0",
				  "safe-buffer": "~5.1.1",
				  "string_decoder": "~1.1.1",
				  "util-deprecate": "~1.0.1"
				}
			  },
			  "string_decoder": {
				"version": "1.1.1",
				"resolved": "https://registry.npmjs.org/string_decoder/-/string_decoder-1.1.1.tgz",
				"integrity": "sha512-n/ShnvDi6FHbbVfviro+WojiFzv+s8MPMHBczVePfUpDJLwoLT0ht1l4YwBCbi8pJAveEEdnkHyPyTP/mzRfwg==",
				"dev": true,
				"requires": {
				  "safe-buffer": "~5.1.0"
				}
			  }
			}
		  },
		  "webpack-node-externals": {
			"version": "1.7.2",
			"resolved": "https://registry.npmjs.org/webpack-node-externals/-/webpack-node-externals-1.7.2.tgz",
			"integrity": "sha512-ajerHZ+BJKeCLviLUUmnyd5B4RavLF76uv3cs6KNuO8W+HuQaEs0y0L7o40NQxdPy5w0pcv8Ew7yPUAQG0UdCg==",
			"dev": true
		  },
		  "webpack-sources": {
			"version": "1.4.3",
			"resolved": "https://registry.npmjs.org/webpack-sources/-/webpack-sources-1.4.3.tgz",
			"integrity": "sha512-lgTS3Xhv1lCOKo7SA5TjKXMjpSM4sBjNV5+q2bqesbSPs5FjGmU6jjtBSkX9b4qW87vDIsCIlUPOEhbZrMdjeQ==",
			"dev": true,
			"requires": {
			  "source-list-map": "^2.0.0",
			  "source-map": "~0.6.1"
			},
			"dependencies": {
			  "source-map": {
				"version": "0.6.1",
				"resolved": "https://registry.npmjs.org/source-map/-/source-map-0.6.1.tgz",
				"integrity": "sha512-UjgapumWlbMhkBgzT7Ykc5YXUT46F0iKu8SGXq0bcwP5dz/h0Plj6enJqjz1Zbq2l5WaqYnrVbwWOWMyF3F47g==",
				"dev": true
			  }
			}
		  },
		  "whatwg-encoding": {
			"version": "1.0.5",
			"resolved": "https://registry.npmjs.org/whatwg-encoding/-/whatwg-encoding-1.0.5.tgz",
			"integrity": "sha512-b5lim54JOPN9HtzvK9HFXvBma/rnfFeqsic0hSpjtDbVxR3dJKLc+KB4V6GgiGOvl7CY/KNh8rxSo9DKQrnUEw==",
			"dev": true,
			"requires": {
			  "iconv-lite": "0.4.24"
			}
		  },
		  "whatwg-mimetype": {
			"version": "2.3.0",
			"resolved": "https://registry.npmjs.org/whatwg-mimetype/-/whatwg-mimetype-2.3.0.tgz",
			"integrity": "sha512-M4yMwr6mAnQz76TbJm914+gPpB/nCwvZbJU28cUD6dR004SAxDLOOSUaB1JDRqLtaOV/vi0IC5lEAGFgrjGv/g==",
			"dev": true
		  },
		  "whatwg-url": {
			"version": "8.1.0",
			"resolved": "https://registry.npmjs.org/whatwg-url/-/whatwg-url-8.1.0.tgz",
			"integrity": "sha512-vEIkwNi9Hqt4TV9RdnaBPNt+E2Sgmo3gePebCRgZ1R7g6d23+53zCTnuB0amKI4AXq6VM8jj2DUAa0S1vjJxkw==",
			"dev": true,
			"requires": {
			  "lodash.sortby": "^4.7.0",
			  "tr46": "^2.0.2",
			  "webidl-conversions": "^5.0.0"
			},
			"dependencies": {
			  "webidl-conversions": {
				"version": "5.0.0",
				"resolved": "https://registry.npmjs.org/webidl-conversions/-/webidl-conversions-5.0.0.tgz",
				"integrity": "sha512-VlZwKPCkYKxQgeSbH5EyngOmRp7Ww7I9rQLERETtf5ofd9pGeswWiOtogpEO850jziPRarreGxn5QIiTqpb2wA==",
				"dev": true
			  }
			}
		  },
		  "which": {
			"version": "1.3.1",
			"resolved": "https://registry.npmjs.org/which/-/which-1.3.1.tgz",
			"integrity": "sha512-HxJdYWq1MTIQbJ3nw0cqssHoTNU267KlrDuGZ1WYlxDStUtKUhOaJmh112/TZmHxxUfuJqPXSOm7tDyas0OSIQ==",
			"dev": true,
			"requires": {
			  "isexe": "^2.0.0"
			}
		  },
		  "which-module": {
			"version": "2.0.0",
			"resolved": "https://registry.npmjs.org/which-module/-/which-module-2.0.0.tgz",
			"integrity": "sha1-2e8H3Od7mQK4o6j6SzHD4/fm6Ho=",
			"dev": true
		  },
		  "windows-release": {
			"version": "3.3.0",
			"resolved": "https://registry.npmjs.org/windows-release/-/windows-release-3.3.0.tgz",
			"integrity": "sha512-2HetyTg1Y+R+rUgrKeUEhAG/ZuOmTrI1NBb3ZyAGQMYmOJjBBPe4MTodghRkmLJZHwkuPi02anbeGP+Zf401LQ==",
			"dev": true,
			"requires": {
			  "execa": "^1.0.0"
			}
		  },
		  "word-wrap": {
			"version": "1.2.3",
			"resolved": "https://registry.npmjs.org/word-wrap/-/word-wrap-1.2.3.tgz",
			"integrity": "sha512-Hz/mrNwitNRh/HUAtM/VT/5VH+ygD6DV7mYKZAtHOrbs8U7lvPS6xf7EJKMF0uW1KJCl0H701g3ZGus+muE5vQ==",
			"dev": true
		  },
		  "worker-farm": {
			"version": "1.7.0",
			"resolved": "https://registry.npmjs.org/worker-farm/-/worker-farm-1.7.0.tgz",
			"integrity": "sha512-rvw3QTZc8lAxyVrqcSGVm5yP/IJ2UcB3U0graE3LCFoZ0Yn2x4EoVSqJKdB/T5M+FLcRPjz4TDacRf3OCfNUzw==",
			"dev": true,
			"requires": {
			  "errno": "~0.1.7"
			}
		  },
		  "worker-rpc": {
			"version": "0.1.1",
			"resolved": "https://registry.npmjs.org/worker-rpc/-/worker-rpc-0.1.1.tgz",
			"integrity": "sha512-P1WjMrUB3qgJNI9jfmpZ/htmBEjFh//6l/5y8SD9hg1Ef5zTTVVoRjTrTEzPrNBQvmhMxkoTsjOXN10GWU7aCg==",
			"dev": true,
			"requires": {
			  "microevent.ts": "~0.1.1"
			}
		  },
		  "wrap-ansi": {
			"version": "6.2.0",
			"resolved": "https://registry.npmjs.org/wrap-ansi/-/wrap-ansi-6.2.0.tgz",
			"integrity": "sha512-r6lPcBGxZXlIcymEu7InxDMhdW0KDxpLgoFLcguasxCaJ/SOIZwINatK9KY/tf+ZrlywOKU0UDj3ATXUBfxJXA==",
			"dev": true,
			"requires": {
			  "ansi-styles": "^4.0.0",
			  "string-width": "^4.1.0",
			  "strip-ansi": "^6.0.0"
			},
			"dependencies": {
			  "ansi-regex": {
				"version": "5.0.0",
				"resolved": "https://registry.npmjs.org/ansi-regex/-/ansi-regex-5.0.0.tgz",
				"integrity": "sha512-bY6fj56OUQ0hU1KjFNDQuJFezqKdrAyFdIevADiqrWHwSlbmBNMHp5ak2f40Pm8JTFyM2mqxkG6ngkHO11f/lg==",
				"dev": true
			  },
			  "ansi-styles": {
				"version": "4.2.1",
				"resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.2.1.tgz",
				"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
				"dev": true,
				"requires": {
				  "@types/color-name": "^1.1.1",
				  "color-convert": "^2.0.1"
				}
			  },
			  "color-convert": {
				"version": "2.0.1",
				"resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
				"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
				"dev": true,
				"requires": {
				  "color-name": "~1.1.4"
				}
			  },
			  "color-name": {
				"version": "1.1.4",
				"resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
				"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA==",
				"dev": true
			  },
			  "strip-ansi": {
				"version": "6.0.0",
				"resolved": "https://registry.npmjs.org/strip-ansi/-/strip-ansi-6.0.0.tgz",
				"integrity": "sha512-AuvKTrTfQNYNIctbR1K/YGTR1756GycPsg7b9bdV9Duqur4gv6aKqHXah67Z8ImS7WEz5QVcOtlfW2rZEugt6w==",
				"dev": true,
				"requires": {
				  "ansi-regex": "^5.0.0"
				}
			  }
			}
		  },
		  "wrappy": {
			"version": "1.0.2",
			"resolved": "https://registry.npmjs.org/wrappy/-/wrappy-1.0.2.tgz",
			"integrity": "sha1-tSQ9jz7BqjXxNkYFvA0QNuMKtp8="
		  },
		  "write": {
			"version": "1.0.3",
			"resolved": "https://registry.npmjs.org/write/-/write-1.0.3.tgz",
			"integrity": "sha512-/lg70HAjtkUgWPVZhZcm+T4hkL8Zbtp1nFNOn3lRrxnlv50SRBv7cR7RqR+GMsd3hUXy9hWBo4CHTbFTcOYwig==",
			"dev": true,
			"requires": {
			  "mkdirp": "^0.5.1"
			}
		  },
		  "write-file-atomic": {
			"version": "3.0.3",
			"resolved": "https://registry.npmjs.org/write-file-atomic/-/write-file-atomic-3.0.3.tgz",
			"integrity": "sha512-AvHcyZ5JnSfq3ioSyjrBkH9yW4m7Ayk8/9My/DD9onKeu/94fwrMocemO2QAJFAlnnDN+ZDS+ZjAR5ua1/PV/Q==",
			"dev": true,
			"requires": {
			  "imurmurhash": "^0.1.4",
			  "is-typedarray": "^1.0.0",
			  "signal-exit": "^3.0.2",
			  "typedarray-to-buffer": "^3.1.5"
			}
		  },
		  "ws": {
			"version": "7.3.0",
			"resolved": "https://registry.npmjs.org/ws/-/ws-7.3.0.tgz",
			"integrity": "sha512-iFtXzngZVXPGgpTlP1rBqsUK82p9tKqsWRPg5L56egiljujJT3vGAYnHANvFxBieXrTFavhzhxW52jnaWV+w2w==",
			"dev": true
		  },
		  "xml-name-validator": {
			"version": "3.0.0",
			"resolved": "https://registry.npmjs.org/xml-name-validator/-/xml-name-validator-3.0.0.tgz",
			"integrity": "sha512-A5CUptxDsvxKJEU3yO6DuWBSJz/qizqzJKOMIfUJHETbBw/sFaDxgd6fxm1ewUaM0jZ444Fc5vC5ROYurg/4Pw==",
			"dev": true
		  },
		  "xmlchars": {
			"version": "2.2.0",
			"resolved": "https://registry.npmjs.org/xmlchars/-/xmlchars-2.2.0.tgz",
			"integrity": "sha512-JZnDKK8B0RCDw84FNdDAIpZK+JuJw+s7Lz8nksI7SIuU3UXJJslUthsi+uWBUYOwPFwW7W7PRLRfUKpxjtjFCw==",
			"dev": true
		  },
		  "xtend": {
			"version": "4.0.2",
			"resolved": "https://registry.npmjs.org/xtend/-/xtend-4.0.2.tgz",
			"integrity": "sha512-LKYU1iAXJXUgAXn9URjiu+MWhyUXHsvfp7mcuYm9dSUKK0/CjtrUwFAxD82/mCWbtLsGjFIad0wIsod4zrTAEQ=="
		  },
		  "y18n": {
			"version": "4.0.0",
			"resolved": "https://registry.npmjs.org/y18n/-/y18n-4.0.0.tgz",
			"integrity": "sha512-r9S/ZyXu/Xu9q1tYlpsLIsa3EeLXXk0VwlxqTcFRfg9EhMW+17kbt9G0NrgCmhGb5vT2hyhJZLfDGx+7+5Uj/w==",
			"dev": true
		  },
		  "yallist": {
			"version": "3.1.1",
			"resolved": "https://registry.npmjs.org/yallist/-/yallist-3.1.1.tgz",
			"integrity": "sha512-a4UGQaWPH59mOXUYnAG2ewncQS4i4F43Tv3JoAM+s2VDAmS9NsK8GpDMLrCHPksFT7h3K6TOoUNn2pb7RoXx4g==",
			"dev": true
		  },
		  "yargs": {
			"version": "15.3.1",
			"resolved": "https://registry.npmjs.org/yargs/-/yargs-15.3.1.tgz",
			"integrity": "sha512-92O1HWEjw27sBfgmXiixJWT5hRBp2eobqXicLtPBIDBhYB+1HpwZlXmbW2luivBJHBzki+7VyCLRtAkScbTBQA==",
			"dev": true,
			"requires": {
			  "cliui": "^6.0.0",
			  "decamelize": "^1.2.0",
			  "find-up": "^4.1.0",
			  "get-caller-file": "^2.0.1",
			  "require-directory": "^2.1.1",
			  "require-main-filename": "^2.0.0",
			  "set-blocking": "^2.0.0",
			  "string-width": "^4.2.0",
			  "which-module": "^2.0.0",
			  "y18n": "^4.0.0",
			  "yargs-parser": "^18.1.1"
			},
			"dependencies": {
			  "find-up": {
				"version": "4.1.0",
				"resolved": "https://registry.npmjs.org/find-up/-/find-up-4.1.0.tgz",
				"integrity": "sha512-PpOwAdQ/YlXQ2vj8a3h8IipDuYRi3wceVQQGYWxNINccq40Anw7BlsEXCMbt1Zt+OLA6Fq9suIpIWD0OsnISlw==",
				"dev": true,
				"requires": {
				  "locate-path": "^5.0.0",
				  "path-exists": "^4.0.0"
				}
			  },
			  "locate-path": {
				"version": "5.0.0",
				"resolved": "https://registry.npmjs.org/locate-path/-/locate-path-5.0.0.tgz",
				"integrity": "sha512-t7hw9pI+WvuwNJXwk5zVHpyhIqzg2qTlklJOf0mVxGSbe3Fp2VieZcduNYjaLDoy6p9uGpQEGWG87WpMKlNq8g==",
				"dev": true,
				"requires": {
				  "p-locate": "^4.1.0"
				}
			  },
			  "p-locate": {
				"version": "4.1.0",
				"resolved": "https://registry.npmjs.org/p-locate/-/p-locate-4.1.0.tgz",
				"integrity": "sha512-R79ZZ/0wAxKGu3oYMlz8jy/kbhsNrS7SKZ7PxEHBgJ5+F2mtFW2fK2cOtBh1cHYkQsbzFV7I+EoRKe6Yt0oK7A==",
				"dev": true,
				"requires": {
				  "p-limit": "^2.2.0"
				}
			  },
			  "path-exists": {
				"version": "4.0.0",
				"resolved": "https://registry.npmjs.org/path-exists/-/path-exists-4.0.0.tgz",
				"integrity": "sha512-ak9Qy5Q7jYb2Wwcey5Fpvg2KoAc/ZIhLSLOSBmRmygPsGwkVVt0fZa0qrtMz+m6tJTAHfZQ8FnmB4MG4LWy7/w==",
				"dev": true
			  }
			}
		  },
		  "yargs-parser": {
			"version": "18.1.3",
			"resolved": "https://registry.npmjs.org/yargs-parser/-/yargs-parser-18.1.3.tgz",
			"integrity": "sha512-o50j0JeToy/4K6OZcaQmW6lyXXKhq7csREXcDwk2omFPJEwUNOVtJKvmDr9EI1fAJZUyZcRF7kxGBWmRXudrCQ==",
			"dev": true,
			"requires": {
			  "camelcase": "^5.0.0",
			  "decamelize": "^1.2.0"
			}
		  },
		  "yn": {
			"version": "3.1.1",
			"resolved": "https://registry.npmjs.org/yn/-/yn-3.1.1.tgz",
			"integrity": "sha512-Ux4ygGWsu2c7isFWe8Yu1YluJmqVhxqK2cLXNQA5AcC3QfbGNpM7fu0Y8b/z16pXLnFxZYvWhd3fhBY9DLmC6Q==",
			"dev": true
		  }
		}
	  }
	  
	`)

	return packageLockJson
}

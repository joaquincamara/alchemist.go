package nestContent

func TsConfig() []byte {
	tsConfig := []byte(`
	{
		"compilerOptions": {
		  "module": "commonjs",
		  "declaration": true,
		  "removeComments": true,
		  "emitDecoratorMetadata": true,
		  "experimentalDecorators": true,
		  "allowSyntheticDefaultImports": true,
		  "target": "es2017",
		  "sourceMap": true,
		  "outDir": "./dist",
		  "baseUrl": "./",
		  "incremental": true
		}
	  }
	  
	`)

	return tsConfig
}

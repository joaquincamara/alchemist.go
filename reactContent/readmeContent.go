package reactContent

func ReadmeContentGenerator() []byte {
	readmeContent := []byte(`{
		"name": "my-app",
		"version": "0.1.0",
		"private": true,
		"dependencies": {
			"@testing-library/jest-dom": "^4.2.4",
			"@testing-library/react": "^9.3.2",
			"@testing-library/user-event": "^7.1.2",
			"react": "^16.13.1",
			"react-dom": "^16.13.1",
			"react-scripts": "3.4.3"
		},
		"scripts": {
			"start": "react-scripts start",
			"build": "react-scripts build",
			"test": "react-scripts test",
			"eject": "react-scripts eject"
		},
		"eslintConfig": {
			"extends": "react-app"
		},
		"browserslist": {
			"production": [
				">0.2%",
				"not dead",
				"not op_mini all"
			],
			"development": [
				"last 1 chrome version",
				"last 1 firefox version",
				"last 1 safari version"
			]
		}
		}
		`)

	return readmeContent

}

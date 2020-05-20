module.exports = {
    "env": {
        "browser": true,
        "es6": true
    },
    'settings': {
        'react': {
            'version': require('./package.json').dependencies.react,
        },
    },
    "extends": [
        "eslint:recommended",
        "plugin:react/recommended"
    ],
    "globals": {
        "Atomics": "readonly",
        "SharedArrayBuffer": "readonly"
    },
    'parser': 'babel-eslint',
    "parserOptions": {
        "ecmaFeatures": {
            "jsx": true
        },
        "ecmaVersion": 2018,
        "sourceType": "module"
    },
    "plugins": ["react"],
    "rules": {
        "arrow-spacing": [
            "error",
            {
                "after": true,
                "before": true
            }
        ],
        "block-scoped-var": "error",
        "block-spacing": "error",
        "brace-style": [
            "error",
            "1tbs"
        ],
        'indent': [
            'error', 2, {
            "FunctionExpression": {"parameters": "first"},
            "FunctionDeclaration": {"parameters": "first"},
            "MemberExpression": 1,
            "SwitchCase": 1,
            "outerIIFEBody": 0,
            "VariableDeclarator": 2
          }
        ],
        "react/react-in-jsx-scope": "off",
        "no-undef": "off",
        "react/prop-types": 0,
        "react/jsx-uses-vars": [
            2
        ],
        "react/jsx-no-undef": "error",
        "no-console": 0,
        "no-unused-vars": "warn",
        "react/jsx-key": "off",
        "no-dupe-keys": "off",
        "react/jsx-filename-extension": [1, { "extensions": [".js", ".jsx"] }],
        "react/prop-types": "off"
    }
};
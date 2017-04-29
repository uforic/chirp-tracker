const path = require('path');

module.exports = {
    entry: "./components/index.jsx",
    output: {
        path: path.resolve(__dirname, "assets"),
        filename: "bundle.js"
    },
    devtool: "source-map",
    module: {
 rules: [
       { test: /\.css$/, loader: 'style-loader!css-loader' },
      {
        test: /\.jsx?$/,
              exclude: /node_modules/,
        use: {
        loader: 'babel-loader',
        options: {
          presets: ['env', 'react']
        }
      }
      }
]
    }
};

const path = require('path')
const ExtractTextPlugin = require('extract-text-webpack-plugin')
const UglifyJSPlugin = require('uglifyjs-webpack-plugin')
const I18nPlugin = require('i18n-webpack-plugin')
const OptimizeCssAssetsPlugin = require('optimize-css-assets-webpack-plugin')

const languages = {
  'en': require('./locals/en.json'),
  'fr': require('./locals/fr.json')
}

let configs = Object.keys(languages).map(function (language) {
  return {
    entry: './src/index.js',
    output: {
      path: path.resolve(__dirname, './public'),
      filename: language + '.output.js'
    },
    module: {
      rules: [
        {
          test: /\.js$/, // files ending with .js
          exclude: /node_modules/, // exclude the node_modules directory
          loader: 'babel-loader' // use this (babel-core) loader
        },
        {
          test: /\.scss$/,
          use: ['css-hot-loader'].concat(ExtractTextPlugin.extract({
            use: ['css-loader', 'sass-loader', 'postcss-loader'],
            fallback: 'style-loader'
          }))
        },
        {
          test: /\.html$/,
          loader: 'html-loader?interpolate'
        }
      ] // end rules
    },
    plugins: [
      new ExtractTextPlugin('styles.css'),
      new I18nPlugin(languages[language])
    ],
    devServer: {
      contentBase: path.resolve(__dirname, './public'),
      historyApiFallback: true,
      inline: true,
      open: true
    },
    devtool: 'eval-source-map'
  }
})

module.exports = configs

if (process.env.NODE_ENV === 'production') {
  for (let config of configs) {
    config.plugins.push(
      new UglifyJSPlugin(),
      new OptimizeCssAssetsPlugin()
    )
  }
}

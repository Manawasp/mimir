const path = require('path')
const ExtractTextPlugin = require('extract-text-webpack-plugin')
const UglifyJSPlugin = require('uglifyjs-webpack-plugin')
const I18nPlugin = require('i18n-webpack-plugin')
const OptimizeCssAssetsPlugin = require('optimize-css-assets-webpack-plugin')
const HtmlWebpackPlugin = require('html-webpack-plugin')

const languages = {
  'en': require('./locals/en.json'),
  'fr': require('./locals/fr.json'),
  'cn': require('./locals/cn.json')
}

module.exports = env => {
  let configs = Object.keys(languages).map(function (language) {
    return {
      entry: './src/index.js',
      output: {
        path: path.resolve(__dirname, './public'),
        filename: 'mimir.' + language + '.js'
      },
      module: {
        rules: [
          {
            test: /\.js$/, // files ending with .js
            exclude: /node_modules/, // exclude the node_modules directory
            loader: 'babel-loader' // use this (babel-core) loader
          },
          {
            test: /\.html$/,
            loader: 'html-loader?interpolate'
          }
        ] // end rules
      },
      plugins: [
        new I18nPlugin(languages[language]),
        new HtmlWebpackPlugin({ // Also generate a test.html
          filename: 'mimir.' + language + '.html',
          template: 'src/mimir.html'
        })
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

  // Generate css
  let scssOptions = {}
  if (process.env.MIMIR_COLOR) {
    scssOptions.data = '$main-color: ' + process.env.MIMIR_COLOR + ';'
  }

  configs.push({
    entry: './src/mimir.scss',
    output: {
      path: path.resolve(__dirname, './public'),
      filename: 'mimir.css'
    },
    module: {
      rules: [
        {
          test: /\.scss$/,
          use: ['css-hot-loader'].concat(ExtractTextPlugin.extract({
            use: ['css-loader',
              {
                loader: 'sass-loader',
                options: scssOptions
              }, 'postcss-loader'],
            fallback: 'style-loader'
          }))
        }
      ]
    },
    plugins: [
      new ExtractTextPlugin('mimir.css')
    ]
  })

  if (env && env.production) {
    for (let config of configs) {
      config.devtool = false
      config.plugins.push(
        new UglifyJSPlugin({
          minimize: true,
          sourceMap: 'hidden-source-map',
          comments: false,
          uglifyOptions: {
            ie8: false,
            ecma: 8,
            mangle: true,
            compress: true,
            output: {
              comments: false,
              beautify: false
            },
            warnings: false
          }
        }),
        new OptimizeCssAssetsPlugin()
      )
    }
  }

  // return configs
  return configs
}

module.exports = {
  webpack(config) {
    config.devtool = 'cheap-module-eval-source-map';
    loaders: [
      { test: /\.(png|jpg)$/, loader: 'url-loader?limit=8192' }
    ]
    return config;
  },
  distDir: 'build',
};

module.exports = {
  webpack(config) {
    config.devtool = 'cheap-module-eval-source-map';
    return config;
  }
};

const path = require("path");

module.exports = {
  pages: {
    index: {
      entry: "src/webapp/main.js"
    }
  },
  configureWebpack: {
    resolve: {
      alias: {
        "@": path.resolve(__dirname, "src/webapp/")
      }
    }
  }
};

//import { request } from 'http';

/*const proxy = require('http-proxy-middleware');
module.exports = {
    devServer: {
        port: 8080,     // 端口np
        proxy: {
          '/apis': {
            target: 'http://api.douban.com/v2', // 你请求的第三方接口
            changeOrigin: true,// 在本地会创建一个虚拟服务端，然后发送请求的数据，并同时接收请求的数据，这样服务端和服务端进行数据的交互就不会有跨域问题
            //ws: true, // proxy websockets
            secure: false,  // 如果是https接口，需要配置这个参数
            pathRewrite:{  // 路径重写，
              '^/apis': '/'  // 替换target中的请求地址，也就是说以后你在请求http://api.douban.com/v2/XXXXX这个地址的时候直接写成/api即可。
            }
          }
        }
    },
    lintOnSave: false,   // 取消 eslint 验证
    runtimeCompiler: true,
    publicPath: '/', // 设置打包文件相对路径
};*/

/*// include dependencies
const express = require('express');
const { createProxyMiddleware } = require('http-proxy-middleware');

// proxy middleware options
const options = {
  target: 'http://api.douban.com/v2', // target host
  changeOrigin: true, // needed for virtual hosted sites
  ws: true, // proxy websockets
  pathRewrite: {
    '^/apis/': '/', // rewrite path
  },
  router: {
    // when request.headers.host == 'dev.localhost:3000',
    // override target 'http://www.example.org' to 'http://localhost:8000'
    //'localhost:3000': 'http://localhost:8080',
  },
};

// create the proxy (without context)
const exampleProxy = createProxyMiddleware(options);

// mount `exampleProxy` in web server
const app = express();
app.use('/api', exampleProxy);
app.listen(3000);
*/

/*const express = require('express');
var app = express()
const { createProxyMiddleware } = require('http-proxy-middleware');

app.middleware = [
      createProxyMiddleware(['/apis/*'], {target: 'http://api.douban.com/v2/', changeOrigin: true,pathRewrite: {
        '^/apis': '', 
      }}),
];

app.use(app.middleware);
//app.listen(3000)*/

module.exports = {
  // 基本路径  
  publicPath: './',
  // 输出路径   
  outputDir: 'dist',
  // 静态资源    
  assetsDir: './',
  // eslint-loader是否在保存时候检查  
  lintOnSave: true,
  // 服务项配置    
  devServer: {
      host: 'localhost',
      port: 8080,
      https: false,
      open: true,
　　 // 设置代理proxy
      proxy: {
         '/apis':{
             'target':'http://localhost:8070/v1',
             changeOrigin:true,    //表示是否跨域，
             pathRewrite:{           //表示需要rewrite重写的
                 '^/apis':'',
             }
         }
      }   
  }
}

/**
 * 最后请求的时候：
  譬如："http://localhost:3000/login"
  现在写成："/api/login".
  用 /api 替换原来的 http://localhost:3000
 */
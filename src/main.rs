use actix_web::{middleware, web, App, HttpServer, Responder};

#[actix_rt::main]
async fn main() -> std::io::Result<()> {
    std::env::set_var("RUST_LOG", "actix_web=info,actix_server=info");
    env_logger::init();

    HttpServer::new(|| {
        App::new()
            .wrap(middleware::Logger::default())
            .service(web::resource("/index.html").route(web::get().to(root)))
    })
    .bind("0.0.0.0:8080")?
    .run()
    .await
}

async fn root() -> impl Responder {
    "Hello, world!"
}

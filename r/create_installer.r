library(RInno)
install_inno()

create_app(
  app_name = "CreateAlluvialApp",
  app_dir = "C:\\Users\\nxl16\\code\\dowser\\r\\alluvial.r",
  pkgs = c("ggplot2", "ggalluvial", "colorspace", "ggrepel"),
  R_version = "4.4"
)

compile_iss()

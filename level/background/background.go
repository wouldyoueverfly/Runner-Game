components {
  id: "background"
  component: "/level/background/background.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"bg\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/background/background.atlas\"\n"
  "}\n"
  ""
  position {
    z: -0.6
  }
}

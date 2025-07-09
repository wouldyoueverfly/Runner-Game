components {
  id: "small_planets"
  component: "/level/background/small_planets.script"
}
embedded_components {
  id: "sprite1"
  type: "sprite"
  data: "default_animation: \"earthlike_planet\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/background/background.atlas\"\n"
  "}\n"
  ""
  position {
    x: 119.0
    y: 139.0
    z: -0.6
  }
  scale {
    x: 0.4
    y: 0.4
  }
}
embedded_components {
  id: "sprite2"
  type: "sprite"
  data: "default_animation: \"gas_planet\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/background/background.atlas\"\n"
  "}\n"
  ""
  position {
    x: 347.0
    y: 444.0
    z: -0.6
  }
  scale {
    x: 0.2
    y: 0.2
  }
}
embedded_components {
  id: "sprite3"
  type: "sprite"
  data: "default_animation: \"ice_planet\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/background/background.atlas\"\n"
  "}\n"
  ""
  position {
    x: 1401.0
    y: 426.0
    z: -0.6
  }
  scale {
    x: 0.6
    y: 0.6
  }
}
embedded_components {
  id: "sprite4"
  type: "sprite"
  data: "default_animation: \"ring_planet\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/background/background.atlas\"\n"
  "}\n"
  ""
  position {
    x: 1004.0
    y: 116.0
    z: -0.6
  }
  scale {
    x: 0.1
    y: 0.1
  }
}

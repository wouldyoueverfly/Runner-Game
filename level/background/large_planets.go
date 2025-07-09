components {
  id: "large_planets"
  component: "/level/background/large_planets.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"ice_planet\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/background/background.atlas\"\n"
  "}\n"
  ""
  position {
    x: 295.0
    y: 449.0
    z: -0.5
  }
  scale {
    x: 1.5
    y: 1.5
  }
}
embedded_components {
  id: "sprite1"
  type: "sprite"
  data: "default_animation: \"gas_planet\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/background/background.atlas\"\n"
  "}\n"
  ""
  position {
    x: 2420.0
    y: 92.0
    z: -0.5
  }
}
embedded_components {
  id: "sprite3"
  type: "sprite"
  data: "default_animation: \"ring_planet\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/background/background.atlas\"\n"
  "}\n"
  ""
  position {
    x: 2648.0
    y: 592.0
    z: -0.5
  }
  scale {
    x: 1.5
    y: 1.5
  }
}
embedded_components {
  id: "sprite4"
  type: "sprite"
  data: "default_animation: \"earthlike_planet\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/background/background.atlas\"\n"
  "}\n"
  ""
  position {
    x: 3856.0
    y: 45.0
    z: -0.5
  }
  scale {
    x: 2.0
    y: 2.0
  }
}

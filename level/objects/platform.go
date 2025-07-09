components {
  id: "platform"
  component: "/level/platform.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"rock_planks\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/level.atlas\"\n"
  "}\n"
  ""
  position {
    z: 0.1
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"geometry\"\n"
  "mask: \"hero\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 193.27869\n"
  "  data: 73.66418\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "sprite1"
  type: "sprite"
  data: "default_animation: \"spikes\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: -205.0
  }
}
embedded_components {
  id: "sprite2"
  type: "sprite"
  data: "default_animation: \"spikes\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: 192.0
  }
  rotation {
    z: 1.0
    w: 6.123234E-17
  }
}
embedded_components {
  id: "sprite3"
  type: "sprite"
  data: "default_animation: \"spikes\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: -99.0
    y: -87.0
  }
  rotation {
    z: 0.70710677
    w: 0.70710677
  }
}
embedded_components {
  id: "sprite4"
  type: "sprite"
  data: "default_animation: \"spikes\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 51.0\n"
  "  y: 131.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: 76.0
    y: -87.0
  }
  rotation {
    z: 0.70710677
    w: 0.70710677
  }
}
embedded_components {
  id: "collisionobject1"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"danger\"\n"
  "mask: \"hero\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: -195.0\n"
  "      y: -1.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 189.0\n"
  "      y: -2.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 3\n"
  "    count: 3\n"
  "  }\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 76.0\n"
  "      y: -74.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 6\n"
  "    count: 3\n"
  "  }\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: -100.0\n"
  "      y: -74.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 9\n"
  "    count: 3\n"
  "  }\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 76.0\n"
  "      y: -74.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 12\n"
  "    count: 3\n"
  "  }\n"
  "  data: 22.898434\n"
  "  data: 68.25427\n"
  "  data: 10.0\n"
  "  data: 24.552155\n"
  "  data: 66.64934\n"
  "  data: 10.0\n"
  "  data: 63.354355\n"
  "  data: 25.09535\n"
  "  data: 10.0\n"
  "  data: 63.354355\n"
  "  data: 25.09535\n"
  "  data: 10.0\n"
  "  data: 63.354355\n"
  "  data: 25.09535\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "coin_factory"
  type: "factory"
  data: "prototype: \"/level/objects/coin.go\"\n"
  ""
}

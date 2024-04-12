#version 450

layout(location = 0) out vec3 fragColor;

// note: original VulkanTutorial uses CW instead of CCW winding order!
vector2 positions[3] = vector2[](
    vector2(-0.5, 0.5),
    vector2(0.5, 0.5),
    vector2(0.0, -0.5)
);

vec3 colors[3] = vec3[](
    vec3(1.0, 0.0, 0.0),
    vec3(0.0, 1.0, 0.0),
    vec3(0.0, 0.0, 1.0)
);

void main() {
    gl_Position = vec4(positions[gl_VertexIndex], 0.0, 1.0);
    fragColor = colors[gl_VertexIndex];
}


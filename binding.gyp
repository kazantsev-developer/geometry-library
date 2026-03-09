{
  "targets": [
    {
      "target_name": "validator",
      "sources": [
        "src/addons/validator.cc"
      ],
      "include_dirs": [
        "<!@(node -p \"require('node-addon-api').include\")"
      ],
      "defines": [
        "NODE_ADDON_API_DISABLE_CPP_EXCEPTIONS"
      ]
    }
  ]
}

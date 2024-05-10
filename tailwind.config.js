/** @type {import('tailwindcss').Config} */
const { iconsPlugin, getIconCollections } = require("@egoist/tailwindcss-icons")

export default {
  content: ["./views/**/*.{html,js,templ,go}"],
  theme: {
    extend: {},
  },
  plugins: [
      require("daisyui"),
        // Iconify plugin
        iconsPlugin({
          // Select the icon collections you want to use
          // You can also ignore this option to automatically discover all individual icon packages you have installed
          // If you install @iconify/json, you should explicitly specify the collections you want to use, like this:
          collections: getIconCollections(["mdi", "lucide"]),
          // If you want to use all icons from @iconify/json, you can do this:
          // collections: getIconCollections("all"),
          // and the more recommended way is to use `dynamicIconsPlugin`, see below.
        }),
        function({ addBase}) {
            addBase({
                'html': {
                    fontSize: "16px", // Default font size, which is typically 16px
                    '@screen md': {
                        fontSize: "14px", // Font size on medium (md) breakpoint, which is typically 14px
                    },
                },
            })
        },
  ],
  daisyui: {
    themes: ["light", "dark"], // false: only light + dark | true: all themes | array: specific themes like this ["light", "dark", "cupcake"]
    darkTheme: "dark", // name of one of the included themes for dark mode
    base: true, // applies background color and foreground color for root element by default
    styled: true, // include daisyUI colors and design decisions for all components
    utils: true, // adds responsive and modifier utility classes
    prefix: "", // prefix for daisyUI classnames (components, modifiers and responsive class names. Not colors)
    logs: true, // Shows info about daisyUI version and used config in the console when building your CSS
    themeRoot: ":root", // The element that receives theme color CSS variables
  },
}


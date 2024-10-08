// module.exports = {
//     mode: "jit",
//     purge: [
//         "./views/**/*.go",
//     ],
//     darkMode: false, // or 'media' or 'class'
//     theme: {
//         extend: {},
//     },
//     variants: {},
//     plugins: [
//         require('daisyui'),
//     ],
// }

module.exports = {
    mode: "jit",
    purge: [
        "./pages/**/*.go",
    ],
    darkMode: false, // or 'media' or 'class'
    theme: {
        extend: {},
    },
    variants: {},
    plugins: [
        require('daisyui'),
    ],
}
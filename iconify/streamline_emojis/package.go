package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Package(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M8.75 43.01a15.25 1.99 0 1 0 30.5 0a15.25 1.99 0 1 0-30.5 0Z" opacity=".15"/><path fill="#f0d5a8" d="m5.78 36l16.74 6.7A4 4 0 0 0 24 43V20.75L4.63 13a2.68 2.68 0 0 0-.52 1.59v19A2.65 2.65 0 0 0 5.78 36Z"/><path fill="#debb7e" d="M5.78 31a2.64 2.64 0 0 1-1.67-2.46v5A2.65 2.65 0 0 0 5.78 36l16.74 6.7A4 4 0 0 0 24 43v-5a4 4 0 0 1-1.48-.29Z"/><path fill="none" stroke="#4f4b45" stroke-linejoin="round" d="m5.78 36l16.74 6.7A4 4 0 0 0 24 43V20.75L4.63 13a2.68 2.68 0 0 0-.52 1.59v19A2.65 2.65 0 0 0 5.78 36Z"/><path fill="#f0d5a8" d="M42.22 36a2.65 2.65 0 0 0 1.67-2.46V14.59a2.68 2.68 0 0 0-.52-1.59L24 20.75V43a4 4 0 0 0 1.48-.29Z"/><path fill="#debb7e" d="m42.22 31.22l-16.74 6.7a4 4 0 0 1-1.48.28V43a4 4 0 0 0 1.48-.29L42.22 36a2.65 2.65 0 0 0 1.67-2.46v-4.78a2.65 2.65 0 0 1-1.67 2.46Z"/><path fill="none" stroke="#4f4b45" stroke-linejoin="round" d="M42.22 36a2.65 2.65 0 0 0 1.67-2.46V14.59a2.68 2.68 0 0 0-.52-1.59L24 20.75V43a4 4 0 0 0 1.48-.29Z"/><path fill="#f0d5a8" d="M42.22 12.12L25.48 5.43a4 4 0 0 0-3 0l-16.7 6.69a2.73 2.73 0 0 0-1.15.88L24 20.75L43.37 13a2.73 2.73 0 0 0-1.15-.88Z"/><path fill="#debb7e" d="m42.22 12.12l-.56-.22l-2-.79l-9.14 3.66a17.65 17.65 0 0 1-13.1 0l-9.11-3.66l-2.53 1a2.73 2.73 0 0 0-1.15.89L24 20.75L43.37 13a2.73 2.73 0 0 0-1.15-.88Z"/><path fill="none" stroke="#4f4b45" stroke-linejoin="round" d="M42.22 12.12L25.48 5.43a4 4 0 0 0-3 0l-16.7 6.69a2.73 2.73 0 0 0-1.15.88L24 20.75L43.37 13a2.73 2.73 0 0 0-1.15-.88Z"/><path fill="#fff" stroke="#4f4b45" stroke-linejoin="round" d="m33.41 27.49l-5.73 2.29v-6.49L33.41 21v6.49z"/><path fill="#fff5e3" d="m17.2 18.03l19.89-7.96l-6.13-2.45l-19.89 7.96l6.13 2.45z"/><path fill="#f7e5c6" d="m17.45 14.77l-2.18-.87l-4.2 1.68L17.2 18l5.19-2a17.8 17.8 0 0 1-4.94-1.23Z"/><path fill="none" stroke="#4f4b45" stroke-linejoin="round" d="m17.2 18.03l19.89-7.96l-6.13-2.45l-19.89 7.96l6.13 2.45z"/><path fill="#fff5e3" stroke="#4f4b45" stroke-linejoin="round" d="M11.07 15.58v5.17l6.16 2.46l-.03-5.18l-6.13-2.45z"/><path fill="#f7e5c6" stroke="#4f4b45" stroke-linejoin="round" d="M11.07 34.34v3.81l6.16 2.47l-.03-3.82l-6.13-2.46z"/>`),
		g.Group(children),
	)
}
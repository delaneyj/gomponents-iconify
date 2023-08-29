package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Necktie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M5.25 44.92a18.75 2.08 0 1 0 37.5 0a18.75 2.08 0 1 0-37.5 0Z" opacity=".15"/><path fill="#80ddff" d="M5.25 9.28h37.5V44H5.25Z"/><path fill="#b8ecff" d="M37.19 9.28H10.81a5.55 5.55 0 0 0-5.56 5.55v4.46a5.56 5.56 0 0 1 5.56-5.56h26.38a5.56 5.56 0 0 1 5.56 5.56v-4.46a5.55 5.55 0 0 0-5.56-5.55Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M5.25 9.28h37.5V44H5.25Z"/><path fill="#fff48c" d="M26.39 17.61h-4.78a1.39 1.39 0 0 1-1.35-1.05l-1.82-7.28h11.12l-1.82 7.28a1.39 1.39 0 0 1-1.35 1.05Z"/><path fill="#ebcb00" d="m20.14 16.06l3.26-2.38a1 1 0 0 1 1.2 0l3.26 2.38l1.7-6.78H18.44Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M26.39 17.61h-4.78a1.39 1.39 0 0 1-1.35-1.05l-1.82-7.28h11.12l-1.82 7.28a1.39 1.39 0 0 1-1.35 1.05Z"/><path fill="#fff48c" d="M29.55 44H18.44l2.78-26.39h5.56L29.55 44z"/><path fill="#ebcb00" d="m27.05 20.21l-.27-2.6h-5.56l-.27 2.6h6.1z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M29.55 44H18.44l2.78-26.39h5.56L29.55 44z"/><path fill="#80ddff" d="M30.25 4.42L24 10.67l6.92 5a.7.7 0 0 0 1-.31l2.45-6.11Zm-12.5 0L24 10.67l-6.92 5a.7.7 0 0 1-1-.31l-2.5-6.08Z"/><path fill="#b8ecff" d="m32.08 6.55l-1.83-2.13L24 10.67l2.09 1.52l5.99-5.64zm-10.17 5.64L24 10.67l-6.25-6.25l-1.83 2.13l5.99 5.64z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M30.25 4.42L24 10.67l6.92 5a.7.7 0 0 0 1-.31l2.45-6.11Zm-12.5 0L24 10.67l-6.92 5a.7.7 0 0 1-1-.31l-2.5-6.08Z"/><path fill="#4acfff" d="M17.75 4.42L24 10.67l6.25-6.25h-12.5z"/><path fill="#00b8f0" d="M27.67 7h-7.34L24 10.67L27.67 7z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M17.75 4.42L24 10.67l6.25-6.25h-12.5z"/><path fill="#4acfff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="m39.2 29.8l-2.77 1.39a1.44 1.44 0 0 1-1.25 0l-2.77-1.39a1.39 1.39 0 0 1-.77-1.24v-4.7H40v4.7a1.39 1.39 0 0 1-.8 1.24Z"/><path fill="#b8ecff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M33 21.08h5.56A1.39 1.39 0 0 1 40 22.47v1.39h0h-8.36h0v-1.39A1.39 1.39 0 0 1 33 21.08Z"/>`),
		g.Group(children),
	)
}
package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoSmoking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#ff5a79"/><circle cx="32" cy="32" r="24" fill="#333"/><path fill="#eda454" d="M12 32h15.4v8H12z"/><path fill="#fff" d="M27.4 32h20.2v8H27.4z"/><g fill="#ffce31"><ellipse cx="13.3" cy="36" rx=".7" ry=".9"/><path d="M14 33.4c0 .5-.3.9-.7.9c-.4 0-.7-.4-.7-.9s.3-.9.7-.9c.4.1.7.5.7.9"/><ellipse cx="13.3" cy="38.6" rx=".7" ry=".9"/><ellipse cx="15.2" cy="37.1" rx=".7" ry=".9"/><ellipse cx="15.2" cy="34.9" rx=".7" ry=".9"/><ellipse cx="17.1" cy="36" rx=".7" ry=".9"/><path d="M17.8 33.4c0 .5-.3.9-.7.9c-.4 0-.7-.4-.7-.9s.3-.9.7-.9c.4.1.7.5.7.9"/><ellipse cx="17.1" cy="38.6" rx=".7" ry=".9"/><path d="M15.2 38.6c-.4 0-.7.4-.7.9h1.4c0-.6-.3-.9-.7-.9m0-5.2c.4 0 .7-.4.7-.9h-1.4c0 .5.3.9.7.9"/><ellipse cx="18.9" cy="37.1" rx=".7" ry=".9"/><ellipse cx="18.9" cy="34.9" rx=".7" ry=".9"/><path d="M18.9 38.6c-.4 0-.7.4-.7.9h1.4c0-.6-.3-.9-.7-.9m0-5.2c.4 0 .7-.4.7-.9h-1.4c0 .5.3.9.7.9"/><ellipse cx="20.5" cy="36" rx=".7" ry=".9"/><path d="M21.2 33.4c0 .5-.3.9-.7.9s-.7-.4-.7-.9s.3-.9.7-.9c.4.1.7.5.7.9"/><ellipse cx="20.5" cy="38.6" rx=".7" ry=".9"/><ellipse cx="22.5" cy="37.1" rx=".7" ry=".9"/><ellipse cx="22.5" cy="34.9" rx=".7" ry=".9"/><ellipse cx="24.4" cy="36" rx=".7" ry=".9"/><path d="M25.1 33.4c0 .5-.3.9-.7.9s-.7-.4-.7-.9s.3-.9.7-.9s.7.5.7.9"/><ellipse cx="24.4" cy="38.6" rx=".7" ry=".9"/><path d="M22.5 38.6c-.4 0-.7.4-.7.9h1.4c-.1-.6-.4-.9-.7-.9m0-5.2c.4 0 .7-.4.7-.9h-1.4c0 .5.3.9.7.9m4.3 3.7c0 .5-.3.9-.7.9c-.4 0-.7-.4-.7-.9s.3-.9.7-.9c.4 0 .7.4.7.9"/><ellipse cx="26.2" cy="34.9" rx=".7" ry=".9"/><path d="M26.2 38.6c-.4 0-.7.4-.7.9h1.4c-.1-.6-.4-.9-.7-.9m0-5.2c.4 0 .7-.4.7-.9h-1.4c0 .5.3.9.7.9"/></g><path fill="#ff8736" d="M47.6 32H50v8h-2.4z"/><path fill="#ff5a79" d="m9.23 13.474l4.243-4.242l41.294 41.294l-4.242 4.243z"/><path fill="#abc2c6" d="M30.6 11.9C27.5 23.7 37.1 7.5 49.5 31c-6.6-8.7-27.9-2.4-36.8-12C.2 5.6 35.9-7.6 30.6 11.9z"/>`),
		g.Group(children),
	)
}
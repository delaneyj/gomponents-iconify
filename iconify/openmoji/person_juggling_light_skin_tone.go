package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonJugglingLightSkinTone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<circle cx="19.969" cy="18.094" r="2" fill="#EA5A47"/><circle cx="40.969" cy="6.094" r="2" fill="#92D3F5"/><circle cx="55.969" cy="27.094" r="2" fill="#B1CC33"/><g fill="#fadcbc" stroke="#fadcbc"><circle cx="35.969" cy="13.094" r="3"/><path d="m30 26l-1.7 41h3.3L35 42h2l3.4 25h3.3L42 26s.214 4.116 1.417 6c.857 1.343 2.666 3.375 6.125 2.875c0 0 4.52-.563 5.333-2.25c.465-.966.125-1.584.042-1.875s-2.667.709-2.667.709L49.167 32l-2.083-1.125l-1.334-1.458l-1.958-7.292l-3.458-1.333l-9.709-.208l-3.416 4.458l-.542 4.291L25 30.5l-6.666-6.167s-1.459 1.73-.459 3.105s3.788 5.324 5.916 5.728c2.417.459 3.793-.615 4.375-1.416c.667-.917 1.501-3.083 1.834-5.75z"/></g><g fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2"><circle cx="35.969" cy="13.094" r="3"/><path stroke-linecap="round" stroke-linejoin="round" d="m41 26l2.568 39.004c.073 1.098-.61 1.996-1.517 1.996c-.908 0-1.772-.892-1.92-1.981L37.27 43.98C37.12 42.891 36.55 42 36 42s-1.121.892-1.27 1.981L31.87 65.02c-.15 1.089-1.014 1.98-1.92 1.98c-.908 0-1.59-.898-1.518-1.996L31 26"/><path stroke-linecap="round" stroke-linejoin="round" d="m19 25l3.707 4.237c1.812 2.07 3.663 1.543 4.115-1.17l.356-2.135C27.63 23.22 29.8 20.775 32 20.5c2.2-.275 5.8-.275 8 0s4.37 2.72 4.822 5.432l.356 2.136C45.63 30.78 47.8 32.55 50 32l4-1"/></g>`),
		g.Group(children),
	)
}
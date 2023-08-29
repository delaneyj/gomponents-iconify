package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Footprints(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#597B91" d="M54.4 34.1c-5.4-1.4-6-3.1-10-4.2c-2.4-.6-4.9.8-5.5 3.2c-1.6 6.1 8.7 9.1 6.1 18.8c-1.3 5-4.1 5.8-5.6 11.7c-.9 3.3 1.1 6.6 4.3 7.5c2.2.6 4.4-.1 5.9-1.6c.7-.5 1.2-1.4 1.6-2.8c1.1-4 .6-7.1 1.9-11.8c1.5-5.8 5-10.5 5.7-13.3c1.4-5.1-2.5-7.1-4.4-7.5z"/><ellipse cx="43" cy="23.8" fill="#597B91" rx="2.6" ry="3.9" transform="rotate(-165.001 43.03 23.763)"/><ellipse cx="48.1" cy="25.8" fill="#597B91" rx="2.1" ry="3.1" transform="rotate(-165.001 48.114 25.776)"/><ellipse cx="52.4" cy="27.8" fill="#597B91" rx="1.8" ry="2.7" transform="rotate(-165.001 52.366 27.805)"/><ellipse cx="55.7" cy="29.8" fill="#597B91" rx="1.4" ry="2.2" transform="rotate(-165.001 55.7 29.786)"/><ellipse cx="58.5" cy="31.6" fill="#597B91" rx="1.3" ry="2" transform="rotate(-159.826 58.462 31.56)"/><path fill="#597B91" d="M27 33c-2.6-9.7 7.7-12.7 6.1-18.8c-.6-2.4-3.1-3.9-5.5-3.2c-3.9 1.1-4.6 2.7-10 4.2c-1.9.5-5.7 2.5-4.4 7.6c.8 2.8 4.2 7.6 5.7 13.3c1.3 4.7.8 7.8 1.9 11.8c.4 1.4.9 2.3 1.6 2.8c1.5 1.5 3.7 2.2 5.9 1.6c3.3-.9 5.2-4.2 4.3-7.5c-1.5-6-4.2-6.8-5.6-11.8z"/><ellipse cx="29" cy="4.8" fill="#597B91" rx="2.6" ry="3.9" transform="rotate(-14.999 28.973 4.842)"/><ellipse cx="23.9" cy="6.9" fill="#597B91" rx="2.1" ry="3.1" transform="rotate(-14.999 23.888 6.855)"/><ellipse cx="19.6" cy="8.9" fill="#597B91" rx="1.8" ry="2.7" transform="rotate(-14.999 19.636 8.884)"/><ellipse cx="16.3" cy="10.9" fill="#597B91" rx="1.4" ry="2.2" transform="rotate(-14.999 16.301 10.866)"/><ellipse cx="13.5" cy="12.6" fill="#597B91" rx="1.3" ry="2" transform="rotate(-20.174 13.532 12.642)"/>`),
		g.Group(children),
	)
}
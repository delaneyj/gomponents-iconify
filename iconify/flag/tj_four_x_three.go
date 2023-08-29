package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TjFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#060" d="M0 0h640v480H0z"/><path fill="#fff" d="M0 0h640v342.9H0z"/><path fill="#c00" d="M0 0h640v137.1H0z"/><path fill="#f8c300" d="M300.8 233.6a8.6 8.6 0 0 1 16 4V272h6.4v-34.3a8.6 8.6 0 0 1 16-4a20.2 20.2 0 1 0-38.4 0"/><path fill="#fff" d="M305.4 224.7a13.7 13.7 0 0 1 14.6 6.5a13.7 13.7 0 0 1 14.6-6.5a14.7 14.7 0 0 0-29.2 0"/><path id="flagTj4x30" fill="#f8c300" d="M316.8 258.3a26 26 0 0 1-43.8 16.6a27 27 0 0 1-41 12c2.5 25 40 19.9 42.8-4.4c11.7 20.7 37.6 14.7 45.2-10.6z"/><use width="100%" height="100%" fill="#f8c300" href="#flagTj4x30" transform="matrix(-1 0 0 1 640 0)"/><path id="flagTj4x31" fill="#f8c300" d="M291.8 302.6c-5.3 11.3-15.7 13.2-24.8 4.1c0 0 3.6-2.6 7.6-3.3c-.8-3.1.7-7.5 2.9-9.8a15 15 0 0 1 6.1 8.1c5.5-.7 8.2 1 8.2 1z"/><use width="100%" height="100%" fill="#f8c300" href="#flagTj4x31" transform="rotate(9.4 320 551.3)"/><use width="100%" height="100%" fill="#f8c300" href="#flagTj4x31" transform="rotate(18.7 320 551.3)"/><path fill="none" stroke="#f8c300" stroke-width="11" d="M253.5 327.8a233.1 233.1 0 0 1 133 0"/><g fill="#f8c300" transform="translate(320 164.6) scale(.68571)"><path id="flagTj4x32" d="m301930 415571l-790463-574305h977066l-790463 574305L0-513674z" transform="scale(.00005)"/></g><g id="flagTj4x33" fill="#f8c300" transform="translate(320 260.6) scale(.68571)"><use width="100%" height="100%" href="#flagTj4x32" transform="translate(-70 -121.2)"/><use width="100%" height="100%" href="#flagTj4x32" transform="translate(-121.2 -70)"/><use width="100%" height="100%" href="#flagTj4x32" transform="translate(-140)"/></g><use width="100%" height="100%" fill="#f8c300" href="#flagTj4x33" transform="matrix(-1 0 0 1 640 0)"/>`),
		g.Group(children),
	)
}
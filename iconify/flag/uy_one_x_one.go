package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UyOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#fff" d="M0 0h512v512H0z"/><path fill="#0038a8" d="M284 56.9h228v56.9H284zm0 113.8h228v56.9H284zM0 284.4h512v57H0zm0 113.8h512v57H0z"/><g fill="#fcd116" stroke="#000" stroke-miterlimit="20" stroke-width=".6" transform="translate(142.2 142.2) scale(3.12889)"><g id="flagUy1x10"><g id="flagUy1x11"><g id="flagUy1x12"><path stroke-linecap="square" d="m-2 8.9l3 4.5c-12.4 9-4.9 14.2-13.6 17c5.4-5.2-.9-5.7 3.7-16.8"/><path fill="none" d="M-4.2 10.2c-6.8 11.2-2.4 17.4-8.4 20.3"/><path d="M0 0h6L0 33L-6 0h6v33"/></g><use width="100%" height="100%" href="#flagUy1x12" transform="rotate(45)"/></g><use width="100%" height="100%" href="#flagUy1x11" transform="rotate(90)"/></g><use width="100%" height="100%" href="#flagUy1x10" transform="scale(-1)"/><circle r="11"/></g><g transform="translate(142.2 142.2) scale(.31289)"><g id="flagUy1x13"><path d="M81-44c-7 8-11-6-36-6S16-35 12-38s21-21 29-22s31 7 40 16m-29 9c7 6 1 19-6 19S26-28 32-36"/><path d="M19-26c1-12 11-14 27-14s23 12 29 15c-7 0-13-10-29-10s-16 0-27 10m3 2c4-6 9 6 20 6s17-3 24-8s-10 12-21 12s-26-6-23-10"/><path d="M56-17c13-7 5-17 0-19c2 2 10 12 0 19M0 43c6 0 8-2 16-2s27 11 38 7c-23 9-14 3-54 3h-5m63 6c-4-7-3-5-11-16c8 6 10 9 11 16M0 67c25 0 21-5 54-19c-24 3-29 11-54 11h-5m5-29c7 0 9-5 17-5s19 3 24 7c1 1-3-8-11-9S25 9 16 7c0 4 3 3 4 9c0 5-9 5-11 0c2 8-4 8-9 8"/></g><use width="100%" height="100%" href="#flagUy1x13" transform="scale(-1 1)"/><path d="M0 76c-5 0-18 3 0 3s5-3 0-3"/></g>`),
		g.Group(children),
	)
}
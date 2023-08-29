package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Knotwords(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m22.812 6.292l9.053-.792l.792 9.053l-9.053.792z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.604 15.344l9.053-.792l.792 9.053l-9.053.792z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m32.656 14.552l9.053-.792l.792 9.053l-9.053.792zm-18.104 1.584l9.053-.792l.792 9.053l-9.053.792z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.5 16.928l9.053-.792l.792 9.053l-9.053.792zm18.896 7.468l9.053-.792l.792 9.053l-9.053.792z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m15.344 25.188l9.053-.792l.792 9.053l-9.053.792z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m25.188 33.448l9.053-.792l.792 9.053l-9.053.792zm2.554-2.023l-.479-5.48l1.794-.157c1.014-.089 1.909.663 1.998 1.68s-.662 1.912-1.676 2l-1.794.158m1.794-.157l1.951 1.641m4.193-15.009l3.631-.318m-1.336 5.639l-.48-5.48m-19.44 7.223l-.479-5.48l4.11 5.162l-.479-5.48m9.184-9.925l-.89 5.6l-1.85-5.36l-.89 5.6l-1.85-5.36"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28.534 40.477l-.48-5.48l1.234-.108a2.407 2.407 0 0 1 2.607 2.188l.06.685a2.407 2.407 0 0 1-2.188 2.607l-1.233.108Z"/><rect width="3.644" height="5.501" x="26.704" y="16.724" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.822" ry="1.822" transform="rotate(-5 28.526 19.474)"/><rect width="3.644" height="5.501" x="18.444" y="26.567" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.822" ry="1.822" transform="rotate(-5 20.266 29.318)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m8.71 18.447l.479 5.48m-.167-1.909l2.635-3.811m.477 5.462l-2.496-2.543"/>`),
		g.Group(children),
	)
}
package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CityscapeAtDusk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#fbbf67" d="M0 0h64v64H0z"/><path fill="#f15744" d="M0 0h64v64H0z" opacity=".7"/><path fill="#62727a" d="M48 45.7h16V64H44V49.7c0-2.2 1.8-4 4-4"/><path fill="#d0d0d0" d="M60.8 49h.4c.4 0 .8.4.8 1v14h-2V50c0-.6.4-1 .8-1m-5 0h.4c.4 0 .8.4.8 1v14h-2V50c0-.6.4-1 .8-1"/><g fill="#62727a"><path d="M51.3 43h11v2.9h-11zM0 43.1h13.1v-5.2L0 40.9z"/><path d="M0 43h13v21H0z"/></g><path fill="#ed4c5c" d="M9.5 51c0 .5-.4 1-1 1h-3c-.6 0-1-.5-1-1v-2c0-.6.4-1 1-1h3c.6 0 1 .4 1 1v2m0 8c0 .5-.4 1-1 1h-3c-.6 0-1-.5-1-1v-2c0-.6.4-1 1-1h3c.6 0 1 .4 1 1v2"/><path fill="#3e4347" d="M50 38v-4l-6-2V21h-2v11l-6 2v4h-3v-7.7L23 19L13 30.3V64h40V38z"/><path fill="#6adbc6" d="M41.5 45c0 .5-.5 1-1 1h-3c-.5 0-1-.5-1-1v-2c0-.6.5-1 1-1h3c.5 0 1 .4 1 1v2m8 8c0 .5-.5 1-1 1h-3c-.5 0-1-.5-1-1v-2c0-.6.5-1 1-1h3c.5 0 1 .4 1 1v2m-8 8c0 .5-.5 1-1 1h-3c-.5 0-1-.5-1-1v-2c0-.6.5-1 1-1h3c.5 0 1 .4 1 1v2m8 0c0 .5-.5 1-1 1h-3c-.5 0-1-.5-1-1v-2c0-.6.5-1 1-1h3c.5 0 1 .4 1 1v2"/><path fill="#ffdd7d" d="M30 43c0 .5-.5 1-1 1H17c-.5 0-1-.5-1-1v-2c0-.6.5-1 1-1h12c.5 0 1 .4 1 1v2m0 16c0 .5-.5 1-1 1H17c-.5 0-1-.5-1-1v-2c0-.6.5-1 1-1h12c.5 0 1 .4 1 1v2m0-24c0 .5-.5 1-1 1H17c-.5 0-1-.5-1-1v-2c0-.6.5-1 1-1h12c.5 0 1 .4 1 1v2m0 16c0 .5-.5 1-1 1H17c-.5 0-1-.5-1-1v-2c0-.6.5-1 1-1h12c.5 0 1 .4 1 1v2"/>`),
		g.Group(children),
	)
}
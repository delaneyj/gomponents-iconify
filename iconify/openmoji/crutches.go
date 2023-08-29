package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crutches(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d0cfce" d="m35.665 44.073l.371-.939l4.365-1.5l3.9 5.521l-4.308 3.044l-4.328-6.126zm-8.425 2.345l4.492-6.335l4.303 3.05l-4.492 6.336z"/><path fill="#92d3f5" d="m53.825 25.64l-10.752-7.522l1.455-2.08c.701-1.003 1.796-1.406 2.36-.868l2.855 2.724a8.09 8.09 0 0 0 2.006 1.403l3.537 1.749c.698.345.695 1.512-.006 2.515Zm-24.754-7.688l-10.914 7.483l-1.447-2.111c-.698-1.018-.69-2.195.017-2.536l3.584-1.73a8.161 8.161 0 0 0 2.036-1.396l2.905-2.719c.573-.537 1.674-.12 2.372.898Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m35.665 44.073l.371-.939l4.365-1.5l3.9 5.521l-4.308 3.044l-4.328-6.126zm-8.425 2.345l4.492-6.335l4.303 3.05l-4.492 6.336zM53.825 25.64l-10.752-7.522l1.455-2.08c.701-1.003 1.796-1.406 2.36-.868l2.855 2.724a8.09 8.09 0 0 0 2.006 1.403l3.537 1.749c.698.345.695 1.512-.006 2.515Zm-24.754-7.688l-10.914 7.483l-1.447-2.111c-.698-1.018-.69-2.195.017-2.536l3.584-1.73a8.161 8.161 0 0 0 2.036-1.396l2.905-2.719c.573-.537 1.674-.12 2.372.898Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M43.045 18.892L32.885 32.77a3.961 3.961 0 0 0-.787 2.412l-.365 4.901l-4.492 6.335l4.303 3.052l4.492-6.335l4.803-1.66a3.962 3.962 0 0 0 2.06-1.48l10.109-13.808m-17.119 2.481l-6.609-9.611"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m35.494 43.898l4.376 6.161l4.301-3.054l-3.77-5.371M19.106 26.053l9.745 14.172a4.034 4.034 0 0 0 1.673 1.41"/><rect width="6.791" height="2.291" x="24.127" y="49.423" rx=".8" transform="rotate(-54.232 27.523 50.569)"/><rect width="3.417" height="3.417" x="22.398" y="53.603" rx="1.175" transform="rotate(35.768 24.106 55.311)"/><rect width="6.791" height="2.291" x="40.603" y="50.007" rx=".8" transform="rotate(-123.03 43.998 51.152)"/><rect width="3.417" height="3.417" x="45.476" y="54.344" rx="1.175" transform="rotate(-33.03 47.184 56.052)"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m39.515 24.604l9.551 7.054M21.51 29.55l10.128-7.063"/>`),
		g.Group(children),
	)
}
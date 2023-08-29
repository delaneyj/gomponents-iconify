package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonWearingTurbanOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M11.5 44.5a12.5 1.5 0 1 0 25 0a12.5 1.5 0 1 0-25 0Z" opacity=".15"/><path fill="#ffe500" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24 5A14.59 14.59 0 0 0 9.41 19.6v7.64h29.18V19.6A14.59 14.59 0 0 0 24 5Z"/><path fill="#ffe500" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M40.15 26.13a2.73 2.73 0 0 0-2.07-2.3l-.87-.24A3 3 0 0 1 35 20.66v-3.07a2.77 2.77 0 0 0-2.08-2.68A23 23 0 0 1 24 16.52a23 23 0 0 1-8.9-1.61a2.77 2.77 0 0 0-2.1 2.68v3.07a3 3 0 0 1-2.23 2.93l-.87.24a2.73 2.73 0 0 0-2.07 2.3a2.69 2.69 0 0 0 2.68 2.92h.17a13.45 13.45 0 0 0 26.6 0h.17a2.69 2.69 0 0 0 2.7-2.92Z"/><path fill="#45413c" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M15.93 25.81a1.15 1.15 0 1 0 1.14-1.15a1.15 1.15 0 0 0-1.14 1.15Zm16.14 0a1.15 1.15 0 1 1-1.14-1.15a1.15 1.15 0 0 1 1.14 1.15Z"/><path fill="#ffaa54" d="M13.69 30.69a1.49.89 0 1 0 2.98 0a1.49.89 0 1 0-2.98 0Zm17.64 0a1.49.89 0 1 0 2.98 0a1.49.89 0 1 0-2.98 0Z"/><path fill="#f0f0f0" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M35.35 21.42v6.94a11.39 11.39 0 0 1-5 9.46v-1.27c0-2.59-2.83-4.68-6.31-4.68s-6.35 2.13-6.35 4.68v1.27a11.39 11.39 0 0 1-5-9.46v-6.94l-1.92 2.17v4.94c0 9.27 8.11 14.9 13.27 14.86s13.27-5.57 13.27-14.86v-4.94Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M21.19 36.12s.69.93 2.81.93s2.81-.93 2.81-.93"/><path fill="#f0f0f0" d="M39.38 16.66v-2.84C39.42 6.9 32.82 1.65 24 1.65S8.58 6.9 8.58 13.82c0 8.92.83 10.21.83 10.21a88 88 0 0 1 15-5.39A58.43 58.43 0 0 1 38.59 24s.64-1 .79-7.34Z"/><path fill="#fff" d="M24 5.74c8.42 0 14.81 4.78 15.38 11.24v-3.16C39.42 6.9 32.82 1.65 24 1.65S8.58 6.9 8.58 13.82v3.16C9.18 10.53 15.57 5.74 24 5.74Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M39.38 16.66v-2.84C39.42 6.9 32.82 1.65 24 1.65S8.58 6.9 8.58 13.82c0 8.92.83 10.21.83 10.21a88 88 0 0 1 15-5.39A58.43 58.43 0 0 1 38.59 24s.64-1 .79-7.34Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M8.58 13.82s5.77-10.52 22.35-11m4.45 9.11a43.66 43.66 0 0 0-22.13 6.62m11.15.09c10.5-2.51 15-2 15-2"/>`),
		g.Group(children),
	)
}
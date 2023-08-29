package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartWithArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M15.5 45.5a9.5 1.5 0 1 0 19 0a9.5 1.5 0 1 0-19 0Z" opacity=".15"/><path fill="#bf8256" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M34.497 13.117L41.6 6.192l2.311 2.37l-7.102 6.925z"/><path fill="#00b8f0" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M41.15 3.93a.83.83 0 0 0-1.4-.53l-2.52 2.45a3.66 3.66 0 0 0-.36 5l4.74-4.65a11.44 11.44 0 0 1-.46-2.27Zm5.02 5.15a.83.83 0 0 1 .5 1.42L44.15 13a3.63 3.63 0 0 1-5 .23l4.73-4.61a12 12 0 0 0 2.29.46Z"/><path fill="#ff6242" d="M21.19 18a8.25 8.25 0 0 0-9.8-6.36C5.74 12.81 4.51 19 5.45 23.43C8 35.19 20.74 40.39 25 41.81a2.45 2.45 0 0 0 2.44-.52c3.29-3 12.87-13 10.37-24.73c-.96-4.46-4.59-9.56-10.25-8.39A8.28 8.28 0 0 0 21.19 18Z"/><path fill="#ff866e" d="M28.83 12.13a8.68 8.68 0 0 0-6.69 6.69a.82.82 0 0 1-1.46.31a8.69 8.69 0 0 0-8.83-3.4c-4 .86-5.93 4.13-6.42 7.6c-.91-4.45.34-10.53 6-11.72A8.25 8.25 0 0 1 21.19 18a8.28 8.28 0 0 1 6.37-9.8c5.62-1.2 9.23 3.8 10.21 8.27c-1.85-2.98-4.9-5.2-8.94-4.34Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M21.19 18a8.25 8.25 0 0 0-9.8-6.36C5.74 12.81 4.51 19 5.45 23.43C8 35.19 20.74 40.39 25 41.81a2.45 2.45 0 0 0 2.44-.52c3.29-3 12.87-13 10.37-24.73c-.96-4.46-4.59-9.56-10.25-8.39A8.28 8.28 0 0 0 21.19 18Z"/><path fill="#c0dceb" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="m8.37 41.84l4.34-1.39a1.23 1.23 0 0 0 .51-2l-1.82-1.9a1.23 1.23 0 0 0-2.05.45l-1.5 4.31a.42.42 0 0 0 .52.53Z"/><path fill="#bf8256" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M21.7 28.15a4.62 4.62 0 0 1-2-.22l-8.57 8.36l2.31 2.37l8.77-8.56a3.93 3.93 0 0 1-.51-1.95Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M18.39 26.57c0 1 1.55 1.74 3.31 1.58c0 2.12 2.24 4.48 3.51 3.73M6.45 10.37a6.62 6.62 0 0 0-5.14 7.82M7.22 7.51a3.72 3.72 0 0 0-4.84 2.07"/>`),
		g.Group(children),
	)
}
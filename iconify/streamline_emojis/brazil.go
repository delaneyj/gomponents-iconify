package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Brazil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#46b000" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M45 12.25h-6.32A45.89 45.89 0 0 1 24 9.88A45.73 45.73 0 0 0 9.37 7.5H3c-.58 0-1 .35-1 .79v26a.94.94 0 0 0 1 .79h6.37A46 46 0 0 1 24 37.46a45.62 45.62 0 0 0 14.65 2.38H45a.93.93 0 0 0 1-.79V13a.94.94 0 0 0-1-.75Z"/><path fill="#ffe500" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24 34.47s-3.55-2.77-7.59-6.11c-3.7-3.06-8-6.33-8-6.33S12 19.53 16.34 17S24 12.87 24 12.87s3.42 2.68 7.58 6c4.44 3.56 8 6.42 8 6.42s-3.8 2.71-8.05 5.19S24 34.47 24 34.47Z"/><path fill="#009fd9" d="M18.319 25.225a6.99 5.92 74.78 1 0 11.425-3.108a6.99 5.92 74.78 1 0-11.425 3.108Z"/><path fill="#fff" d="M18.22 21.38a7.86 7.86 0 0 0-.22 1.71v.59a21.14 21.14 0 0 1 7.57.77a15.62 15.62 0 0 1 4.13 1.85a7.58 7.58 0 0 0 .3-2.05s0 0 0-.05a14.82 14.82 0 0 0-4.4-2a21.33 21.33 0 0 0-7.38-.82Zm2.52 4.78a.58.54 0 1 0 1.16 0a.58.54 0 1 0-1.16 0Zm3.15 2.61a.44.41 0 1 0 .88 0a.44.41 0 1 0-.88 0Zm-.44-1.66a.44.41 0 1 0 .88 0a.44.41 0 1 0-.88 0Zm3.75-5.78a.44.41 0 1 0 .88 0a.44.41 0 1 0-.88 0Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M18.319 25.225a6.99 5.92 74.78 1 0 11.425-3.108a6.99 5.92 74.78 1 0-11.425 3.108Z"/><path fill="#45413c" d="M12.5 45.5a11.5 1.5 0 1 0 23 0a11.5 1.5 0 1 0-23 0Z" opacity=".15"/>`),
		g.Group(children),
	)
}
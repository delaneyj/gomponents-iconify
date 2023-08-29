package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SouthAfrica(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#46b000" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M45 11.75h-6.32A45.89 45.89 0 0 1 24 9.38A45.73 45.73 0 0 0 9.37 7H3c-.58 0-1 .35-1 .79v26a.94.94 0 0 0 1 .79h6.37A46 46 0 0 1 24 37a45.62 45.62 0 0 0 14.65 2.38H45a.93.93 0 0 0 1-.79v-26a.94.94 0 0 0-1-.84Z"/><path fill="#fff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24 27.65a42.82 42.82 0 0 0-5.56-1.46a67.75 67.75 0 0 0-8.33 8.41A45.94 45.94 0 0 1 24 37a45.62 45.62 0 0 0 14.65 2.38H45a.93.93 0 0 0 1-.79V30h-7.32A45.62 45.62 0 0 1 24 27.65Z"/><path fill="#ffe500" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M13.44 20.14c-2.91-4.33-11.44-9-11.44-9v20.52a66.72 66.72 0 0 1 11.44-11.52Z"/><path fill="#fff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M45 11.75h-6.32A45.89 45.89 0 0 1 24 9.38A45.44 45.44 0 0 0 10.13 7c3.8 2.21 7.54 6.34 10.06 10.41c1.3.31 2.59.66 3.84 1.09a45.89 45.89 0 0 0 14.65 2.37h7.37v-8.33a.94.94 0 0 0-1.05-.79Z"/><path fill="#fff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M12.28 7.1a47.123 47.123 0 0 1 0 0Z"/><path fill="#656769" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M9.23 19.85A33.69 33.69 0 0 0 2 14.36V27a86.33 86.33 0 0 1 7.23-7.15Z"/><path fill="#fff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M9.37 7c.82 0 1.64 0 2.46.07C11 7 10.19 7 9.37 7Z"/><path fill="#e04122" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M45 11.75h-6.32A45.89 45.89 0 0 1 24 9.38a44.31 44.31 0 0 0-9.35-2.07a32.12 32.12 0 0 1 7.06 7.94c.77.21 1.53.44 2.29.7a45.89 45.89 0 0 0 14.65 2.37h7.37v-5.78a.94.94 0 0 0-1.02-.79Z"/><path fill="#fff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M12.05 34.67h0Z"/><path fill="#009fd9" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24 30.52c-1.37-.46-2.78-.84-4.21-1.17a47.41 47.41 0 0 0-5.48 5.51A45.45 45.45 0 0 1 24 37a45.62 45.62 0 0 0 14.65 2.38H45a.93.93 0 0 0 1-.79V32.9h-7.32A45.62 45.62 0 0 1 24 30.52Z"/><path fill="#45413c" d="M12.5 45.5a11.5 1.5 0 1 0 23 0a11.5 1.5 0 1 0-23 0Z" opacity=".15"/>`),
		g.Group(children),
	)
}
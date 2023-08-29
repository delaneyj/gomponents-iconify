package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedPaperLantern(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<defs><path id="streamlineEmojisRedPaperLantern0" fill="#ff6242" d="M10.55 22.6a13.45 16.69 0 1 0 26.9 0a13.45 16.69 0 1 0-26.9 0Z"/></defs><path fill="#656769" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24 34.49c-3.34 0-6 1.12-6 2.51v2.49c0 1.37 2.66 2.51 6 2.51s6.05-1.11 6.05-2.49V37c0-1.39-2.71-2.51-6.05-2.51Z"/><path fill="#45413c" d="M10.74 45.27a13.26 1.73 0 1 0 26.52 0a13.26 1.73 0 1 0-26.52 0Z" opacity=".15"/><use href="#streamlineEmojisRedPaperLantern0"/><use href="#streamlineEmojisRedPaperLantern0"/><use href="#streamlineEmojisRedPaperLantern0"/><path fill="#ff866e" d="M24 17.16c6.36 0 11.68.85 13.08 10.84a36 36 0 0 0 .37-5.4c0-15.53-6-16.68-13.45-16.68S10.55 7.07 10.55 22.6a36 36 0 0 0 .37 5.4c1.4-10 6.72-10.84 13.08-10.84Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M10.55 22.6a13.45 16.69 0 1 0 26.9 0a13.45 16.69 0 1 0-26.9 0Z"/><path fill="#87898c" d="M24 4c-3.34 0-6 1.12-6 2.49V9c0 1.37 2.71 2.49 6.05 2.49s6-1.13 6-2.49V6.49C30.05 5.12 27.34 4 24 4Z"/><path fill="#656769" d="M24 9.65c-3.34 0-6-1.11-6-2.49V9c0 1.37 2.71 2.49 6.05 2.49s6-1.13 6-2.49V7.16c0 1.38-2.71 2.49-6.05 2.49Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24 4c-3.34 0-6 1.12-6 2.49V9c0 1.37 2.71 2.49 6.05 2.49s6-1.13 6-2.49V6.49C30.05 5.12 27.34 4 24 4Z"/><path fill="#ffe500" d="M19.06 5.92a4.94 1.82 0 1 0 9.88 0a4.94 1.82 0 1 0-9.88 0Z"/><path fill="#ffaa54" d="M19.75 6.83a4.25.91 0 1 0 8.5 0a4.25.91 0 1 0-8.5 0Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M19.06 5.92a4.94 1.82 0 1 0 9.88 0a4.94 1.82 0 1 0-9.88 0Zm-7.19 6.71s1.66 4.52 12.13 4.52s12.13-4.52 12.13-4.52m-25.57 8.82c.17.37 2.34 4.49 13.44 4.49s13.27-4.13 13.44-4.49M11.29 30s1.76 4.49 12.71 4.49S36.71 30 36.71 30"/>`),
		g.Group(children),
	)
}
package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonGettingHaircutMediumSkinTone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#92d3f5" d="M7.814 61.021v-3.958c0-4.994 5.008-9 10-9q9 7.5 18 0c4.994 0 10 4.006 10 9v3.958"/><path fill="#9b9b9a" d="M57.708 40.014a5 5 0 1 1-8.145-.202l-.499-.749l.688-.562l-2.75-3.063l-1.5-3.625l-.195-4.328l2.574 1.203l9.683 9.875s1.063.936 1 1.25"/><path fill="#d0cfce" d="M53.564 31.813a5.035 5.035 0 0 0 .1 1l-15.85.25s.75 3.75 3.75 3.75h17a5 5 0 1 0-5-5Z"/><path fill="#c19a65" d="M15.75 31.063c0 9 4.937 14 11 14c5.937 0 11.064-5 11.064-14a12.133 12.133 0 0 0-1-5c-3-3-7-8-7-8c-4 3-7 6-13 7c0 0-1.063 1-1.063 6Z"/><path fill="#6a462f" d="M16.814 39.063c-4 0-4-6-4-13s4-14 14-14s13.122 7.018 17 14c.706 1.272 1 3 1 6l-7 1c.946 2.77 3.196 4.357 6.942 4.586c0 0-.817 3.414-1.63 3.414s-8.122-.25-8.122-.25s4.826-11.483 1.054-14.777l-6.322-7.49l-12.908 7.225s-2.258 9.639-.014 13.292Z"/><circle cx="58.564" cy="31.813" r="2"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M44.814 37.063c-.042 1.108-.692 4-2 4h-7m-19-2c-4 0-4-6-4-13s4-14 14-14s13.122 7.018 17 14c.706 1.272 1 3 1 6m-37 28v-3c0-4.994 5.008-9 10-9q9 7.5 18 0c4.994 0 10 4.006 10 9v3"/><path d="M32.687 30.063a2 2 0 1 1-2-2a2 2 0 0 1 2 2m-8 0a2 2 0 1 1-2-2a2 2 0 0 1 2 2"/><path fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2" d="M15.75 31.063c0 9 4.937 14 11 14c5.937 0 11.064-5 11.064-14a12.133 12.133 0 0 0-1-5c-3-3-7-8-7-8c-4 3-7 6-13 7c0 0-1.063 1-1.063 6Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M23.814 38.063a6.553 6.553 0 0 0 6 0m23.75-6.25a5.035 5.035 0 0 0 .1 1l-15.85.25s.75 3.75 3.75 3.75h17a5 5 0 1 0-5-5Zm4.144 8.201a5 5 0 1 1-8.145-.202"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M44.293 26.803a3.064 3.064 0 0 1 .749.444a3.218 3.218 0 0 1 .266.238l2.573 2.574"/><circle cx="53.564" cy="42.813" r="2"/>`),
		g.Group(children),
	)
}
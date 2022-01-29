/*
0 - Без стиля.По умолчанию.
1 - Теории заговора.
2 - ТВ-репортажи.
3 - Тосты.
4 - Пацанские цитаты.
5 - Рекламные слоганы.
6 - Короткие истории.
7 - Подписи в Instagram.
8 - Короче, Википедия.
9 - Синопсисы фильмов.
10 - Гороскоп.
11 - Народные мудрости.
18 - Новый Европейский Театр.
19 - Яндекс.Директ.
20 - Новогодние открытки.
*/
package balabola

type TextStyle int

const (
	TextStyleWithoutStyle TextStyle = iota
	TextStyleConspiracyTheories
	TextStyleTVReports
	TextStyleToasts
	TextStyleBoyQuotes
	TextStyleAdvertisingSlogans
	TextStyleShortStories
	TextStyleInstagramCaptions
	TextStyleInShortWikipedia
	TextStyleMovieSynopses
	TextStyleHoroscope
	TextStyleFolkWisdom
	TextStyleTheNewEuropeanTheater = 18
	TextStyleYandexDirect          = 19
	TextStyleNewYearsCards         = 20
)

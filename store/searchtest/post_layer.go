// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package searchtest

import (
	"testing"
	"time"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/store"
	"github.com/stretchr/testify/require"
)

var searchPostStoreTests = []searchTest{
	{
		Name: "Should be able to search posts including results from DMs",
		Fn:   testSearchPostsIncludingDMs,
		Tags: []string{ENGINE_ALL},
	},
	{
		Name: "Should be able to search posts using pagination",
		Fn:   testSearchPostsWithPagination,
		Tags: []string{ENGINE_ELASTICSEARCH, ENGINE_BLEVE},
	},
	{
		Name: "Should return pinned and unpinned posts",
		Fn:   testSearchReturnPinnedAndUnpinned,
		Tags: []string{ENGINE_ALL},
	},
	{
		Name: "Should be able to search for exact phrases in quotes",
		Fn:   testSearchExactPhraseInQuotes,
		Tags: []string{ENGINE_POSTGRES, ENGINE_MYSQL, ENGINE_ELASTICSEARCH},
	},
	{
		Name: "Should be able to search for email addresses with or without quotes",
		Fn:   testSearchEmailAddresses,
		Tags: []string{ENGINE_ELASTICSEARCH},
	},
	{
		Name: "Should be able to search when markdown underscores are applied",
		Fn:   testSearchMarkdownUnderscores,
		Tags: []string{ENGINE_POSTGRES, ENGINE_ELASTICSEARCH},
	},
	{
		Name: "Should be able to search for non-latin words",
		Fn:   testSearchNonLatinWords,
		Tags: []string{ENGINE_ELASTICSEARCH},
	},
	{
		Name: "Should be able to search for alternative spellings of words",
		Fn:   testSearchAlternativeSpellings,
		Tags: []string{ENGINE_ELASTICSEARCH},
	},
	{
		Name: "Should be able to search for alternative spellings of words with and without accents",
		Fn:   testSearchAlternativeSpellingsAccents,
		Tags: []string{ENGINE_ELASTICSEARCH},
	},
	{
		Name: "Should be able to search or exclude messages written by a specific user",
		Fn:   testSearchOrExcludePostsBySpecificUser,
		Tags: []string{ENGINE_ALL},
	},
	{
		Name: "Should be able to search or exclude messages written in a specific channel",
		Fn:   testSearchOrExcludePostsInChannel,
		Tags: []string{ENGINE_ALL},
	},
	{
		Name: "Should be able to search or exclude messages written in a DM or GM",
		Fn:   testSearchOrExcludePostsInDMGM,
		Tags: []string{ENGINE_ALL},
	},
	{
		Name: "Should be able to filter messages written after a specific date",
		Fn:   testFilterMessagesAfterSpecificDate,
		Tags: []string{ENGINE_ALL},
	},
	{
		Name: "Should be able to filter messages written before a specific date",
		Fn:   testFilterMessagesBeforeSpecificDate,
		Tags: []string{ENGINE_ALL},
	},
	{
		Name: "Should be able to filter messages written on a specific date",
		Fn:   testFilterMessagesInSpecificDate,
		Tags: []string{ENGINE_ALL},
	},
	{
		Name: "Should be able to exclude messages that contain a serch term",
		Fn:   testFilterMessagesWithATerm,
		Tags: []string{ENGINE_MYSQL, ENGINE_POSTGRES},
	},
	{
		Name: "Should be able to search using boolean operators",
		Fn:   testSearchUsingBooleanOperators,
		Tags: []string{ENGINE_ELASTICSEARCH},
	},
	{
		Name: "Should be able to search with combined filters",
		Fn:   testSearchUsingCombinedFilters,
		Tags: []string{ENGINE_ALL},
	},
	{
		Name: "Should be able to ignore stop words",
		Fn:   testSearchIgnoringStopWords,
		Tags: []string{ENGINE_ELASTICSEARCH},
	},
	{
		Name: "Should support search stemming",
		Fn:   testSupportStemming,
		Tags: []string{ENGINE_POSTGRES, ENGINE_ELASTICSEARCH},
	},
	{
		Name: "Should support search with wildcards",
		Fn:   testSupportWildcards,
		Tags: []string{ENGINE_POSTGRES, ENGINE_MYSQL, ENGINE_ELASTICSEARCH},
	},
	{
		Name: "Should not support search with preceding wildcards",
		Fn:   testNotSupportPrecedingWildcards,
		Tags: []string{ENGINE_ALL},
	},
	{
		Name: "Should discard a wildcard if it's not placed immediately by text",
		Fn:   testSearchDiscardWildcardAlone,
		Tags: []string{ENGINE_ELASTICSEARCH},
	},
	{
		Name: "Should support terms with dash",
		Fn:   testSupportTermsWithDash,
		Tags: []string{ENGINE_ALL},
		Skip: true,
	},
	{
		Name: "Should support terms with underscore",
		Fn:   testSupportTermsWithUnderscore,
		Tags: []string{ENGINE_MYSQL, ENGINE_ELASTICSEARCH},
	},
	{
		Name: "Should search or exclude post using hashtags",
		Fn:   testSearchOrExcludePostsWithHashtags,
		Tags: []string{ENGINE_ALL},
	},
	{
		Name: "Should support searching for hashtags surrounded by markdown",
		Fn:   testSearchHashtagWithMarkdown,
		Tags: []string{ENGINE_ALL},
	},
	{
		Name: "Should support searching for multiple hashtags",
		Fn:   testSearcWithMultipleHashtags,
		Tags: []string{ENGINE_ELASTICSEARCH},
	},
	{
		Name: "Should support searching hashtags with dots",
		Fn:   testSearchPostsWithDotsInHashtags,
		Tags: []string{ENGINE_ALL},
	},
	{
		Name: "Should be able to search or exclude messages with hashtags in a case insensitive manner",
		Fn:   testSearchHashtagCaseInsensitive,
		Tags: []string{ENGINE_ALL},
	},
	{
		Name: "Should be able to search by hashtags with dashes",
		Fn:   testSearchHashtagWithDash,
		Tags: []string{ENGINE_ALL},
	},
	{
		Name: "Should be able to search by hashtags with numbers",
		Fn:   testSearchHashtagWithNumbers,
		Tags: []string{ENGINE_ALL},
	},
	{
		Name: "Should be able to search by hashtags with dots",
		Fn:   testSearchHashtagWithDots,
		Tags: []string{ENGINE_ALL},
	},
	{
		Name: "Should be able to search by hashtags with underscores",
		Fn:   testSearchHashtagWithUnderscores,
		Tags: []string{ENGINE_ALL},
	},
	{
		Name: "Should not return system messages",
		Fn:   testSearchShouldExcludeSytemMessages,
		Tags: []string{ENGINE_ALL},
	},
	{
		Name: "Should be able to search matching by mentions",
		Fn:   testSearchShouldBeAbleToMatchByMentions,
		Tags: []string{ENGINE_ALL},
	},
	{
		Name:        "Should be able to search in deleted/archived channels",
		Fn:          testSearchInDeletedOrArchivedChannels,
		Tags:        []string{ENGINE_ALL},
		Skip:        true,
		SkipMessage: "Not working",
	},
	{
		Name:        "Should be able to search terms with dashes",
		Fn:          testSearchTermsWithDashes,
		Tags:        []string{ENGINE_ALL},
		Skip:        true,
		SkipMessage: "Not working",
	},
	{
		Name: "Should be able to search terms with dots",
		Fn:   testSearchTermsWithDots,
		Tags: []string{ENGINE_ELASTICSEARCH},
	},
	{
		Name: "Should be able to search terms with underscores",
		Fn:   testSearchTermsWithUnderscores,
		Tags: []string{ENGINE_ELASTICSEARCH},
	},
	{
		Name: "Should be able to search posts made by bot accounts",
		Fn:   testSearchBotAccountsPosts,
		Tags: []string{ENGINE_ALL},
	},
	{
		Name: "Should be able to combine stemming and wildcards",
		Fn:   testSupportStemmingAndWildcards,
		Tags: []string{ENGINE_ELASTICSEARCH},
	},
	{
		Name: "Should support wildcard outside quotes",
		Fn:   testSupportWildcardOutsideQuotes,
		Tags: []string{ENGINE_ELASTICSEARCH},
	},
	{
		Name: "Should support hashtags with 3 or more characters",
		Fn:   testHashtagSearchShouldSupportThreeOrMoreCharacters,
		Tags: []string{ENGINE_ALL},
	},
	{
		Name: "Should not support slash as character separator",
		Fn:   testSlashShouldNotBeCharSeparator,
		Tags: []string{ENGINE_MYSQL, ENGINE_ELASTICSEARCH},
	},
	{
		Name: "Should be able to search emails without quoting them",
		Fn:   testSearchEmailsWithoutQuotes,
		Tags: []string{ENGINE_ELASTICSEARCH},
	},
	{
		Name: "Should be able to search in comments",
		Fn:   testSupportSearchInComments,
		Tags: []string{ENGINE_ALL},
	},
	{
		Name: "Should be able to search terms within links",
		Fn:   testSupportSearchTermsWithinLinks,
		Tags: []string{ENGINE_MYSQL, ENGINE_ELASTICSEARCH},
	},
	{
		Name: "Should not return links that are embedded in markdown",
		Fn:   testShouldNotReturnLinksEmbeddedInMarkdown,
		Tags: []string{ENGINE_POSTGRES, ENGINE_ELASTICSEARCH},
	},
}

func TestSearchPostStore(t *testing.T, s store.Store, testEngine *SearchTestEngine) {
	if s.DriverName() == model.DATABASE_DRIVER_COCKROACH {
		t.Skip("Cockroach db doesn't support full text search")
	}
	th := &SearchTestHelper{
		Store: s,
	}
	err := th.SetupBasicFixtures()
	require.Nil(t, err)
	defer th.CleanFixtures()

	runTestSearch(t, testEngine, searchPostStoreTests, th)
}

func testSearchPostsIncludingDMs(t *testing.T, th *SearchTestHelper) {
	direct, err := th.createDirectChannel(th.Team.Id, "direct", "direct", []*model.User{th.User, th.User2})
	require.Nil(t, err)
	defer th.deleteChannel(direct)

	p1, err := th.createPost(th.User.Id, direct.Id, "dm test", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	_, err = th.createPost(th.User.Id, direct.Id, "dm other", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	p2, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "channel test", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	params := &model.SearchParams{Terms: "test"}
	results, err := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
	require.Nil(t, err)

	require.Len(t, results.Posts, 2)
	th.checkPostInSearchResults(t, p1.Id, results.Posts)
	th.checkPostInSearchResults(t, p2.Id, results.Posts)
}

func testSearchPostsWithPagination(t *testing.T, th *SearchTestHelper) {
	direct, err := th.createDirectChannel(th.Team.Id, "direct", "direct", []*model.User{th.User, th.User2})
	require.Nil(t, err)
	defer th.deleteChannel(direct)

	p1, err := th.createPost(th.User.Id, direct.Id, "dm test", "", model.POST_DEFAULT, 10000, false)
	require.Nil(t, err)
	_, err = th.createPost(th.User.Id, direct.Id, "dm other", "", model.POST_DEFAULT, 20000, false)
	require.Nil(t, err)
	p2, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "channel test", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	params := &model.SearchParams{Terms: "test"}
	results, err := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 1)
	require.Nil(t, err)

	require.Len(t, results.Posts, 1)
	th.checkPostInSearchResults(t, p2.Id, results.Posts)

	results, err = th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 1, 1)
	require.Nil(t, err)

	require.Len(t, results.Posts, 1)
	th.checkPostInSearchResults(t, p1.Id, results.Posts)
}

func testSearchReturnPinnedAndUnpinned(t *testing.T, th *SearchTestHelper) {
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "channel test unpinned", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	p2, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "channel test pinned", "", model.POST_DEFAULT, 0, true)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	params := &model.SearchParams{Terms: "test"}
	results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
	require.Nil(t, apperr)

	require.Len(t, results.Posts, 2)
	th.checkPostInSearchResults(t, p1.Id, results.Posts)
	th.checkPostInSearchResults(t, p2.Id, results.Posts)
}

func testSearchExactPhraseInQuotes(t *testing.T, th *SearchTestHelper) {
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "channel test 1 2 3", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	_, err = th.createPost(th.User.Id, th.ChannelBasic.Id, "channel test 123", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	params := &model.SearchParams{Terms: "\"channel test 1 2 3\""}
	results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
	require.Nil(t, apperr)

	require.Len(t, results.Posts, 1)
	th.checkPostInSearchResults(t, p1.Id, results.Posts)
}

func testSearchEmailAddresses(t *testing.T, th *SearchTestHelper) {
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "test email test@test.com", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	_, err = th.createPost(th.User.Id, th.ChannelBasic.Id, "test email test2@test.com", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	t.Run("Should search email addresses enclosed by quotes", func(t *testing.T) {
		params := &model.SearchParams{Terms: "\"test@test.com\""}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 1)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
	})

	t.Run("Should search email addresses without quotes", func(t *testing.T) {
		params := &model.SearchParams{Terms: "test@test.com"}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 1)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
	})
}

func testSearchMarkdownUnderscores(t *testing.T, th *SearchTestHelper) {
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "_start middle end_ _another_", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	t.Run("Should search the start inside the markdown underscore", func(t *testing.T) {
		params := &model.SearchParams{Terms: "start"}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 1)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
	})

	t.Run("Should search a word in the middle of the markdown underscore", func(t *testing.T) {
		params := &model.SearchParams{Terms: "middle"}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 1)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
	})

	t.Run("Should search in the end of the markdown underscore", func(t *testing.T) {
		params := &model.SearchParams{Terms: "end"}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 1)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
	})

	t.Run("Should search inside markdown underscore", func(t *testing.T) {
		params := &model.SearchParams{Terms: "another"}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 1)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
	})
}

func testSearchNonLatinWords(t *testing.T, th *SearchTestHelper) {
	t.Run("Should be able to search chinese words", func(t *testing.T) {
		p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "你好", "", model.POST_DEFAULT, 0, false)
		require.Nil(t, err)
		p2, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "你", "", model.POST_DEFAULT, 0, false)
		require.Nil(t, err)
		defer th.deleteUserPosts(th.User.Id)

		t.Run("Should search one word", func(t *testing.T) {
			params := &model.SearchParams{Terms: "你"}
			results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
			require.Nil(t, apperr)

			require.Len(t, results.Posts, 1)
			th.checkPostInSearchResults(t, p2.Id, results.Posts)
		})
		t.Run("Should search two words", func(t *testing.T) {
			params := &model.SearchParams{Terms: "你好"}
			results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
			require.Nil(t, apperr)

			require.Len(t, results.Posts, 1)
			th.checkPostInSearchResults(t, p1.Id, results.Posts)
		})
		t.Run("Should search with wildcard", func(t *testing.T) {
			params := &model.SearchParams{Terms: "你*"}
			results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
			require.Nil(t, apperr)

			require.Len(t, results.Posts, 2)
			th.checkPostInSearchResults(t, p1.Id, results.Posts)
			th.checkPostInSearchResults(t, p2.Id, results.Posts)
		})
	})
	t.Run("Should be able to search cyrillic words", func(t *testing.T) {
		p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "слово test", "", model.POST_DEFAULT, 0, false)
		require.Nil(t, err)
		defer th.deleteUserPosts(th.User.Id)

		t.Run("Should search one word", func(t *testing.T) {
			params := &model.SearchParams{Terms: "слово"}
			results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
			require.Nil(t, apperr)

			require.Len(t, results.Posts, 1)
			th.checkPostInSearchResults(t, p1.Id, results.Posts)
		})
		t.Run("Should search using wildcard", func(t *testing.T) {
			params := &model.SearchParams{Terms: "слов*"}
			results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
			require.Nil(t, apperr)

			require.Len(t, results.Posts, 1)
			th.checkPostInSearchResults(t, p1.Id, results.Posts)
		})
	})

	t.Run("Should be able to search japanese words", func(t *testing.T) {
		p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "本", "", model.POST_DEFAULT, 0, false)
		require.Nil(t, err)
		p2, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "本木", "", model.POST_DEFAULT, 0, false)
		require.Nil(t, err)
		defer th.deleteUserPosts(th.User.Id)

		t.Run("Should search one word", func(t *testing.T) {
			params := &model.SearchParams{Terms: "本"}
			results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
			require.Nil(t, apperr)

			require.Len(t, results.Posts, 2)
			th.checkPostInSearchResults(t, p1.Id, results.Posts)
			th.checkPostInSearchResults(t, p2.Id, results.Posts)
		})
		t.Run("Should search two words", func(t *testing.T) {
			params := &model.SearchParams{Terms: "本木"}
			results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
			require.Nil(t, apperr)

			require.Len(t, results.Posts, 1)
			th.checkPostInSearchResults(t, p2.Id, results.Posts)
		})
		t.Run("Should search with wildcard", func(t *testing.T) {
			params := &model.SearchParams{Terms: "本*"}
			results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
			require.Nil(t, apperr)

			require.Len(t, results.Posts, 2)
			th.checkPostInSearchResults(t, p1.Id, results.Posts)
			th.checkPostInSearchResults(t, p2.Id, results.Posts)
		})
	})

	t.Run("Should be able to search korean words", func(t *testing.T) {
		p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "불", "", model.POST_DEFAULT, 0, false)
		require.Nil(t, err)
		p2, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "불다", "", model.POST_DEFAULT, 0, false)
		require.Nil(t, err)
		defer th.deleteUserPosts(th.User.Id)

		t.Run("Should search one word", func(t *testing.T) {
			params := &model.SearchParams{Terms: "불"}
			results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
			require.Nil(t, apperr)

			require.Len(t, results.Posts, 1)
			th.checkPostInSearchResults(t, p1.Id, results.Posts)
		})
		t.Run("Should search two words", func(t *testing.T) {
			params := &model.SearchParams{Terms: "불다"}
			results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
			require.Nil(t, apperr)

			require.Len(t, results.Posts, 1)
			th.checkPostInSearchResults(t, p2.Id, results.Posts)
		})
		t.Run("Should search with wildcard", func(t *testing.T) {
			params := &model.SearchParams{Terms: "불*"}
			results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
			require.Nil(t, apperr)

			require.Len(t, results.Posts, 2)
			th.checkPostInSearchResults(t, p1.Id, results.Posts)
			th.checkPostInSearchResults(t, p2.Id, results.Posts)
		})
	})
}

func testSearchAlternativeSpellings(t *testing.T, th *SearchTestHelper) {
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "Straße test", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	p2, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "Strasse test", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	params := &model.SearchParams{Terms: "Straße"}
	results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
	require.Nil(t, apperr)

	require.Len(t, results.Posts, 2)
	th.checkPostInSearchResults(t, p1.Id, results.Posts)
	th.checkPostInSearchResults(t, p2.Id, results.Posts)

	params = &model.SearchParams{Terms: "Strasse"}
	results, apperr = th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
	require.Nil(t, apperr)

	require.Len(t, results.Posts, 2)
	th.checkPostInSearchResults(t, p1.Id, results.Posts)
	th.checkPostInSearchResults(t, p2.Id, results.Posts)
}

func testSearchAlternativeSpellingsAccents(t *testing.T, th *SearchTestHelper) {
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "café", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	p2, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "café", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	params := &model.SearchParams{Terms: "café"}
	results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
	require.Nil(t, apperr)

	require.Len(t, results.Posts, 2)
	th.checkPostInSearchResults(t, p1.Id, results.Posts)
	th.checkPostInSearchResults(t, p2.Id, results.Posts)

	params = &model.SearchParams{Terms: "café"}
	results, apperr = th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
	require.Nil(t, apperr)

	require.Len(t, results.Posts, 2)
	th.checkPostInSearchResults(t, p1.Id, results.Posts)
	th.checkPostInSearchResults(t, p2.Id, results.Posts)

	params = &model.SearchParams{Terms: "cafe"}
	results, apperr = th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
	require.Nil(t, apperr)

	require.Len(t, results.Posts, 0)
}

func testSearchOrExcludePostsBySpecificUser(t *testing.T, th *SearchTestHelper) {
	p1, err := th.createPost(th.User.Id, th.ChannelPrivate.Id, "test fromuser", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	_, err = th.createPost(th.User2.Id, th.ChannelPrivate.Id, "test fromuser 2", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)
	defer th.deleteUserPosts(th.User2.Id)

	params := &model.SearchParams{
		Terms:     "fromuser",
		FromUsers: []string{th.User.Id},
	}
	results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
	require.Nil(t, apperr)

	require.Len(t, results.Posts, 1)
	th.checkPostInSearchResults(t, p1.Id, results.Posts)
}

func testSearchOrExcludePostsInChannel(t *testing.T, th *SearchTestHelper) {
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "test fromuser", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	_, err = th.createPost(th.User2.Id, th.ChannelPrivate.Id, "test fromuser 2", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)
	defer th.deleteUserPosts(th.User2.Id)

	params := &model.SearchParams{
		Terms:      "fromuser",
		InChannels: []string{th.ChannelBasic.Id},
	}
	results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
	require.Nil(t, apperr)

	require.Len(t, results.Posts, 1)
	th.checkPostInSearchResults(t, p1.Id, results.Posts)
}

func testSearchOrExcludePostsInDMGM(t *testing.T, th *SearchTestHelper) {
	direct, err := th.createDirectChannel(th.Team.Id, "direct", "direct", []*model.User{th.User, th.User2})
	require.Nil(t, err)
	defer th.deleteChannel(direct)

	group, err := th.createGroupChannel(th.Team.Id, "test group", []*model.User{th.User, th.User2})
	require.Nil(t, err)
	defer th.deleteChannel(group)

	p1, err := th.createPost(th.User.Id, direct.Id, "test fromuser", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	p2, err := th.createPost(th.User2.Id, group.Id, "test fromuser 2", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)
	defer th.deleteUserPosts(th.User2.Id)

	t.Run("Should be able to search in both DM and GM channels", func(t *testing.T) {
		params := &model.SearchParams{
			Terms:      "fromuser",
			InChannels: []string{direct.Id, group.Id},
		}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 2)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
		th.checkPostInSearchResults(t, p2.Id, results.Posts)
	})

	t.Run("Should be able to search only in DM channel", func(t *testing.T) {
		params := &model.SearchParams{
			Terms:      "fromuser",
			InChannels: []string{direct.Id},
		}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 1)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
	})

	t.Run("Should be able to search only in GM channel", func(t *testing.T) {
		params := &model.SearchParams{
			Terms:      "fromuser",
			InChannels: []string{group.Id},
		}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 1)
		th.checkPostInSearchResults(t, p2.Id, results.Posts)
	})
}

func testFilterMessagesInSpecificDate(t *testing.T, th *SearchTestHelper) {
	creationDate := model.GetMillisForTime(time.Date(2020, 03, 22, 12, 0, 0, 0, time.UTC))
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "test in specific date", "", model.POST_DEFAULT, creationDate, false)
	require.Nil(t, err)
	creationDate2 := model.GetMillisForTime(time.Date(2020, 03, 23, 0, 0, 0, 0, time.UTC))
	p2, err := th.createPost(th.User.Id, th.ChannelPrivate.Id, "test in the present", "", model.POST_DEFAULT, creationDate2, false)
	require.Nil(t, err)
	creationDate3 := model.GetMillisForTime(time.Date(2020, 03, 21, 23, 59, 59, 0, time.UTC))
	p3, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "test in the present", "", model.POST_DEFAULT, creationDate3, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	t.Run("Should be able to search posts on date", func(t *testing.T) {
		params := &model.SearchParams{
			Terms:  "test",
			OnDate: "2020-03-22",
		}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 1)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
	})
	t.Run("Should be able to exclude posts on date", func(t *testing.T) {
		params := &model.SearchParams{
			Terms:        "test",
			ExcludedDate: "2020-03-22",
		}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 2)
		th.checkPostInSearchResults(t, p2.Id, results.Posts)
		th.checkPostInSearchResults(t, p3.Id, results.Posts)
	})
}

func testFilterMessagesBeforeSpecificDate(t *testing.T, th *SearchTestHelper) {
	creationDate := model.GetMillisForTime(time.Date(2020, 03, 01, 12, 0, 0, 0, time.UTC))
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "test in specific date", "", model.POST_DEFAULT, creationDate, false)
	require.Nil(t, err)
	creationDate2 := model.GetMillisForTime(time.Date(2020, 03, 22, 23, 59, 59, 0, time.UTC))
	p2, err := th.createPost(th.User.Id, th.ChannelPrivate.Id, "test in specific date 2", "", model.POST_DEFAULT, creationDate2, false)
	require.Nil(t, err)
	creationDate3 := model.GetMillisForTime(time.Date(2020, 03, 26, 16, 55, 0, 0, time.UTC))
	p3, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "test in the present", "", model.POST_DEFAULT, creationDate3, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	t.Run("Should be able to search posts before a date", func(t *testing.T) {
		params := &model.SearchParams{
			Terms:      "test",
			BeforeDate: "2020-03-23",
		}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 2)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
		th.checkPostInSearchResults(t, p2.Id, results.Posts)
	})

	t.Run("Should be able to exclude posts before a date", func(t *testing.T) {
		params := &model.SearchParams{
			Terms:              "test",
			ExcludedBeforeDate: "2020-03-23",
		}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 1)
		th.checkPostInSearchResults(t, p3.Id, results.Posts)
	})
}

func testFilterMessagesAfterSpecificDate(t *testing.T, th *SearchTestHelper) {
	creationDate := model.GetMillisForTime(time.Date(2020, 03, 01, 12, 0, 0, 0, time.UTC))
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "test in specific date", "", model.POST_DEFAULT, creationDate, false)
	require.Nil(t, err)
	creationDate2 := model.GetMillisForTime(time.Date(2020, 03, 22, 23, 59, 59, 0, time.UTC))
	p2, err := th.createPost(th.User.Id, th.ChannelPrivate.Id, "test in specific date 2", "", model.POST_DEFAULT, creationDate2, false)
	require.Nil(t, err)
	creationDate3 := model.GetMillisForTime(time.Date(2020, 03, 26, 16, 55, 0, 0, time.UTC))
	p3, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "test in the present", "", model.POST_DEFAULT, creationDate3, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	t.Run("Should be able to search posts after a date", func(t *testing.T) {
		params := &model.SearchParams{
			Terms:     "test",
			AfterDate: "2020-03-23",
		}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 1)
		th.checkPostInSearchResults(t, p3.Id, results.Posts)
	})

	t.Run("Should be able to exclude posts after a date", func(t *testing.T) {
		params := &model.SearchParams{
			Terms:             "test",
			ExcludedAfterDate: "2020-03-23",
		}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 2)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
		th.checkPostInSearchResults(t, p2.Id, results.Posts)
	})
}

func testFilterMessagesWithATerm(t *testing.T, th *SearchTestHelper) {
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "one two three", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	p2, err := th.createPost(th.User.Id, th.ChannelPrivate.Id, "one four five six", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	_, err = th.createPost(th.User.Id, th.ChannelBasic.Id, "one seven eight nine", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	t.Run("Should exclude terms", func(t *testing.T) {
		params := &model.SearchParams{
			Terms:         "one",
			ExcludedTerms: "five eight",
		}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 1)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
	})

	t.Run("Should exclude quoted terms", func(t *testing.T) {
		params := &model.SearchParams{
			Terms:         "one",
			ExcludedTerms: "\"eight nine\"",
		}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 2)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
		th.checkPostInSearchResults(t, p2.Id, results.Posts)
	})
}

func testSearchUsingBooleanOperators(t *testing.T, th *SearchTestHelper) {
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "one two three message", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	p2, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "two messages", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	_, err = th.createPost(th.User.Id, th.ChannelBasic.Id, "another message", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	t.Run("Should search posts using OR operator", func(t *testing.T) {
		params := &model.SearchParams{
			Terms:   "one two",
			OrTerms: true,
		}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 2)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
		th.checkPostInSearchResults(t, p2.Id, results.Posts)
	})

	t.Run("Should search posts using AND operator", func(t *testing.T) {
		params := &model.SearchParams{
			Terms:   "one two",
			OrTerms: false,
		}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 1)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
	})
}

func testSearchUsingCombinedFilters(t *testing.T, th *SearchTestHelper) {
	creationDate := model.GetMillisForTime(time.Date(2020, 03, 01, 12, 0, 0, 0, time.UTC))
	p1, err := th.createPost(th.User.Id, th.ChannelPrivate.Id, "one two three message", "", model.POST_DEFAULT, creationDate, false)
	require.Nil(t, err)
	creationDate2 := model.GetMillisForTime(time.Date(2020, 03, 10, 12, 0, 0, 0, time.UTC))
	p2, err := th.createPost(th.User2.Id, th.ChannelPrivate.Id, "two messages", "", model.POST_DEFAULT, creationDate2, false)
	require.Nil(t, err)
	creationDate3 := model.GetMillisForTime(time.Date(2020, 03, 20, 12, 0, 0, 0, time.UTC))
	p3, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "two another message", "", model.POST_DEFAULT, creationDate3, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)
	defer th.deleteUserPosts(th.User2.Id)

	t.Run("Should search combining from user and in channel filters", func(t *testing.T) {
		params := &model.SearchParams{
			Terms:      "two",
			FromUsers:  []string{th.User2.Id},
			InChannels: []string{th.ChannelPrivate.Id},
		}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 1)
		th.checkPostInSearchResults(t, p2.Id, results.Posts)
	})

	t.Run("Should search combining excluding users and in channel filters", func(t *testing.T) {
		params := &model.SearchParams{
			Terms:         "two",
			ExcludedUsers: []string{th.User2.Id},
			InChannels:    []string{th.ChannelPrivate.Id},
		}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 1)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
	})

	t.Run("Should search combining excluding dates and in channel filters", func(t *testing.T) {
		params := &model.SearchParams{
			Terms:              "two",
			ExcludedBeforeDate: "2020-03-09",
			ExcludedAfterDate:  "2020-03-11",
			InChannels:         []string{th.ChannelPrivate.Id},
		}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 1)
		th.checkPostInSearchResults(t, p2.Id, results.Posts)
	})
	t.Run("Should search combining excluding dates and in channel filters", func(t *testing.T) {
		params := &model.SearchParams{
			Terms:            "two",
			AfterDate:        "2020-03-11",
			ExcludedChannels: []string{th.ChannelPrivate.Id},
		}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 1)
		th.checkPostInSearchResults(t, p3.Id, results.Posts)
	})
}

func testSearchIgnoringStopWords(t *testing.T, th *SearchTestHelper) {
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "the search for a bunch of stop words", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	p2, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "the objective is to avoid a bunch of stop words", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	p3, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "in the a on to where you", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	t.Run("Should avoid stop word 'the'", func(t *testing.T) {
		params := &model.SearchParams{
			Terms: "the search",
		}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 1)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
	})

	t.Run("Should avoid stop word 'a'", func(t *testing.T) {
		params := &model.SearchParams{
			Terms: "a avoid",
		}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 1)
		th.checkPostInSearchResults(t, p2.Id, results.Posts)
	})

	t.Run("Should avoid stop word 'in'", func(t *testing.T) {
		params := &model.SearchParams{
			Terms: "in where",
		}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 1)
		th.checkPostInSearchResults(t, p3.Id, results.Posts)
	})
}

func testSupportStemming(t *testing.T, th *SearchTestHelper) {
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "search post", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	p2, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "searching post", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	_, err = th.createPost(th.User.Id, th.ChannelBasic.Id, "another post", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	params := &model.SearchParams{
		Terms: "search",
	}
	results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
	require.Nil(t, apperr)

	require.Len(t, results.Posts, 2)
	th.checkPostInSearchResults(t, p1.Id, results.Posts)
	th.checkPostInSearchResults(t, p2.Id, results.Posts)
}

func testSupportWildcards(t *testing.T, th *SearchTestHelper) {
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "search post", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	p2, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "searching post", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	_, err = th.createPost(th.User.Id, th.ChannelBasic.Id, "another post", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	params := &model.SearchParams{
		Terms: "search*",
	}
	results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
	require.Nil(t, apperr)

	require.Len(t, results.Posts, 2)
	th.checkPostInSearchResults(t, p1.Id, results.Posts)
	th.checkPostInSearchResults(t, p2.Id, results.Posts)
}

func testNotSupportPrecedingWildcards(t *testing.T, th *SearchTestHelper) {
	_, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "search post", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	_, err = th.createPost(th.User.Id, th.ChannelBasic.Id, "searching post", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	_, err = th.createPost(th.User.Id, th.ChannelBasic.Id, "another post", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	params := &model.SearchParams{
		Terms: "*earch",
	}
	results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
	require.Nil(t, apperr)

	require.Len(t, results.Posts, 0)
}

func testSearchDiscardWildcardAlone(t *testing.T, th *SearchTestHelper) {
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "qwerty", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	_, err = th.createPost(th.User.Id, th.ChannelBasic.Id, "qwertyjkl", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	params := &model.SearchParams{
		Terms: "qwerty *",
	}
	results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
	require.Nil(t, apperr)

	require.Len(t, results.Posts, 1)
	th.checkPostInSearchResults(t, p1.Id, results.Posts)
}

func testSupportTermsWithDash(t *testing.T, th *SearchTestHelper) {
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "search term-with-dash", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	_, err = th.createPost(th.User.Id, th.ChannelBasic.Id, "searching term with dash", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	t.Run("Should search terms with dash", func(t *testing.T) {
		params := &model.SearchParams{
			Terms: "term-with-dash",
		}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 1)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
	})

	t.Run("Should search terms with dash using quotes", func(t *testing.T) {
		params := &model.SearchParams{
			Terms: "\"term-with-dash\"",
		}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 1)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
	})
}

func testSupportTermsWithUnderscore(t *testing.T, th *SearchTestHelper) {
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "search term_with_underscore", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	_, err = th.createPost(th.User.Id, th.ChannelBasic.Id, "searching term with underscore", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	t.Run("Should search terms with underscore", func(t *testing.T) {
		params := &model.SearchParams{
			Terms: "term_with_underscore",
		}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 1)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
	})

	t.Run("Should search terms with underscore using quotes", func(t *testing.T) {
		params := &model.SearchParams{
			Terms: "\"term_with_underscore\"",
		}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 1)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
	})
}

func testSearchOrExcludePostsWithHashtags(t *testing.T, th *SearchTestHelper) {
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "search post with #hashtag", "#hashtag", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	p2, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "searching term with hashtag", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	p3, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "searching term with", "#hashtag", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	t.Run("Should search terms with hashtags", func(t *testing.T) {
		params := &model.SearchParams{
			Terms:     "#hashtag",
			IsHashtag: true,
		}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 2)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
		th.checkPostInSearchResults(t, p3.Id, results.Posts)
	})

	t.Run("Should search hashtag terms without hashtag option", func(t *testing.T) {
		params := &model.SearchParams{
			Terms:     "#hashtag",
			IsHashtag: false,
		}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 2)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
		th.checkPostInSearchResults(t, p2.Id, results.Posts)
	})
}

func testSearchHashtagWithMarkdown(t *testing.T, th *SearchTestHelper) {
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "searching hashtag #hashtag", "#hashtag", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	p2, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "searching term with `#hashtag`", "#hashtag", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	p3, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "searching term with **#hashtag**", "#hashtag", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	p4, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "searching term with ~~#hashtag~~", "#hashtag", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	p5, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "searching term with _#hashtag_", "#hashtag", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	params := &model.SearchParams{
		Terms:     "#hashtag",
		IsHashtag: true,
	}
	results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
	require.Nil(t, apperr)

	require.Len(t, results.Posts, 5)
	th.checkPostInSearchResults(t, p1.Id, results.Posts)
	th.checkPostInSearchResults(t, p2.Id, results.Posts)
	th.checkPostInSearchResults(t, p3.Id, results.Posts)
	th.checkPostInSearchResults(t, p4.Id, results.Posts)
	th.checkPostInSearchResults(t, p5.Id, results.Posts)
}

func testSearcWithMultipleHashtags(t *testing.T, th *SearchTestHelper) {
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "searching hashtag #hashtag", "#hashtwo #hashone", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	p2, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "searching term with `#hashtag`", "#hashtwo #hashthree", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	t.Run("Should search posts with multiple hashtags", func(t *testing.T) {
		params := &model.SearchParams{
			Terms:     "#hashone #hashtwo",
			IsHashtag: true,
		}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 1)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
	})

	t.Run("Should search posts with multiple hashtags using OR", func(t *testing.T) {
		params := &model.SearchParams{
			Terms:     "#hashone #hashtwo",
			IsHashtag: true,
			OrTerms:   true,
		}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 2)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
		th.checkPostInSearchResults(t, p2.Id, results.Posts)
	})
}

func testSearchPostsWithDotsInHashtags(t *testing.T, th *SearchTestHelper) {
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "searching hashtag #hashtag.dot", "#hashtag.dot", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	params := &model.SearchParams{
		Terms:     "#hashtag.dot",
		IsHashtag: true,
	}
	results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
	require.Nil(t, apperr)

	require.Len(t, results.Posts, 1)
	th.checkPostInSearchResults(t, p1.Id, results.Posts)
}

func testSearchHashtagCaseInsensitive(t *testing.T, th *SearchTestHelper) {
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "searching hashtag #HaShTaG", "#HaShTaG", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	p2, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "searching hashtag #hashtag", "#hashtag", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	p3, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "searching hashtag #HASHTAG", "#HASHTAG", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	t.Run("Lower case hashtag", func(t *testing.T) {
		params := &model.SearchParams{
			Terms:     "#hashtag",
			IsHashtag: true,
		}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 3)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
		th.checkPostInSearchResults(t, p2.Id, results.Posts)
		th.checkPostInSearchResults(t, p3.Id, results.Posts)
	})

	t.Run("Upper case hashtag", func(t *testing.T) {
		params := &model.SearchParams{
			Terms:     "#HASHTAG",
			IsHashtag: true,
		}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 3)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
		th.checkPostInSearchResults(t, p2.Id, results.Posts)
		th.checkPostInSearchResults(t, p3.Id, results.Posts)
	})

	t.Run("Mixed case hashtag", func(t *testing.T) {
		params := &model.SearchParams{
			Terms:     "#HaShTaG",
			IsHashtag: true,
		}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 3)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
		th.checkPostInSearchResults(t, p2.Id, results.Posts)
		th.checkPostInSearchResults(t, p3.Id, results.Posts)
	})
}

func testSearchHashtagWithDash(t *testing.T, th *SearchTestHelper) {
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "searching hashtag #hashtag-test", "#hashtag-test", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	_, err = th.createPost(th.User.Id, th.ChannelBasic.Id, "searching hashtag #hashtagtest", "#hashtagtest", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	params := &model.SearchParams{
		Terms:     "#hashtag-test",
		IsHashtag: true,
	}
	results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
	require.Nil(t, apperr)

	require.Len(t, results.Posts, 1)
	th.checkPostInSearchResults(t, p1.Id, results.Posts)
}

func testSearchHashtagWithNumbers(t *testing.T, th *SearchTestHelper) {
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "searching hashtag #h4sht4g", "#h4sht4g", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	_, err = th.createPost(th.User.Id, th.ChannelBasic.Id, "searching hashtag #hashtag", "#hashtag", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	params := &model.SearchParams{
		Terms:     "#h4sht4g",
		IsHashtag: true,
	}
	results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
	require.Nil(t, apperr)

	require.Len(t, results.Posts, 1)
	th.checkPostInSearchResults(t, p1.Id, results.Posts)
}

func testSearchHashtagWithDots(t *testing.T, th *SearchTestHelper) {
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "searching hashtag #hashtag.test", "#hashtag.test", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	_, err = th.createPost(th.User.Id, th.ChannelBasic.Id, "searching hashtag #hashtagtest", "#hashtagtest", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	params := &model.SearchParams{
		Terms:     "#hashtag.test",
		IsHashtag: true,
	}
	results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
	require.Nil(t, apperr)

	require.Len(t, results.Posts, 1)
	th.checkPostInSearchResults(t, p1.Id, results.Posts)
}

func testSearchHashtagWithUnderscores(t *testing.T, th *SearchTestHelper) {
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "searching hashtag #hashtag_test", "#hashtag_test", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	_, err = th.createPost(th.User.Id, th.ChannelBasic.Id, "searching hashtag #hashtagtest", "#hashtagtest", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	params := &model.SearchParams{
		Terms:     "#hashtag_test",
		IsHashtag: true,
	}
	results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
	require.Nil(t, apperr)

	require.Len(t, results.Posts, 1)
	th.checkPostInSearchResults(t, p1.Id, results.Posts)
}

func testSearchShouldExcludeSytemMessages(t *testing.T, th *SearchTestHelper) {
	_, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "test system message one", "", model.POST_JOIN_CHANNEL, 0, false)
	require.Nil(t, err)
	_, err = th.createPost(th.User.Id, th.ChannelBasic.Id, "test system message two", "", model.POST_LEAVE_CHANNEL, 0, false)
	require.Nil(t, err)
	_, err = th.createPost(th.User.Id, th.ChannelBasic.Id, "test system message three", "", model.POST_LEAVE_TEAM, 0, false)
	require.Nil(t, err)
	_, err = th.createPost(th.User.Id, th.ChannelBasic.Id, "test system message four", "", model.POST_ADD_TO_CHANNEL, 0, false)
	require.Nil(t, err)
	_, err = th.createPost(th.User.Id, th.ChannelBasic.Id, "test system message five", "", model.POST_ADD_TO_TEAM, 0, false)
	require.Nil(t, err)
	_, err = th.createPost(th.User.Id, th.ChannelBasic.Id, "test system message six", "", model.POST_HEADER_CHANGE, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	params := &model.SearchParams{Terms: "test system"}
	results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
	require.Nil(t, apperr)

	require.Len(t, results.Posts, 0)
}

func testSearchShouldBeAbleToMatchByMentions(t *testing.T, th *SearchTestHelper) {
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "test system @testuser", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	p2, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "test system testuser", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	p3, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "test system #testuser", "#testuser", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	params := &model.SearchParams{Terms: "@testuser"}
	results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
	require.Nil(t, apperr)

	require.Len(t, results.Posts, 3)
	th.checkPostInSearchResults(t, p1.Id, results.Posts)
	th.checkPostInSearchResults(t, p2.Id, results.Posts)
	th.checkPostInSearchResults(t, p3.Id, results.Posts)
}

func testSearchInDeletedOrArchivedChannels(t *testing.T, th *SearchTestHelper) {
	p1, err := th.createPost(th.User.Id, th.ChannelDeleted.Id, "message in deleted channel", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	p2, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "message in regular channel", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	p3, err := th.createPost(th.User.Id, th.ChannelPrivate.Id, "message in private channel", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	t.Run("Doesn't include posts in deleted channels", func(t *testing.T) {
		params := &model.SearchParams{Terms: "message", IncludeDeletedChannels: false}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 2)
		th.checkPostInSearchResults(t, p2.Id, results.Posts)
		th.checkPostInSearchResults(t, p3.Id, results.Posts)
	})

	t.Run("Include posts in deleted channels", func(t *testing.T) {
		params := &model.SearchParams{Terms: "message", IncludeDeletedChannels: true}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 3)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
		th.checkPostInSearchResults(t, p2.Id, results.Posts)
		th.checkPostInSearchResults(t, p3.Id, results.Posts)
	})

	t.Run("Include posts in deleted channels using multiple terms", func(t *testing.T) {
		params := &model.SearchParams{Terms: "message channel", IncludeDeletedChannels: true}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 3)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
		th.checkPostInSearchResults(t, p2.Id, results.Posts)
		th.checkPostInSearchResults(t, p3.Id, results.Posts)
	})

	t.Run("Include posts in deleted channels using multiple OR terms", func(t *testing.T) {
		params := &model.SearchParams{
			Terms:                  "message channel",
			IncludeDeletedChannels: true,
			OrTerms:                true,
		}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 3)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
		th.checkPostInSearchResults(t, p2.Id, results.Posts)
		th.checkPostInSearchResults(t, p3.Id, results.Posts)
	})
}

func testSearchTermsWithDashes(t *testing.T, th *SearchTestHelper) {
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "message with-dash-term", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	p2, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "message with dash term", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	t.Run("Search for terms with dash", func(t *testing.T) {
		params := &model.SearchParams{Terms: "with-dash-term"}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 1)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
	})

	t.Run("Search for terms with quoted dash", func(t *testing.T) {
		params := &model.SearchParams{Terms: "\"with-dash-term\""}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 1)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
	})

	t.Run("Search for multiple terms with one having dash", func(t *testing.T) {
		params := &model.SearchParams{Terms: "with-dash-term message"}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 1)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
	})

	t.Run("Search for multiple OR terms with one having dash", func(t *testing.T) {
		params := &model.SearchParams{Terms: "with-dash-term message", OrTerms: true}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 2)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
		th.checkPostInSearchResults(t, p2.Id, results.Posts)
	})
}

func testSearchTermsWithDots(t *testing.T, th *SearchTestHelper) {
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "message with.dots.term", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	p2, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "message with dots term", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	t.Run("Search for terms with dots", func(t *testing.T) {
		params := &model.SearchParams{Terms: "with.dots.term"}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 1)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
	})

	t.Run("Search for terms with quoted dots", func(t *testing.T) {
		params := &model.SearchParams{Terms: "\"with.dots.term\""}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 1)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
	})

	t.Run("Search for multiple terms with one having dots", func(t *testing.T) {
		params := &model.SearchParams{Terms: "with.dots.term message"}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 1)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
	})

	t.Run("Search for multiple OR terms with one having dots", func(t *testing.T) {
		params := &model.SearchParams{Terms: "with.dots.term message", OrTerms: true}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 2)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
		th.checkPostInSearchResults(t, p2.Id, results.Posts)
	})
}

func testSearchTermsWithUnderscores(t *testing.T, th *SearchTestHelper) {
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "message with_underscores_term", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	p2, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "message with underscores term", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	t.Run("Search for terms with underscores", func(t *testing.T) {
		params := &model.SearchParams{Terms: "with_underscores_term"}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 1)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
	})

	t.Run("Search for terms with quoted underscores", func(t *testing.T) {
		params := &model.SearchParams{Terms: "\"with_underscores_term\""}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 1)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
	})

	t.Run("Search for multiple terms with one having underscores", func(t *testing.T) {
		params := &model.SearchParams{Terms: "with_underscores_term message"}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 1)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
	})

	t.Run("Search for multiple OR terms with one having underscores", func(t *testing.T) {
		params := &model.SearchParams{Terms: "with_underscores_term message", OrTerms: true}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 2)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
		th.checkPostInSearchResults(t, p2.Id, results.Posts)
	})
}

func testSearchBotAccountsPosts(t *testing.T, th *SearchTestHelper) {
	bot, err := th.createBot("testbot", "Test Bot", th.User.Id)
	require.Nil(t, err)
	defer th.deleteBot(bot.UserId)
	err = th.addUserToTeams(model.UserFromBot(bot), []string{th.Team.Id})
	require.Nil(t, err)
	p1, err := th.createPost(bot.UserId, th.ChannelBasic.Id, "bot test message", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	p2, err := th.createPost(bot.UserId, th.ChannelPrivate.Id, "bot test message in private", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(bot.UserId)

	params := &model.SearchParams{Terms: "bot"}
	results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
	require.Nil(t, apperr)

	require.Len(t, results.Posts, 2)
	th.checkPostInSearchResults(t, p1.Id, results.Posts)
	th.checkPostInSearchResults(t, p2.Id, results.Posts)
}

func testSupportStemmingAndWildcards(t *testing.T, th *SearchTestHelper) {
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "approve", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	p2, err := th.createPost(th.User.Id, th.ChannelPrivate.Id, "approved", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	p3, err := th.createPost(th.User.Id, th.ChannelPrivate.Id, "approvedz", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	t.Run("Should stem appr", func(t *testing.T) {
		params := &model.SearchParams{Terms: "appr*"}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 3)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
		th.checkPostInSearchResults(t, p2.Id, results.Posts)
		th.checkPostInSearchResults(t, p3.Id, results.Posts)
	})

	t.Run("Should stem approve", func(t *testing.T) {
		params := &model.SearchParams{Terms: "approve*"}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 1)
		th.checkPostInSearchResults(t, p3.Id, results.Posts)
	})
}

func testSupportWildcardOutsideQuotes(t *testing.T, th *SearchTestHelper) {
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "hello world", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	p2, err := th.createPost(th.User.Id, th.ChannelPrivate.Id, "hell or heaven", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	t.Run("Should return results without quotes", func(t *testing.T) {
		params := &model.SearchParams{Terms: "hell*"}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 2)
		th.checkPostInSearchResults(t, p1.Id, results.Posts)
		th.checkPostInSearchResults(t, p2.Id, results.Posts)
	})

	t.Run("Should return just one result with quotes", func(t *testing.T) {
		params := &model.SearchParams{Terms: "\"hell\"*"}
		results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
		require.Nil(t, apperr)

		require.Len(t, results.Posts, 1)
		th.checkPostInSearchResults(t, p2.Id, results.Posts)
	})

}

func testHashtagSearchShouldSupportThreeOrMoreCharacters(t *testing.T, th *SearchTestHelper) {
	_, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "one char hashtag #1", "#1", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	_, err = th.createPost(th.User.Id, th.ChannelPrivate.Id, "two chars hashtag #12", "#12", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	p3, err := th.createPost(th.User.Id, th.ChannelPrivate.Id, "three chars hashtag #123", "#123", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	params := &model.SearchParams{Terms: "#123", IsHashtag: true}
	results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
	require.Nil(t, apperr)

	require.Len(t, results.Posts, 1)
	th.checkPostInSearchResults(t, p3.Id, results.Posts)
}

func testSlashShouldNotBeCharSeparator(t *testing.T, th *SearchTestHelper) {
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "alpha/beta gamma, theta", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	params := &model.SearchParams{Terms: "gamma"}
	results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
	require.Nil(t, apperr)

	require.Len(t, results.Posts, 1)
	th.checkPostInSearchResults(t, p1.Id, results.Posts)

	params = &model.SearchParams{Terms: "beta"}
	results, apperr = th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
	require.Nil(t, apperr)

	require.Len(t, results.Posts, 1)
	th.checkPostInSearchResults(t, p1.Id, results.Posts)

	params = &model.SearchParams{Terms: "alpha"}
	results, apperr = th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
	require.Nil(t, apperr)

	require.Len(t, results.Posts, 1)
	th.checkPostInSearchResults(t, p1.Id, results.Posts)
}

func testSearchEmailsWithoutQuotes(t *testing.T, th *SearchTestHelper) {
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "message test@test.com", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	_, err = th.createPost(th.User.Id, th.ChannelBasic.Id, "message test2@test.com", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	params := &model.SearchParams{Terms: "test@test.com"}
	results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
	require.Nil(t, apperr)

	require.Len(t, results.Posts, 1)
	th.checkPostInSearchResults(t, p1.Id, results.Posts)
}

func testSupportSearchInComments(t *testing.T, th *SearchTestHelper) {
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "message test@test.com", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	r1, err := th.createReply(th.User.Id, "reply check", "", p1, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	params := &model.SearchParams{Terms: "reply"}
	results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
	require.Nil(t, apperr)

	require.Len(t, results.Posts, 1)
	th.checkPostInSearchResults(t, r1.Id, results.Posts)
}

func testSupportSearchTermsWithinLinks(t *testing.T, th *SearchTestHelper) {
	p1, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "message with link http://www.wikipedia.org/dolphins", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	params := &model.SearchParams{Terms: "wikipedia"}
	results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
	require.Nil(t, apperr)

	require.Len(t, results.Posts, 1)
	th.checkPostInSearchResults(t, p1.Id, results.Posts)
}

func testShouldNotReturnLinksEmbeddedInMarkdown(t *testing.T, th *SearchTestHelper) {
	_, err := th.createPost(th.User.Id, th.ChannelBasic.Id, "message with link [here](http://www.wikipedia.org/dolphins)", "", model.POST_DEFAULT, 0, false)
	require.Nil(t, err)
	defer th.deleteUserPosts(th.User.Id)

	params := &model.SearchParams{Terms: "wikipedia"}
	results, apperr := th.Store.Post().SearchPostsInTeamForUser([]*model.SearchParams{params}, th.User.Id, th.Team.Id, false, false, 0, 20)
	require.Nil(t, apperr)

	require.Len(t, results.Posts, 0)
}

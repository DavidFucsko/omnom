package webapp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type method string

const (
	GET   method = "GET"
	POST  method = "POST"
	PUT   method = "PUT"
	PATCH method = "PATCH"
	HEAD  method = "HEAD"
)

type EndpointArg struct {
	Name        string
	Type        string
	Required    bool
	Description string
}

type Endpoint struct {
	Name         string
	Path         string
	Method       method
	AuthRequired bool
	Handler      func(*gin.Context)
	Description  string
	Args         []*EndpointArg
}

var Endpoints []*Endpoint

func init() {
	Endpoints = []*Endpoint{
		&Endpoint{
			Name:         "Index",
			Path:         "/",
			Method:       GET,
			AuthRequired: false,
			Handler:      index,
			Description:  "Landing page",
		},
		&Endpoint{
			Name:         "Signup form",
			Path:         "/signup",
			Method:       GET,
			AuthRequired: false,
			Handler:      signup,
			Description:  "Signup page",
		},
		&Endpoint{
			Name:         "Signup processor",
			Path:         "/signup",
			Method:       POST,
			AuthRequired: false,
			Handler:      signup,
			Description:  "Signup form processor",
			Args: []*EndpointArg{
				&EndpointArg{
					Name:        "username",
					Type:        "string",
					Required:    true,
					Description: "Username of the new account",
				},
				&EndpointArg{
					Name:        "email",
					Type:        "string",
					Required:    true,
					Description: "Email address of the new account",
				},
			},
		},
		&Endpoint{
			Name:         "Login form",
			Path:         "/login",
			Method:       GET,
			AuthRequired: false,
			Handler:      login,
			Description:  "Login page",
		},
		&Endpoint{
			Name:         "Login processor",
			Path:         "/login",
			Method:       POST,
			AuthRequired: false,
			Handler:      login,
			Description:  "Login form processor",
			Args: []*EndpointArg{
				&EndpointArg{
					Name:        "username",
					Type:        "string",
					Required:    true,
					Description: "Username or email address of the new account",
				},
			},
		},
		&Endpoint{
			Name:         "Logout page",
			Path:         "/logout",
			Method:       GET,
			AuthRequired: false,
			Handler:      logout,
			Description:  "Destroys user session",
		},
		&Endpoint{
			Name:         "Public bookmark listing",
			Path:         "/bookmarks",
			Method:       GET,
			AuthRequired: false,
			Handler:      bookmarks,
			Description:  "Displays public bookmarks with optional filters",
			Args: []*EndpointArg{
				&EndpointArg{
					Name:        "query",
					Type:        "string",
					Required:    false,
					Description: "Search term to filter bookmarks by title",
				},
				&EndpointArg{
					Name:        "owner",
					Type:        "string",
					Required:    false,
					Description: "Search term to filter bookmarks by username",
				},
				&EndpointArg{
					Name:        "tag",
					Type:        "string",
					Required:    false,
					Description: "Search term to filter bookmarks by tag",
				},
				&EndpointArg{
					Name:        "domain",
					Type:        "string",
					Required:    false,
					Description: "Search term to filter bookmarks by domain",
				},
				&EndpointArg{
					Name:        "from",
					Type:        "date",
					Required:    false,
					Description: "Display only newer bookmarks. (Format: YYYY.MM.DD)",
				},
				&EndpointArg{
					Name:        "to",
					Type:        "date",
					Required:    false,
					Description: "Display only older bookmarks. (Format: YYYY.MM.DD)",
				},
				&EndpointArg{
					Name:        "search_in_snapshots",
					Type:        "boolean",
					Required:    false,
					Description: "Query parameter also applied to snapshot content. (Values: 0, 1)",
				},
				&EndpointArg{
					Name:        "search_in_notes",
					Type:        "boolean",
					Required:    false,
					Description: "Query parameter also applied to bookmark notes. (Values: 0, 1)",
				},
			},
		},
		&Endpoint{
			Name:         "Snapshot view with details",
			Path:         "/snapshot",
			Method:       GET,
			AuthRequired: false,
			Handler:      snapshotWrapper,
			Description:  "Displays bookmark snapshots with additional bookmark properties",
			Args: []*EndpointArg{
				&EndpointArg{
					Name:        "sid",
					Type:        "string",
					Required:    true,
					Description: "Snapshot key",
				},
				&EndpointArg{
					Name:        "bid",
					Type:        "int",
					Required:    true,
					Description: "Bookmark ID",
				},
			},
		},
		&Endpoint{
			Name:         "Fullscreen snapshot view",
			Path:         "/view_snapshot",
			Method:       GET,
			AuthRequired: false,
			Handler:      snapshot,
			Description:  "Displays bookmark snapshots as a fullscreen page",
			Args: []*EndpointArg{
				&EndpointArg{
					Name:        "sid",
					Type:        "string",
					Required:    true,
					Description: "Snapshot key",
				},
			},
		},
		&Endpoint{
			Name:         "Add bookmark",
			Path:         "/add_bookmark",
			Method:       POST,
			AuthRequired: false,
			Handler:      addBookmark,
			Description:  "Add new bookmark",
			Args: []*EndpointArg{
				&EndpointArg{
					Name:        "token",
					Type:        "string",
					Required:    true,
					Description: "Extension token. It can be found on the profile page",
				},
				&EndpointArg{
					Name:        "url",
					Type:        "URL",
					Required:    true,
					Description: "URL of the bookmark",
				},
				&EndpointArg{
					Name:        "title",
					Type:        "string",
					Required:    true,
					Description: "Title of the bookmark",
				},
				&EndpointArg{
					Name:        "notes",
					Type:        "string",
					Required:    false,
					Description: "Bookmark notes",
				},
				&EndpointArg{
					Name:        "favicon",
					Type:        "string",
					Required:    false,
					Description: "Data URL encoded favicon (https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/Data_URIs)",
				},
				&EndpointArg{
					Name:        "public",
					Type:        "boolean",
					Required:    false,
					Description: "Marks bookmark as public",
				},
				&EndpointArg{
					Name:        "tags",
					Type:        "string",
					Required:    false,
					Description: "Comma separated list of tags",
				},
				&EndpointArg{
					Name:        "snapshot_title",
					Type:        "string",
					Required:    false,
					Description: "Title of the snapshot",
				},
				&EndpointArg{
					Name:        "snapshot_text",
					Type:        "string",
					Required:    false,
					Description: "Text content of the snapshot",
				},
				&EndpointArg{
					Name:        "snapshot",
					Type:        "multipart file",
					Required:    false,
					Description: "Snapshot file",
				},
			},
		},
		&Endpoint{
			Name:         "Check bookmark",
			Path:         "/check_bookmark",
			Method:       GET,
			AuthRequired: false,
			Handler:      checkBookmark,
			Description:  "Checks if a bookmark is already exists",
			Args: []*EndpointArg{
				&EndpointArg{
					Name:        "token",
					Type:        "string",
					Required:    true,
					Description: "Extension token. It can be found on the profile page",
				},
				&EndpointArg{
					Name:        "url",
					Type:        "URL",
					Required:    true,
					Description: "URL of the bookmark",
				},
			},
		},
		&Endpoint{
			Name:         "View bookmark",
			Path:         "/bookmark",
			Method:       GET,
			AuthRequired: false,
			Handler:      viewBookmark,
			Description:  "Displays all details of a bookmark",
			Args: []*EndpointArg{
				&EndpointArg{
					Name:        "id",
					Type:        "int",
					Required:    true,
					Description: "Bookmark ID",
				},
			},
		},
		&Endpoint{
			Name:         "API documentation",
			Path:         "/api",
			Method:       GET,
			AuthRequired: false,
			Handler:      api,
			Description:  "Displays API documentation (this page)",
		},
		/****************************************\
		| LOGIN REQUIRED FOR THE ENDPOINTS BELOW |
		\****************************************/
		&Endpoint{
			Name:         "Profile page",
			Path:         "/profile",
			Method:       GET,
			AuthRequired: true,
			Handler:      profile,
			Description:  "Displays the profile page",
		},
		&Endpoint{
			Name:         "Generate addon token",
			Path:         "/generate_addon_token",
			Method:       GET,
			AuthRequired: true,
			Handler:      generateAddonToken,
			Description:  "Creates a new addon token",
		},
		&Endpoint{
			Name:         "Delete addon token",
			Path:         "/delete_addon_token",
			Method:       POST,
			AuthRequired: true,
			Handler:      deleteAddonToken,
			Description:  "Deletes an addon token",
			Args: []*EndpointArg{
				&EndpointArg{
					Name:        "id",
					Type:        "int",
					Required:    true,
					Description: "Token ID",
				},
			},
		},
		&Endpoint{
			Name:         "View personal bookmarks",
			Path:         "/my_bookmarks",
			Method:       GET,
			AuthRequired: true,
			Handler:      myBookmarks,
			Description:  "Displays bookmarks belongs to the current user with optional filters",
			Args: []*EndpointArg{
				&EndpointArg{
					Name:        "query",
					Type:        "string",
					Required:    false,
					Description: "Search term to filter bookmarks by title",
				},
				&EndpointArg{
					Name:        "tag",
					Type:        "string",
					Required:    false,
					Description: "Search term to filter bookmarks by tag",
				},
				&EndpointArg{
					Name:        "domain",
					Type:        "string",
					Required:    false,
					Description: "Search term to filter bookmarks by domain",
				},
				&EndpointArg{
					Name:        "from",
					Type:        "date",
					Required:    false,
					Description: "Display only newer bookmarks. (Format: YYYY.MM.DD)",
				},
				&EndpointArg{
					Name:        "to",
					Type:        "date",
					Required:    false,
					Description: "Display only older bookmarks. (Format: YYYY.MM.DD)",
				},
				&EndpointArg{
					Name:        "is_public",
					Type:        "boolean",
					Required:    false,
					Description: "Display only public bookmarks. (Values: 0, 1)",
				},
				&EndpointArg{
					Name:        "is_private",
					Type:        "boolean",
					Required:    false,
					Description: "Display only private bookmarks. (Values: 0, 1)",
				},
				&EndpointArg{
					Name:        "search_in_snapshots",
					Type:        "boolean",
					Required:    false,
					Description: "Query parameter also applied to snapshot content. (Values: 0, 1)",
				},
				&EndpointArg{
					Name:        "search_in_notes",
					Type:        "boolean",
					Required:    false,
					Description: "Query parameter also applied to bookmark notes. (Values: 0, 1)",
				},
			},
		},
		&Endpoint{
			Name:         "Edit bookmark",
			Path:         "/edit_bookmark",
			Method:       GET,
			AuthRequired: true,
			Handler:      editBookmark,
			Description:  "Displays a bookmark with all the editable properties",
			Args: []*EndpointArg{
				&EndpointArg{
					Name:        "id",
					Type:        "int",
					Required:    true,
					Description: "Bookmark ID",
				},
			},
		},
		&Endpoint{
			Name:         "Save bookmark",
			Path:         "/save_bookmark",
			Method:       POST,
			AuthRequired: true,
			Handler:      saveBookmark,
			Description:  "Saves an edited bookmark",
			Args: []*EndpointArg{
				&EndpointArg{
					Name:        "id",
					Type:        "int",
					Required:    true,
					Description: "Bookmark ID",
				},
				&EndpointArg{
					Name:        "title",
					Type:        "string",
					Required:    true,
					Description: "Title of the bookmark",
				},
				&EndpointArg{
					Name:        "notes",
					Type:        "string",
					Required:    false,
					Description: "Bookmark notes",
				},
				&EndpointArg{
					Name:        "public",
					Type:        "boolean",
					Required:    false,
					Description: "Bookmark is publicly accessible. (Omit this argument in case of private bookmarks)",
				},
			},
		},
		&Endpoint{
			Name:         "Delete snapshot",
			Path:         "/delete_snapshot",
			Method:       POST,
			AuthRequired: true,
			Handler:      deleteSnapshot,
			Description:  "Deletes a snapshot",
			Args: []*EndpointArg{
				&EndpointArg{
					Name:        "bid",
					Type:        "int",
					Required:    true,
					Description: "Bookmark ID",
				},
				&EndpointArg{
					Name:        "sid",
					Type:        "int",
					Required:    true,
					Description: "Snapshot ID",
				},
			},
		},
		&Endpoint{
			Name:         "Delete bookmark",
			Path:         "/delete_bookmark",
			Method:       POST,
			AuthRequired: true,
			Handler:      deleteBookmark,
			Description:  "Deletes a bookmark",
			Args: []*EndpointArg{
				&EndpointArg{
					Name:        "id",
					Type:        "int",
					Required:    true,
					Description: "Bookmark ID",
				},
			},
		},
		&Endpoint{
			Name:         "Add tag",
			Path:         "/add_tag",
			Method:       POST,
			AuthRequired: true,
			Handler:      addTag,
			Description:  "Add tag to a bookmark",
			Args: []*EndpointArg{
				&EndpointArg{
					Name:        "bid",
					Type:        "int",
					Required:    true,
					Description: "Bookmark ID",
				},
				&EndpointArg{
					Name:        "tag",
					Type:        "string",
					Required:    true,
					Description: "Tag string",
				},
			},
		},
		&Endpoint{
			Name:         "Delete tag",
			Path:         "/delete_tag",
			Method:       POST,
			AuthRequired: true,
			Handler:      deleteTag,
			Description:  "Delete tag's bookmark",
			Args: []*EndpointArg{
				&EndpointArg{
					Name:        "bid",
					Type:        "int",
					Required:    true,
					Description: "Bookmark ID",
				},
				&EndpointArg{
					Name:        "tid",
					Type:        "int",
					Required:    true,
					Description: "Tag ID",
				},
			},
		},
	}
}

func api(c *gin.Context) {
	renderHTML(c, http.StatusOK, "api", map[string]interface{}{
		"Endpoints": Endpoints,
	})
}

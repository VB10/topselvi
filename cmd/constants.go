package cmd

//HTTP Request Type
const GET = "GET"
const POST = "POST"

//Youtube Watch Prefix

const YoutubeWatchPrefix = "https://www.youtube.com/watch?v="

//collections string
const FirestoreVideos = "videos"

const FirebaseAuthEndPoint = "https://securetoken.googleapis.com/v1/token?key=" + FirebaseApiKey
const FirebaseAuthSigninCustomToken = "https://identitytoolkit.googleapis.com/v1/accounts:signInWithCustomToken?key=" + FirebaseApiKey
const FirestoreUsers = "users"
const FirebaseApiKey = "AIzaSyCRyClXcn4igHaFNd9TmZuJL1DkHgAuV9Q"

//custom error
const ModelInvalid = "Your request model body, doesn't validate."

const FirebaseQueryUserID = "user_id"

const QueryApiKey = "token"
const QueryUserId = "userID"
const QueryUserToken = "userToken"
const QueryIDToken = "idToken"
const QueryKey = "key"

//firebase query param

const FbUid = "user_id"

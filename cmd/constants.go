package cmd

//HTTP Request Type
const GET = "GET"
const POST = "POST"

//Youtube Watch Prefix

const YoutubeWatchPrefix = "https://www.youtube.com/watch?v="

//collections string
const FirestoreVideos = "videos"
const FirebaseApiKey = "AIzaSyCRyClXcn4igHaFNd9TmZuJL1DkHgAuV9Q"
const FirebaseApiKey_SUB4 = "AIzaSyCW9_qmsAM0Wm4KDJwI1zozyloDunF_CtY"

const FirebaseAuthEndPoint = "https://securetoken.googleapis.com/v1/token?key=" + FirebaseApiKey
const FirebaseAuthSigninCustomToken = "https://identitytoolkit.googleapis.com/v1/accounts:signInWithCustomToken?key=" + FirebaseApiKey_SUB4
const FirestoreUsers = "users"

//custom error
const ModelInvalid = "Your request model body, doesn't validate."
const ApiKeyEmptyError = "We need api key in request header."

const FirebaseQueryUserID = "user_id"

const QueryApiKey = "token"
const QueryUserId = "userID"
const QueryUserToken = "userToken"
const QueryIDToken = "idToken"
const QueryKey = "key"

const NewUserWallet = 100

//firebase query param

const FbUid = "user_id"
const FBName = "name"
const FBEmail = "email"
const FBPicture= "picture"
const JWTFirebaseKeyError = "key is of invalid type"

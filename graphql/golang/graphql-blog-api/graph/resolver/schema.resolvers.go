package resolver

import (
	"context"
	"log"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"example.com/graphql-blog-api/graph/generated"
	"example.com/graphql-blog-api/graph/model"
	"example.com/graphql-blog-api/graph"
)

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	newPost := &model.Post{
		ID:      uuid.New().String(),
		Title:   input.Title,
		Content: input.Content,
		Author:  input.Author,
	}

	collection := graph.GetCollection("posts")
	_, err := collection.InsertOne(ctx, newPost)
	if err != nil {
		log.Printf("Error inserting post: %v", err)
		return nil, err
	}

	return newPost, nil
}

// UpdatePost is the resolver for the updatePost field.
func (r *mutationResolver) UpdatePost(ctx context.Context, input model.UpdatePost) (*model.Post, error) {
	collection := graph.GetCollection("posts")
	filter := bson.M{"id": input.ID}
	update := bson.M{"$set": bson.M{"title": input.Title, "content": input.Content}}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Printf("Error updating post: %v", err)
		return nil, err
	}

	var updatedPost model.Post
	err = collection.FindOne(ctx, filter).Decode(&updatedPost)
	if err != nil {
		return nil, err
	}

	return &updatedPost, nil
}

// DeletePost is the resolver for the deletePost field.
func (r *mutationResolver) DeletePost(ctx context.Context, id string) (bool, error) {
	collection := graph.GetCollection("posts")
	filter := bson.M{"id": id}

	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Printf("Error deleting post: %v", err)
		return false, err
	}

	return res.DeletedCount > 0, nil
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	collection := graph.GetCollection("posts")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Error retrieving posts: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var posts []*model.Post
	if err := cursor.All(ctx, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	collection := graph.GetCollection("posts")
	filter := bson.M{"id": id}

	var post model.Post
	err := collection.FindOne(ctx, filter).Decode(&post)
	if err != nil {
		log.Printf("Error retrieving post: %v", err)
		return nil, err
	}

	return &post, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

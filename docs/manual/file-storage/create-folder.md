# Create folder

You can easily allow users to create a folder via the File Management module of Space Cloud by calling `createFolder` on frontend. Here's a code snippet to do so:

 <div class="row tabs-wrapper">
  <div class="col s12" style="padding:0">
    <ul class="tabs">
      <li class="tab col s2"><a class="active" href="#create-folder-js">Javascript</a></li>
      <li class="tab col s2"><a href="#create-folder-java">Java</a></li>
      <li class="tab col s2"><a href="#create-folder-python">Python</a></li>
      <li class="tab col s2"><a href="#create-folder-golang">Golang</a></li>
    </ul>
  </div>
  <div id="create-folder-js" class="col s12" style="padding:0">
    <pre>
      <code class="javascript">
import { API } from "space-api";

// Initialize api with the project name and url of the space cloud
const api = new API("todo-app", "http://localhost:8080");

// Create a folder
api.FileStore()
  .createFolder("/some-path", "some-folder")
  .then(res => {
    if (res.status === 200) {
      // Folder created successfully
    }
    // Error creating folder
  })
  .catch(ex => {
    // Exception occured while processing request
  });
      </code>
    </pre>
  </div>
  <div id="create-folder-java" class="col s12" style="padding:0">
    <pre>
      <code class="java">
// Java client coming soon!      
      </code>
    </pre>
  </div>
 <div id="create-folder-python" class="col s12" style="padding:0">
    <pre>
      <code class="python">
# Python client coming soon!
      </code>
    </pre>
  </div>
  <div id="create-folder-golang" class="col s12" style="padding:0">
    <pre>
      <code class="golang">
// Golang client coming soon!
      </code>
    </pre>
  </div>
</div>

The `createFolder` function takes two parameters and creates a folder. The two parameters are as follows:
- **path:** The path at which to create the folder.
- **name:** The name of the folder.

The `path` can be nested as well. For e.g a `path` - /folder1/folder2 would mean to create the folder inside folder2 which is in folder1. If any of the folders mentioned in the `path` were not present, they would be created before creating the specified folder.

## Response

A response object sent by the server contains the **status** fields explained below:

**status:** Number describing the status of the upload operation. Following values are possible:

- 200 - Successful creation of folder
- 401 - Request was unauthenticated
- 403 - Request was unauthorized
- 500 - Internal server error

## Next steps

Now you know how to create a folder. So let's check how to list all files / folders within that folder.

<div class="btns-wrapper">
  <a href="/docs/file-storage/upload-file" class="waves-effect waves-light btn primary-btn-border btn-small">
    <i class="material-icons btn-with-icon">arrow_back</i>Previous
  </a>
  <a href="/docs/file-storage/list-files" class="waves-effect waves-light btn primary-btn-fill btn-small">
    Next<i class="material-icons btn-with-icon">arrow_forward</i>
  </a>
</div>

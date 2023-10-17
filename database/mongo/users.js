db = db.getSiblingDB("videodb");
db.createUser({
  user: "video-user",
  pwd: "video-pwd",
  roles: [
    {
      role: "readWrite",
      db: "videodb",
    },
    {
      role: "dbAdmin",
      db: "videodb",
    },
  ],
});
